package repository

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg"
)

type ImageS3 struct {
	Bucket string `json:"bucket" example:"film/"`
	Key    string `json:"key" example:"posters/hor/1.jpeg"`
	Bytes  []byte `json:"-"`
}

func NewImageS3Pattern(imageParams *models.Image) *ImageS3 {
	image := &ImageS3{}

	switch imageParams.Object {
	case pkg.ImageObjectFilmPosterVer:
		image.Bucket = "film/"
		image.Key = "posters/ver/"
	case pkg.ImageObjectFilmPosterHor:
		image.Bucket = "film/"
		image.Key = "posters/hor/"
	case pkg.ImageObjectAvatar:
		image.Bucket = "users/"
		image.Key = "avatar/"
	case pkg.ImageObjectDefault:
		image.Bucket = "default/"
	}

	image.Key += imageParams.Key + ".jpeg"

	image.Bytes = imageParams.Bytes

	return image
}
