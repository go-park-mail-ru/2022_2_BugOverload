package models

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	innerPKG "go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/security"
)

type ReviewsRequest struct {
	FilmID       int
	CountReviews int
	Offset       int
}

func NewReviewsRequest() *ReviewsRequest {
	return &ReviewsRequest{}
}

func (rr *ReviewsRequest) Bind(r *http.Request) error {
	if r.Header.Get("Content-Type") != "" {
		return errors.ErrUnsupportedMediaType
	}

	var err error

	vars := mux.Vars(r)

	rr.FilmID, err = strconv.Atoi(vars["id"])
	if err != nil {
		return errors.ErrConvertQueryType
	}

	rr.CountReviews, err = strconv.Atoi(r.FormValue("count_reviews"))
	if err != nil {
		return errors.ErrConvertQueryType
	}

	if rr.CountReviews <= 0 {
		return errors.ErrBadQueryParams
	}

	rr.Offset, err = strconv.Atoi(r.FormValue("offset"))
	if err != nil {
		return errors.ErrConvertQueryType
	}

	if rr.Offset < 0 {
		return errors.ErrBadQueryParams
	}

	return nil
}

func (rr *ReviewsRequest) GetParams() *innerPKG.GetReviewsFilmParams {
	return &innerPKG.GetReviewsFilmParams{
		FilmID: rr.FilmID,
		Offset: rr.Offset,
		Count:  rr.CountReviews,
	}
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
	CreateTime string               `json:"create_time,omitempty" example:"2022-10-30 14:48:48.712860"`
	Body       string               `json:"body,omitempty" example:"Столько крови и убийств нет ни в одном из сериалов, из 730 персонажей больше половины полегло"`
	CountLikes int                  `json:"count_likes,omitempty" example:"142"`
	Author     ReviewAuthorResponse `json:"author,omitempty"`
}

func NewReviewsResponse(reviews *[]models.Review) []*ReviewResponse {
	res := make([]*ReviewResponse, len(*reviews))

	for idx, value := range *reviews {
		res[idx] = &ReviewResponse{
			Name:       security.Sanitize(value.Name),
			Type:       value.Type,
			CreateTime: value.CreateTime,
			Body:       security.Sanitize(value.Body),
			CountLikes: value.CountLikes,
			Author: ReviewAuthorResponse{
				ID:           value.Author.ID,
				Nickname:     security.Sanitize(value.Author.Nickname),
				CountReviews: value.Author.Profile.CountReviews,
				Avatar:       value.Author.Profile.Avatar,
			},
		}
	}

	return res
}
