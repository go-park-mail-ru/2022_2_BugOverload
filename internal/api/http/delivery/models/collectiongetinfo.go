package models

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	innerPKG "go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
)

type CollectionGetFilmsRequest struct {
	CollectionID int
	SortParam    string
}

func NewCollectionGetFilmsRequest() *CollectionGetFilmsRequest {
	return &CollectionGetFilmsRequest{}
}

func (p *CollectionGetFilmsRequest) Bind(r *http.Request) error {
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

func (p *CollectionGetFilmsRequest) GetParams() *innerPKG.CollectionGetFilmsRequestParams {
	return &innerPKG.CollectionGetFilmsRequestParams{
		CollectionID: p.CollectionID,
		SortParam:    p.SortParam,
	}
}

type CollectionGetFilmsResponse struct {
	Name        string                      `json:"name,omitempty" example:"Сейчас в кино"`
	Description string                      `json:"description,omitempty" example:"Фильмы, которые можно посмотреть в российском кинопрокате"`
	Films       []FilmTagCollectionResponse `json:"films,omitempty"`
	IsAuthor    bool                        `json:"is_author,omitempty"`
}

func NewCollectionGetFilmsResponse(collection *models.Collection) *CollectionGetFilmsResponse {
	res := &CollectionGetFilmsResponse{
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

	if collection.Author.ID != 0 {
		res.IsAuthor = true
	}

	return res
}
