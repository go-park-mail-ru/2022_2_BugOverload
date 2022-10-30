package models

import (
	"context"
	"github.com/gorilla/mux"
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	innerPKG "go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"net/http"
	"strconv"
)

type GetReviewsParams struct {
	Count     int
	Delimiter string
}

type ReviewsRequest struct {
	FilmID int
	Params GetReviewsParams
}

func NewReviewsRequest() *ReviewsRequest {
	return &ReviewsRequest{}
}

func (rr *ReviewsRequest) Bind(r *http.Request) error {
	if r.Header.Get("Content-Type") != "" {
		return errors.NewErrValidation(errors.ErrUnsupportedMediaType)
	}

	var err error

	vars := mux.Vars(r)
	rr.FilmID, err = strconv.Atoi(vars["id"])
	if err != nil {
		return errors.NewErrValidation(errors.ErrUnsupportedMediaType)
	}

	rr.Params.Count, err = strconv.Atoi(r.FormValue("count"))
	if err != nil {
		return errors.NewErrValidation(errors.ErrUnsupportedMediaType)
	}

	rr.Params.Delimiter = r.FormValue("delimiter")

	return nil
}

func (rr *ReviewsRequest) GetParams(ctx context.Context) (*models.Film, context.Context) {
	ctx = context.WithValue(ctx, innerPKG.GetReviewsRequestParams, rr.Params)

	filmParams := &models.Film{
		ID: rr.FilmID,
	}

	return filmParams, ctx
}

type ReviewAuthorResponse struct {
	ID           int    `json:"id,omitempty" example:"54521"`
	Nickname     string `json:"nickname,omitempty" example:"Инокентий"`
	CountReviews int    `json:"count_reviews,omitempty" example:"42"`
	Avatar       string `json:"avatar,omitempty" example:"54521"`
}

type ReviewResponse struct {
	Name       string               `json:"name,omitempty" example:"Почему Игра престолов всего лишь одно насилие?"`
	Type       string               `json:"type,omitempty" example:"negative"`
	Time       string               `json:"time,omitempty" example:"2022-10-30 14:48:48.712860"`
	Body       string               `json:"body,omitempty" example:"Столько крови и убийств нет ни в одном из сериалов, из 730 персонажей больше половины полегло"`
	CountLikes int                  `json:"count_likes,omitempty" example:"142"`
	Author     ReviewAuthorResponse `json:"author,omitempty"`
}

func NewReviewsResponse(reviews []*models.Review) []*ReviewResponse {
	res := make([]*ReviewResponse, len(reviews))

	for idx, value := range reviews {
		res[idx] = &ReviewResponse{
			Name:       value.Name,
			Type:       value.Type,
			Time:       value.Time,
			Body:       value.Body,
			CountLikes: value.CountLikes,
			Author: ReviewAuthorResponse{
				ID:           value.Author.ID,
				Nickname:     value.Author.Nickname,
				CountReviews: value.Author.Profile.CountReviews,
				Avatar:       value.Author.Profile.Avatar,
			},
		}
	}

	return res
}
