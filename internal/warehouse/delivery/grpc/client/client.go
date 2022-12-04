package client

import (
	"context"

	stdErrors "github.com/pkg/errors"
	"google.golang.org/grpc"

	modelsGlobal "go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/wrapper"
	"go-park-mail-ru/2022_2_BugOverload/internal/warehouse/delivery/grpc/models"
	proto "go-park-mail-ru/2022_2_BugOverload/internal/warehouse/delivery/grpc/protobuf"
)

//go:generate mockgen -source client.go -destination mocks/mockwarehouseclient.go -package mockWarehouseClient

type WarehouseService interface {
	GetRecommendation(ctx context.Context) (modelsGlobal.Film, error)
	GetFilmByID(ctx context.Context, film *modelsGlobal.Film, params *constparams.GetFilmParams) (modelsGlobal.Film, error)
	GetReviewsByFilmID(ctx context.Context, params *constparams.GetFilmReviewsParams) ([]modelsGlobal.Review, error)

	GetStdCollection(ctx context.Context, params *constparams.GetStdCollectionParams) (modelsGlobal.Collection, error)
	GetPremieresCollection(ctx context.Context, params *constparams.GetPremiersCollectionParams) (modelsGlobal.Collection, error)

	GetPersonByID(ctx context.Context, person *modelsGlobal.Person, params *constparams.GetPersonParams) (modelsGlobal.Person, error)

	GetCollectionFilmsAuthorized(ctx context.Context, user *modelsGlobal.User, params *constparams.CollectionGetFilmsRequestParams) (modelsGlobal.Collection, error)
	GetCollectionFilmsNotAuthorized(ctx context.Context, params *constparams.CollectionGetFilmsRequestParams) (modelsGlobal.Collection, error)

	Search(ctx context.Context, params *constparams.SearchParams) (modelsGlobal.Search, error)
}

type WarehouseServiceGRPCClient struct {
	warehouseClient proto.WarehouseServiceClient
}

func NewWarehouseServiceGRPSClient(con *grpc.ClientConn) WarehouseService {
	return &WarehouseServiceGRPCClient{
		warehouseClient: proto.NewWarehouseServiceClient(con),
	}
}

func (c WarehouseServiceGRPCClient) GetRecommendation(ctx context.Context) (modelsGlobal.Film, error) {
	filmProtoResponse, err := c.warehouseClient.GetRecommendation(pkg.GetDefInfoMicroService(ctx), &proto.Nothing{})
	if err != nil {
		return modelsGlobal.Film{}, wrapper.GRPCErrorConvert(stdErrors.Wrap(err, "GetRecommendation"))
	}

	return models.NewFilm(filmProtoResponse), nil
}

func (c WarehouseServiceGRPCClient) GetFilmByID(ctx context.Context, film *modelsGlobal.Film, params *constparams.GetFilmParams) (modelsGlobal.Film, error) {
	filmProtoResponse, err := c.warehouseClient.GetFilmByID(pkg.GetDefInfoMicroService(ctx), models.NewGetFilmParamsProto(film, params))
	if err != nil {
		return modelsGlobal.Film{}, wrapper.GRPCErrorConvert(stdErrors.Wrap(err, "GetFilmByID"))
	}

	return models.NewFilm(filmProtoResponse), nil
}

func (c WarehouseServiceGRPCClient) GetReviewsByFilmID(ctx context.Context, params *constparams.GetFilmReviewsParams) ([]modelsGlobal.Review, error) {
	reviewsProtoResponse, err := c.warehouseClient.GetReviewsByFilmID(pkg.GetDefInfoMicroService(ctx), models.NewGetFilmReviewsParamsProto(params))
	if err != nil {
		return []modelsGlobal.Review{}, wrapper.GRPCErrorConvert(stdErrors.Wrap(err, "GetReviewsByFilmID"))
	}

	return models.NewReviews(reviewsProtoResponse), nil
}

func (c WarehouseServiceGRPCClient) GetStdCollection(ctx context.Context, params *constparams.GetStdCollectionParams) (modelsGlobal.Collection, error) {
	collectionProtoResponse, err := c.warehouseClient.GetStdCollection(pkg.GetDefInfoMicroService(ctx), models.NewGetStdCollectionParamsProto(params))
	if err != nil {
		return modelsGlobal.Collection{}, wrapper.GRPCErrorConvert(stdErrors.Wrap(err, "GetStdCollection"))
	}

	return models.NewCollection(collectionProtoResponse), nil
}

func (c WarehouseServiceGRPCClient) GetPremieresCollection(ctx context.Context, params *constparams.GetPremiersCollectionParams) (modelsGlobal.Collection, error) {
	collectionProtoResponse, err := c.warehouseClient.GetPremieresCollection(pkg.GetDefInfoMicroService(ctx), models.NewPremiersCollectionParamsProto(params))
	if err != nil {
		return modelsGlobal.Collection{}, wrapper.GRPCErrorConvert(stdErrors.Wrap(err, "GetPremieresCollection"))
	}

	return models.NewCollection(collectionProtoResponse), nil
}

func (c WarehouseServiceGRPCClient) GetPersonByID(ctx context.Context, person *modelsGlobal.Person, params *constparams.GetPersonParams) (modelsGlobal.Person, error) {
	personProtoResponse, err := c.warehouseClient.GetPersonByID(pkg.GetDefInfoMicroService(ctx), models.NewGetPersonParamsProto(person, params))
	if err != nil {
		return modelsGlobal.Person{}, wrapper.GRPCErrorConvert(stdErrors.Wrap(err, "GetPersonByID"))
	}

	return models.NewPerson(personProtoResponse), nil
}

func (c WarehouseServiceGRPCClient) GetCollectionFilmsAuthorized(ctx context.Context, user *modelsGlobal.User, params *constparams.CollectionGetFilmsRequestParams) (modelsGlobal.Collection, error) {
	collectionProtoResponse, err := c.warehouseClient.GetCollectionFilmsAuthorized(pkg.GetDefInfoMicroService(ctx), models.NewCollectionGetFilmsAuthParamsProto(user, params))
	if err != nil {
		return modelsGlobal.Collection{}, wrapper.GRPCErrorConvert(stdErrors.Wrap(err, "GetCollectionAuthorized"))
	}

	return models.NewCollection(collectionProtoResponse), nil
}

func (c WarehouseServiceGRPCClient) GetCollectionFilmsNotAuthorized(ctx context.Context, params *constparams.CollectionGetFilmsRequestParams) (modelsGlobal.Collection, error) {
	collectionProtoResponse, err := c.warehouseClient.GetCollectionFilmsNotAuthorized(pkg.GetDefInfoMicroService(ctx), models.NewCollectionGetFilmsNotAuthParamsProto(params))
	if err != nil {
		return modelsGlobal.Collection{}, wrapper.GRPCErrorConvert(stdErrors.Wrap(err, "GetCollectionNotAuthorized"))
	}

	return models.NewCollection(collectionProtoResponse), nil
}

func (c WarehouseServiceGRPCClient) Search(ctx context.Context, params *constparams.SearchParams) (modelsGlobal.Search, error) {
	searchProtoResponse, err := c.warehouseClient.Search(pkg.GetDefInfoMicroService(ctx), models.NewSearchParamsProto(params))
	if err != nil {
		return modelsGlobal.Search{}, wrapper.GRPCErrorConvert(stdErrors.Wrap(err, "Search"))
	}
	return models.NewSearchResponse(searchProtoResponse), nil
}
