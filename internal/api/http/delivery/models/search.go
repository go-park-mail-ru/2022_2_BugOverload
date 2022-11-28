package models

import (
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
