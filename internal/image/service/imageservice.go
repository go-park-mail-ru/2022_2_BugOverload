package service

import (
	"context"

	stdErrors "github.com/pkg/errors"

	"go-park-mail-ru/2022_2_BugOverload/internal/image/repository"
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
)

// ImageService provides universal service for work with images.
type ImageService interface {
	DownloadImage(ctx context.Context, image *models.Image) (models.Image, error)
	UploadImage(ctx context.Context, image *models.Image) error
}

// imageService is implementation for users service corresponding to the ImageService interface.
type imageService struct {
	imageRepo repository.ImageRepository
}

// NewImageService is constructor for imageService. Accepts UserRepository interfaces.
func NewImageService(ur repository.ImageRepository) ImageService {
	return &imageService{
		imageRepo: ur,
	}
}

// DownloadImage is the service that accesses the interface ImageRepository.
func (i *imageService) DownloadImage(ctx context.Context, image *models.Image) (models.Image, error) {
	binImage, err := i.imageRepo.DownloadImage(ctx, image)
	if err != nil {
		return binImage, stdErrors.Wrap(err, "DownloadImage")
	}

	return binImage, nil
}

// UploadImage is the service that accesses the interface ImageRepository
func (i *imageService) UploadImage(ctx context.Context, image *models.Image) error {
	err := i.imageRepo.UploadImage(ctx, image)
	if err != nil {
		return stdErrors.Wrap(err, "UploadImage")
	}

	return nil
}
