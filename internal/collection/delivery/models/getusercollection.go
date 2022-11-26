package models

import (
	innerPKG "go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	"net/http"
	"strconv"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
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
	ID         int    `json:"id,omitempty" example:"12"`
	Name       string `json:"name,omitempty" example:"Избранное"`
	Poster     string `json:"poster,omitempty" example:"42"`
	CountLikes int    `json:"count_likes,omitempty" example:"1023"`
	CountFilms int    `json:"count_films,omitempty"  example:"10"`
	UpdateTime string `json:"update_time,omitempty"  example:"2020.12.12 15:15:15"`
	CreateTime string `json:"create_time,omitempty"  example:"2012.06.05 01:25:00"`
}

func NewShortFilmCollectionResponse(collections []models.Collection) []ShortFilmCollectionResponse {
	res := make([]ShortFilmCollectionResponse, len(collections))

	for idx := range collections {
		res[idx] = ShortFilmCollectionResponse{
			ID:         collections[idx].ID,
			Name:       collections[idx].Name,
			Poster:     collections[idx].Poster,
			CountLikes: collections[idx].CountLikes,
			CountFilms: collections[idx].CountFilms,
			UpdateTime: collections[idx].UpdateTime,
			CreateTime: collections[idx].CreateTime,
		}
	}

	return res
}
