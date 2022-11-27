package server

import (
	"context"
	"net"

	"google.golang.org/grpc"

	"go-park-mail-ru/2022_2_BugOverload/internal/image/delivery/grpc/models"
	proto "go-park-mail-ru/2022_2_BugOverload/internal/image/delivery/grpc/protobuf"
	"go-park-mail-ru/2022_2_BugOverload/internal/image/service"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/wrapper"
)

type ImageServiceGRPCServer struct {
	grpcServer *grpc.Server

	imageManager service.ImageService
}

func NewImageServiceGRPCServer(grpcServer *grpc.Server, im service.ImageService) *ImageServiceGRPCServer {
	return &ImageServiceGRPCServer{
		grpcServer:   grpcServer,
		imageManager: im,
	}
}

func (s *ImageServiceGRPCServer) StartGRPCServer(listenURL string) error {
	lis, err := net.Listen("tcp", listenURL)
	if err != nil {
		return err
	}

	proto.RegisterImageServiceServer(s.grpcServer, s)

	return s.grpcServer.Serve(lis)
}

func (s *ImageServiceGRPCServer) GetImage(ctx context.Context, image *proto.Image) (*proto.Image, error) {
	imageRequest, err := s.imageManager.GetImage(ctx, models.NewImage(image))
	if err != nil {
		return &proto.Image{}, wrapper.DefaultHandlerGRPCError(ctx, err)
	}

	return models.NewImageProto(&imageRequest), nil
}

func (s *ImageServiceGRPCServer) UpdateImage(ctx context.Context, image *proto.Image) (*proto.Nothing, error) {
	err := s.imageManager.UpdateImage(ctx, models.NewImage(image))
	if err != nil {
		return &proto.Nothing{}, wrapper.DefaultHandlerGRPCError(ctx, err)
	}

	return &proto.Nothing{}, nil
}
