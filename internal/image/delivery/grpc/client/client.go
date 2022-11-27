package client

import (
	"context"

	stdErrors "github.com/pkg/errors"
	"google.golang.org/grpc"

	"go-park-mail-ru/2022_2_BugOverload/internal/image/delivery/grpc/models"
	proto "go-park-mail-ru/2022_2_BugOverload/internal/image/delivery/grpc/protobuf"
	modelsGlobal "go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/wrapper"
)

type ImageService interface {
	GetImage(ctx context.Context, image *modelsGlobal.Image) (modelsGlobal.Image, error)
	UpdateImage(ctx context.Context, image *modelsGlobal.Image) error
}

type ImageServiceGRPCClient struct {
	imageClient proto.ImageServiceClient
}

func NewImageServiceGRPSClient(con *grpc.ClientConn) ImageService {
	return &ImageServiceGRPCClient{
		imageClient: proto.NewImageServiceClient(con),
	}
}

func (c *ImageServiceGRPCClient) GetImage(ctx context.Context, image *modelsGlobal.Image) (modelsGlobal.Image, error) {
	imageProtoResponse, err := c.imageClient.GetImage(pkg.GetDefInfoMicroService(ctx), models.NewImageProto(image))
	if err != nil {
		return modelsGlobal.Image{}, wrapper.GRPCErrorConvert(stdErrors.Wrap(err, "GetImage"))
	}

	return *models.NewImage(imageProtoResponse), nil
}

func (c *ImageServiceGRPCClient) UpdateImage(ctx context.Context, image *modelsGlobal.Image) error {
	_, err := c.imageClient.UpdateImage(pkg.GetDefInfoMicroService(ctx), models.NewImageProto(image))
	if err != nil {
		return wrapper.GRPCErrorConvert(stdErrors.Wrap(err, "UpdateImage"))
	}

	return nil
}
