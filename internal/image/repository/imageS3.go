package repository

import (
	"bytes"
	"context"

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
	DownloadImage(ctx context.Context, image *models.Image) (models.Image, error)
	UploadImage(ctx context.Context, image *models.Image) error
}

// imageS3 is implementation repository of users in S3 corresponding to the ImageRepository interface.
type imageS3 struct {
	downloaderS3 *s3manager.Downloader
	uploaderS3   *s3manager.Uploader
}

// NewImageS3 is constructor for imageS3. Accepts only mutex.
func NewImageS3(config *innerPKG.Config) ImageRepository {
	awsConfig := aws.NewConfig()
	awsConfig.Endpoint = aws.String(config.S3.Endpoint)

	awsConfig.Credentials = credentials.NewEnvCredentials()

	sess, err := session.NewSession(awsConfig)
	if err != nil {
		logrus.Fatalf("Init ImageS3Storage fatal error: NewSession: %s", err.Error())
	}

	res := &imageS3{
		s3manager.NewDownloader(sess),
		s3manager.NewUploader(sess),
	}

	return res
}

// DownloadImage getting image by path
func (is *imageS3) DownloadImage(ctx context.Context, image *models.Image) (models.Image, error) {
	imageS3Pattern, err := NewImageS3Pattern(image)
	if err != nil {
		return models.Image{}, stdErrors.WithMessagef(errors.ErrImage,
			"Err: params input: image key - [%s], object - [%s], size image [%d]. Special Error [%s]",
			image.Key, image.Object, len(image.Bytes), err)
	}

	res := make([]byte, innerPKG.BufSizeImage)

	w := aws.NewWriteAtBuffer(res)

	getObjectInput := &s3.GetObjectInput{
		Bucket: aws.String(imageS3Pattern.Bucket),
		Key:    aws.String(imageS3Pattern.Key),
	}

	realSize, err := is.downloaderS3.DownloadWithContext(ctx, w, getObjectInput)
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
				"Err: params input: image key - [%s], object - [%s], size image [%d]. Special Error [%s]",
				image.Key, image.Object, len(image.Bytes), err)
		}
	}

	return models.Image{Bytes: w.Bytes()[:realSize]}, nil
}

// UploadImage download image into storage
func (is *imageS3) UploadImage(ctx context.Context, image *models.Image) error {
	imageS3Pattern, err := NewImageS3Pattern(image)
	if err != nil {
		return stdErrors.WithMessagef(errors.ErrImage,
			"Err: params input: image key - [%s], object - [%s], size image [%d]. Special Error [%s]",
			image.Key, image.Object, len(image.Bytes), err)
	}

	body := bytes.NewReader(image.Bytes)

	getObjectInput := &s3manager.UploadInput{
		Bucket: aws.String(imageS3Pattern.Bucket),
		Key:    aws.String(imageS3Pattern.Key),
		Body:   body,
	}

	_, err = is.uploaderS3.UploadWithContext(ctx, getObjectInput)
	if err != nil {
		return stdErrors.WithMessagef(errors.ErrImage,
			"Err: params input: image key - [%s], object - [%s], size image [%d]. Special Error [%s]",
			image.Key, image.Object, len(image.Bytes), err)
	}

	return nil
}
