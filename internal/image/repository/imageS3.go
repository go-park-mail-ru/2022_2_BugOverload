package repository

import (
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
	"go-park-mail-ru/2022_2_BugOverload/pkg"
)

// ImageRepository provides the versatility of images repositories.
type ImageRepository interface {
	GetImage(ctx context.Context, image *models.Image) ([]byte, error)
	PutImage(ctx context.Context) error
}

// imageS3 is implementation repository of users in S3 corresponding to the ImageRepository interface.
type imageS3 struct {
	downloaderS3 *s3manager.Downloader
	uploaderS3   *s3manager.Uploader
}

// NewImageS3 is constructor for imageS3. Accepts only mutex.
func NewImageS3(config *innerPKG.Config) (ImageRepository, error) {
	awsConfig := aws.NewConfig()
	awsConfig.Region = aws.String(config.S3.Region)
	awsConfig.Endpoint = aws.String(config.S3.Endpoint)

	token, _ := pkg.CryptoRandString(innerPKG.TokenS3Length)
	awsConfig.Credentials = credentials.NewStaticCredentials(config.S3.ID, config.S3.Secret, token)

	sess, err := session.NewSession(awsConfig)
	if err != nil {
		logrus.Error("Init fatal error: NewSession: ", err)

		return nil, err
	}

	res := &imageS3{
		s3manager.NewDownloader(sess),
		s3manager.NewUploader(sess),
	}

	return res, nil
}

// GetImage getting image by path
func (is *imageS3) GetImage(ctx context.Context, image *models.Image) ([]byte, error) {
	res := make([]byte, innerPKG.BufSizeImage)

	w := aws.NewWriteAtBuffer(res)

	getObjectInput := &s3.GetObjectInput{
		Bucket: aws.String(image.Bucket),
		Key:    aws.String(image.Item),
	}

	_, err := is.downloaderS3.Download(w, getObjectInput)
	if err != nil {
		return nil, errors.ErrImageNotFound
	}

	return w.Bytes(), nil
}

// PutImage download image into storage
func (is *imageS3) PutImage(ctx context.Context) error {
	return nil
}
