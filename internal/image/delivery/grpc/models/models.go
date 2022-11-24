package models

import (
	proto "go-park-mail-ru/2022_2_BugOverload/internal/image/delivery/grpc/protobuf"
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
)

func NewImageProto(image *models.Image) *proto.Image {
	return &proto.Image{
		Key:    image.Key,
		Object: image.Object,
		Bytes:  image.Bytes,
	}
}

func NewImage(image *proto.Image) *models.Image {
	return &models.Image{
		Key:    image.Key,
		Object: image.Object,
		Bytes:  image.Bytes,
	}
}
