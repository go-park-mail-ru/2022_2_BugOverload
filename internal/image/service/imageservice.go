package service

import (
	"context"

	stdErrors "github.com/pkg/errors"

	repoImage "go-park-mail-ru/2022_2_BugOverload/internal/image/repository"
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
)

//go:generate mockgen -source imageservice.go -destination mocks/mockimageservice.go -package mockImageService

// ImageService provides universal service for work with images.
type ImageService interface {
	GetImage(ctx context.Context, image *models.Image) (models.Image, error)
	UpdateImage(ctx context.Context, image *models.Image) error
}

// imageService is implementation for users service corresponding to the ImageService interface.
type imageService struct {
	imageRepo repoImage.ImageRepository
}

// NewImageService is constructor for imageService. Accepts ImageService interfaces.
func NewImageService(ur repoImage.ImageRepository) ImageService {
	return &imageService{
		imageRepo: ur,
	}
}

// GetImage is the service that accesses the interface ImageRepository.
func (i *imageService) GetImage(ctx context.Context, image *models.Image) (models.Image, error) {
	binImage, err := i.imageRepo.GetImage(ctx, image)
	if err != nil {
		return models.Image{}, stdErrors.Wrap(err, "GetImage")
	}

	return binImage, nil
}

// UpdateImage is the service that accesses the interface ImageRepository
func (i *imageService) UpdateImage(ctx context.Context, image *models.Image) error {
	err := i.imageRepo.UpdateImage(ctx, image)
	if err != nil {
		return stdErrors.Wrap(err, "UpdateImage")
	}

	return nil
}
