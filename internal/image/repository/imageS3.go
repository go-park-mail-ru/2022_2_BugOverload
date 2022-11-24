package repository

import (
	"bytes"
	"context"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/sqltools"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	stdErrors "github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	innerPKG "go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
)

// ImageRepository provides the versatility of images repositories.
type ImageRepository interface {
	GetImage(ctx context.Context, image *models.Image) (models.Image, error)
	UpdateImage(ctx context.Context, image *models.Image) error
}

// imageS3WithPostgres is implementation repository of users in S3 corresponding to the ImageRepository interface.
type imageS3WithPostgres struct {
	downloaderS3 *s3manager.Downloader
	uploaderS3   *s3manager.Uploader

	database *sqltools.Database
}

// NewImageS3 is constructor for imageS3WithPostgres. Accepts only mutex.
func NewImageS3(config *innerPKG.Config, database *sqltools.Database) ImageRepository {
	awsConfig := aws.NewConfig()
	awsConfig.Endpoint = aws.String(config.S3.Endpoint)

	awsConfig.Credentials = credentials.NewEnvCredentials()

	sess, err := session.NewSession(awsConfig)
	if err != nil {
		logrus.Fatalf("Init ImageS3Storage fatal error: NewSession: %s", err)
	}

	res := &imageS3WithPostgres{
		s3manager.NewDownloader(sess),
		s3manager.NewUploader(sess),
		database,
	}

	return res
}

// GetImage getting image by path
func (i *imageS3WithPostgres) GetImage(ctx context.Context, image *models.Image) (models.Image, error) {
	imageS3Pattern, err := NewImageS3Pattern(image)
	if err != nil {
		return models.Image{}, stdErrors.WithMessagef(err,
			"Err: params input: image key - [%s], object - [%s]",
			image.Key, image.Object)
	}

	res := make([]byte, innerPKG.BufSizeImage)

	w := aws.NewWriteAtBuffer(res)

	getObjectInput := &s3.GetObjectInput{
		Bucket: aws.String(imageS3Pattern.Bucket),
		Key:    aws.String(imageS3Pattern.Key),
	}

	realSize, err := i.downloaderS3.DownloadWithContext(ctx, w, getObjectInput)
	if err != nil {
		var awsErr awserr.Error
		var errOut error

		if stdErrors.As(err, &awsErr) {
			switch awsErr.Code() {
			case s3.ErrCodeNoSuchBucket:
				errOut = errors.ErrImageNotFound
			case s3.ErrCodeNoSuchKey:
				errOut = errors.ErrImageNotFound
			default:
				errOut = errors.ErrImage
			}

			return models.Image{}, stdErrors.WithMessagef(errOut,
				"Err: params input: image key - [%s], object - [%s]. Special Error [%s]",
				image.Key, image.Object, err)
		}
	}

	return models.Image{Bytes: w.Bytes()[:realSize]}, nil
}

// UpdateImage download image into storage
func (i *imageS3WithPostgres) UpdateImage(ctx context.Context, image *models.Image) error {
	imageS3Pattern, err := NewImageS3Pattern(image)
	if err != nil {
		return stdErrors.WithMessagef(err,
			"Err: params input: image key - [%s], object - [%s], size image [%d]",
			image.Key, image.Object, len(image.Bytes))
	}

	body := bytes.NewReader(image.Bytes)

	getObjectInput := &s3manager.UploadInput{
		Bucket: aws.String(imageS3Pattern.Bucket),
		Key:    aws.String(imageS3Pattern.Key),
		Body:   body,
	}

	_, err = i.uploaderS3.UploadWithContext(ctx, getObjectInput)
	if err != nil {
		return stdErrors.WithMessagef(errors.ErrImage,
			"Err: params input: image key - [%s], object - [%s], size image [%d]. Special Error [%s]",
			imageS3Pattern.Key, imageS3Pattern.Bucket, len(image.Bytes), err)
	}

	err = i.UpdateImageInfo(ctx, image)
	if err != nil {
		return stdErrors.Wrap(err, "UpdateImageInfo")
	}

	return nil
}
