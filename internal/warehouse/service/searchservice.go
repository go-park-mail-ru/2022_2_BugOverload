package service

import (
	"context"

	stdErrors "github.com/pkg/errors"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	"go-park-mail-ru/2022_2_BugOverload/internal/warehouse/repository/search"
)

//go:generate mockgen -source searchservice.go -destination mocks/mocksearchservice.go -package mockWarehouseService

// SearchService provides universal service for work with search.
type SearchService interface {
	Search(ctx context.Context, params *constparams.SearchParams) (models.Search, error)
}

// searchService is implementation for users service corresponding to the SearchService interface.
type searchService struct {
	searchRepo search.Repository
}

// NewSearchService is constructor for searchService.
func NewSearchService(sr search.Repository) SearchService {
	return &searchService{
		searchRepo: sr,
	}
}

func (s *searchService) Search(ctx context.Context, params *constparams.SearchParams) (models.Search, error) {
	var searchResponseRepo models.Search
	var err error

	params.Query = "%" + params.Query + "%"

	searchResponseRepo.Films, err = s.searchRepo.SearchFilms(ctx, params)
	if err != nil {
		return models.Search{}, stdErrors.Wrap(err, "SearchFilms")
	}

	searchResponseRepo.Serials, err = s.searchRepo.SearchSeries(ctx, params)
	if err != nil {
		return models.Search{}, stdErrors.Wrap(err, "SearchSeries")
	}

	searchResponseRepo.Persons, err = s.searchRepo.SearchPersons(ctx, params)
	if err != nil {
		return models.Search{}, stdErrors.Wrap(err, "SearchPersons")
	}

	return searchResponseRepo, nil
}
