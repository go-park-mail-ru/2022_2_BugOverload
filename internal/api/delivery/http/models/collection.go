package models

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	innerPKG "go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
)

//go:generate easyjson -omit_empty -disallow_unknown_fields collection.go

type CollectionRequest struct {
	CollectionID int
	SortParam    string
}

func NewCollectionRequest() *CollectionRequest {
	return &CollectionRequest{}
}

func (p *CollectionRequest) Bind(r *http.Request) error {
	var err error
	vars := mux.Vars(r)

	p.CollectionID, err = strconv.Atoi(vars["id"])
	if err != nil {
		return errors.ErrConvertQueryType
	}

	p.SortParam = r.FormValue("sort_param")
	if p.SortParam == "" {
		return errors.ErrBadRequestParamsEmptyRequiredFields
	}

	return nil
}

func (p *CollectionRequest) GetParams() *innerPKG.CollectionGetFilmsRequestParams {
	return &innerPKG.CollectionGetFilmsRequestParams{
		CollectionID: p.CollectionID,
		SortParam:    p.SortParam,
	}
}

//easyjson:json
type CollectionAuthorResponse struct {
	ID               int    `json:"id,omitempty" example:"54521"`
	Nickname         string `json:"nickname,omitempty" example:"Инокентий"`
	CountCollections int    `json:"count_collections,omitempty" example:"42"`
	Avatar           string `json:"avatar,omitempty" example:"54521"`
}

//easyjson:json
type CollectionResponse struct {
	Name        string                      `json:"name,omitempty" example:"Сейчас в кино"`
	Description string                      `json:"description,omitempty" example:"Фильмы, которые можно посмотреть в российском кинопрокате"`
	Films       []FilmTagCollectionResponse `json:"films,omitempty"`
	IsAuthor    bool                        `json:"is_author,omitempty" example:"true"`
	Author      *CollectionAuthorResponse   `json:"author"`
}

func NewCollectionResponse(collection *models.Collection) *CollectionResponse {
	res := &CollectionResponse{
		Name:        collection.Name,
		Description: collection.Description,
		Films:       make([]FilmTagCollectionResponse, len(collection.Films)),
	}

	for idx, value := range collection.Films {
		res.Films[idx] = FilmTagCollectionResponse{
			ID:        value.ID,
			Name:      value.Name,
			ProdYear:  value.ProdDate,
			EndYear:   value.EndYear,
			PosterVer: value.PosterVer,
			Rating:    value.Rating,
			Genres:    value.Genres,
		}
	}

	if collection.Author.ID == 0 {
		res.IsAuthor = true

		return res
	}

	res.Author = &CollectionAuthorResponse{
		ID:               collection.Author.ID,
		Nickname:         collection.Author.Nickname,
		Avatar:           collection.Author.Avatar,
		CountCollections: collection.Author.CountCollections,
	}

	return res
}
