package repository

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
)

const (
	ImageObjectFilmPosterHor = "film_hor"
	ImageObjectFilmPosterVer = "film_ver"
	ImageObjectDefault       = "default"
)

type ImageS3 struct {
	Bucket string `json:"bucket" example:"films/"`
	Key    string `json:"key" example:"posters/hor/1.jpeg"`
	Bytes  []byte `json:"-"`
}

func NewImageS3Pattern(imageParams *models.Image) *ImageS3 {
	image := &ImageS3{}

	switch imageParams.Object {
	case ImageObjectFilmPosterVer:
		image.Bucket = "films/"
		image.Key = "posters/ver/"
	case ImageObjectFilmPosterHor:
		image.Bucket = "films/"
		image.Key = "posters/hor/"
	case ImageObjectDefault:
		image.Bucket = "default/"
	}

	image.Key += imageParams.Key + ".jpeg"

	image.Bytes = imageParams.Bytes

	return image
}
