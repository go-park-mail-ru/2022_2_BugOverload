package models

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	innerPKG "go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"net/http"
	"strconv"
)

type GetUserCollectionRequest struct {
	Target           string
	Key              string
	SortParam        string
	SortDir          string
	CountCollections int
	Delimiter        string
}

func NewGetUserCollectionRequest() *GetUserCollectionRequest {
	return &GetUserCollectionRequest{}
}

func (p *GetUserCollectionRequest) Bind(r *http.Request) error {
	var err error

	p.SortParam = r.FormValue("sort_param")
	if p.SortParam == "" {
		return errors.ErrBadRequestParamsEmptyRequiredFields
	}

	countCollections := r.FormValue("count_collections")
	if countCollections == "" {
		return errors.ErrBadRequestParamsEmptyRequiredFields
	}

	p.CountCollections, err = strconv.Atoi(countCollections)
	if err != nil {
		return errors.ErrConvertQueryType
	}

	if p.CountCollections < 0 {
		return errors.ErrBadRequestParams
	}

	p.Delimiter = r.FormValue("delimiter")
	if p.Delimiter == "" {
		return errors.ErrBadRequestParamsEmptyRequiredFields
	}

	return nil
}

func (p *GetUserCollectionRequest) GetParams() *innerPKG.GetUserCollectionsParams {
	return &innerPKG.GetUserCollectionsParams{
		SortParam:        p.SortParam,
		CountCollections: p.CountCollections,
		Delimiter:        p.Delimiter,
	}
}

type ShortFilmCollectionResponse struct {
	Name       string `json:"name,omitempty" example:"Популярное"`
	Poster     string `json:"poster,omitempty" example:"42"`
	CountLikes int    `json:"count_likes,omitempty" example:"1023"`
	CountFilms int    `json:"count_films,omitempty"  example:"10"`
}

func NewShortFilmCollectionResponse(collections []models.Collection) []ShortFilmCollectionResponse {
	res := make([]ShortFilmCollectionResponse, len(collections))

	for idx := range collections {
		res[idx] = ShortFilmCollectionResponse{
			Name:       collections[idx].Name,
			Poster:     collections[idx].Poster,
			CountLikes: collections[idx].CountLikes,
			CountFilms: collections[idx].CountFilms,
		}
	}

	return res
}
