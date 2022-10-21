package repository

import (
	"bytes"
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

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
	awsConfig.Region = aws.String(config.S3.Region)
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
	imageS3Pattern := NewImageS3Pattern(image)

	res := make([]byte, innerPKG.BufSizeImage)

	w := aws.NewWriteAtBuffer(res)

	getObjectInput := &s3.GetObjectInput{
		Bucket: aws.String(imageS3Pattern.Bucket),
		Key:    aws.String(imageS3Pattern.Key),
	}

	realSize, err := is.downloaderS3.DownloadWithContext(ctx, w, getObjectInput)
	if err != nil {
		return models.Image{}, errors.ErrImageNotFound
	}

	return models.Image{Bytes: w.Bytes()[:realSize]}, nil
}

// UploadImage download image into storage
func (is *imageS3) UploadImage(ctx context.Context, image *models.Image) error {
	imageS3Pattern := NewImageS3Pattern(image)

	body := bytes.NewReader(image.Bytes)

	getObjectInput := &s3manager.UploadInput{
		Bucket: aws.String(imageS3Pattern.Bucket),
		Key:    aws.String(imageS3Pattern.Key),
		Body:   body,
	}

	_, err := is.uploaderS3.UploadWithContext(ctx, getObjectInput)
	if err != nil {
		return errors.ErrImageNotFound
	}

	return nil
}
