package server

import (
	"context"
	"net"

	"google.golang.org/grpc"

	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/wrapper"
	"go-park-mail-ru/2022_2_BugOverload/internal/warehouse/delivery/grpc/models"
	proto "go-park-mail-ru/2022_2_BugOverload/internal/warehouse/delivery/grpc/protobuf"
	"go-park-mail-ru/2022_2_BugOverload/internal/warehouse/service"
)

type WarehouseServiceGRPCServer struct {
	grpcServer *grpc.Server

	collectionManager service.CollectionService
	filmManager       service.FilmService
	personManager     service.PersonService
}

func NewWarehouseServiceGRPCServer(grpcServer *grpc.Server, cm service.CollectionService, fm service.FilmService, pm service.PersonService) *WarehouseServiceGRPCServer {
	return &WarehouseServiceGRPCServer{
		grpcServer:        grpcServer,
		collectionManager: cm,
		filmManager:       fm,
		personManager:     pm,
	}
}

func (s *WarehouseServiceGRPCServer) StartGRPCServer(listenURL string) error {
	lis, err := net.Listen("tcp", listenURL)
	if err != nil {
		return err
	}

	proto.RegisterWarehouseServiceServer(s.grpcServer, s)

	return s.grpcServer.Serve(lis)
}

func (s *WarehouseServiceGRPCServer) GetRecommendation(ctx context.Context, nothing *proto.Nothing) (*proto.Film, error) {
	film, err := s.filmManager.GetRecommendation(ctx)
	if err != nil {
		return &proto.Film{}, wrapper.DefaultHandlerGRPCError(ctx, err)
	}

	return models.NewFilmProto(&film), nil
}

func (s *WarehouseServiceGRPCServer) GetFilmByID(ctx context.Context, params *proto.GetFilmParams) (*proto.Film, error) {
	filmRequest, paramsRequest := models.NewGetFilmParams(params)

	film, err := s.filmManager.GetFilmByID(ctx, filmRequest, paramsRequest)
	if err != nil {
		return &proto.Film{}, wrapper.DefaultHandlerGRPCError(ctx, err)
	}

	return models.NewFilmProto(&film), nil
}

func (s *WarehouseServiceGRPCServer) GetReviewsByFilmID(ctx context.Context, params *proto.GetFilmReviewsParams) (*proto.Reviews, error) {
	reviews, err := s.filmManager.GetReviewsByFilmID(ctx, models.NewGetFilmReviewsParams(params))
	if err != nil {
		return &proto.Reviews{}, wrapper.DefaultHandlerGRPCError(ctx, err)
	}

	return models.NewReviewsProto(reviews), nil
}

func (s *WarehouseServiceGRPCServer) GetStdCollection(ctx context.Context, params *proto.GetStdCollectionParams) (*proto.Collection, error) {
	collection, err := s.collectionManager.GetStdCollection(ctx, models.NewGetStdCollectionParams(params))
	if err != nil {
		return &proto.Collection{}, wrapper.DefaultHandlerGRPCError(ctx, err)
	}

	return models.NewCollectionProto(&collection), nil
}

func (s *WarehouseServiceGRPCServer) GetPremieresCollection(ctx context.Context, params *proto.GetStdCollectionParams) (*proto.Collection, error) {
	collection, err := s.collectionManager.GetPremieresCollection(ctx, models.NewGetStdCollectionParams(params))
	if err != nil {
		return &proto.Collection{}, wrapper.DefaultHandlerGRPCError(ctx, err)
	}

	return models.NewCollectionProto(&collection), nil
}

func (s *WarehouseServiceGRPCServer) GetPersonByID(ctx context.Context, params *proto.GetPersonParams) (*proto.Person, error) {
	personRequest, paramsRequest := models.NewGetPersonParams(params)

	personRepo, err := s.personManager.GetPersonByID(ctx, personRequest, paramsRequest)
	if err != nil {
		return &proto.Person{}, wrapper.DefaultHandlerGRPCError(ctx, err)
	}

	return models.NewPersonProto(&personRepo), nil
}
