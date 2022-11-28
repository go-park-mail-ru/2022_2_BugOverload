package models

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"net/http"

	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
)

type SearchRequest struct {
	query string
}

func NewSearchRequest() *SearchRequest {
	return &SearchRequest{}
}

func (p *SearchRequest) Bind(r *http.Request) error {
	p.query = r.FormValue("q")
	if p.query == "" {
		return errors.ErrQueryRequiredEmpty
	}
	return nil
}

func (p *SearchRequest) GetParams() *constparams.SearchParams {
	return &constparams.SearchParams{
		Query: p.query,
	}
}

func NewSearchResponse(resp models.SearchResponse) (models.SearchResponse, error) {
	if len(resp.Films) == 0 && len(resp.Series) == 0 && len(resp.Persons) == 0 {
		return models.SearchResponse{}, errors.ErrNotFoundInDB
	}
	return resp, nil
}
