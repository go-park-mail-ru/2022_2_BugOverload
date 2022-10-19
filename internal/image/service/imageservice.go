package service

import (
	"context"

	stdErrors "github.com/pkg/errors"

	"go-park-mail-ru/2022_2_BugOverload/internal/image/repository"
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
)

// ImageService provides universal service for work with images.
type ImageService interface {
	GetImage(ctx context.Context, image *models.Image) ([]byte, error)
	PutImage(ctx context.Context) error
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

// GetImage is the service that accesses the interface ImageRepository.
func (i *imageService) GetImage(ctx context.Context, image *models.Image) ([]byte, error) {
	binImage, err := i.imageRepo.GetImage(ctx, image)
	if err != nil {
		return binImage, stdErrors.Wrap(err, "GetImage")
	}

	return binImage, nil
}

// PutImage is the service that accesses the interface ImageRepository
func (i *imageService) PutImage(ctx context.Context) error {
	err := i.imageRepo.PutImage(ctx)
	if err != nil {
		return stdErrors.Wrap(err, "PutImage")
	}

	return nil
}
