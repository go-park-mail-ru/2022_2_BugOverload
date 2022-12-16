package repository

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	innerPKG "go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/pkg"
	"strconv"
)

type ImageS3 struct {
	Bucket string `json:"bucket" example:"film/"`
	Key    string `json:"key" example:"posters/hor/1.jpeg"`
	Bytes  []byte `json:"-"`
}

func NewImageS3Pattern(imageParams *models.Image) (*ImageS3, error) {
	image := &ImageS3{}

	switch imageParams.Object {
	case innerPKG.ImageObjectFilmPosterVer:
		image.Bucket = innerPKG.FilmsBucket
		image.Key = "posters/ver/"
	case innerPKG.ImageObjectFilmPosterHor:
		image.Bucket = innerPKG.FilmsBucket
		image.Key = "posters/hor/"
	case innerPKG.ImageObjectFilmImage:
		image.Bucket = innerPKG.FilmsBucket
		image.Key = "images/"

	case innerPKG.ImageObjectUserAvatar:
		image.Bucket = innerPKG.UsersBucket
		image.Key = "avatars/"

	case innerPKG.ImageObjectPersonAvatar:
		image.Bucket = innerPKG.PersonsBucket
		image.Key = "avatars/"
	case innerPKG.ImageObjectPersonImage:
		image.Bucket = innerPKG.PersonsBucket
		image.Key = "images/"

	case innerPKG.ImageObjectCollectionImage:
		image.Bucket = innerPKG.CollectionsBucket
		image.Key = "posters/"

	case innerPKG.ImageObjectDefault:
		image.Bucket = innerPKG.DefBucket

		if imageParams.Key == "login" || imageParams.Key == "signup" {
			randID := pkg.RandMaxInt(innerPKG.ImageCountSignupLogin) + 1

			imageParams.Key = strconv.Itoa(randID)
		}
	default:
		return nil, errors.ErrBadImageType
	}

	image.Key += imageParams.Key + ".webp"

	image.Bytes = imageParams.Bytes

	return image, nil
}
