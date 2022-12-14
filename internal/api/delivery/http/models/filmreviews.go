package models

import (
	innerPKG "go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/security"
)

//go:generate easyjson  -disallow_unknown_fields filmreviews.go

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

	rr.FilmID, _ = strconv.Atoi(vars["id"])

	countReviews := r.FormValue("count_reviews")
	if countReviews == "" {
		return errors.ErrBadRequestParamsEmptyRequiredFields
	}

	rr.CountReviews, err = strconv.Atoi(countReviews)
	if err != nil {
		return errors.ErrConvertQueryType
	}

	if rr.CountReviews <= 0 {
		return errors.ErrBadRequestParams
	}

	offset := r.FormValue("offset")
	if offset == "" {
		return errors.ErrBadRequestParamsEmptyRequiredFields
	}

	rr.Offset, err = strconv.Atoi(offset)
	if err != nil {
		return errors.ErrConvertQueryType
	}

	if rr.Offset < 0 {
		return errors.ErrBadRequestParams
	}

	return nil
}

func (rr *ReviewsRequest) GetParams() *innerPKG.GetFilmReviewsParams {
	return &innerPKG.GetFilmReviewsParams{
		FilmID:       rr.FilmID,
		Offset:       rr.Offset,
		CountReviews: rr.CountReviews,
	}
}

//easyjson:json
type ReviewAuthorResponse struct {
	ID           int    `json:"id,omitempty" example:"54521"`
	Nickname     string `json:"nickname,omitempty" example:"Инокентий"`
	CountReviews int    `json:"count_reviews,omitempty" example:"42"`
	Avatar       string `json:"avatar,omitempty" example:"54521"`
}

//easyjson:json
type ReviewResponse struct {
	Name       string               `json:"name,omitempty" example:"Почему Игра престолов всего лишь одно насилие?"`
	Type       string               `json:"type,omitempty" example:"negative"`
	CreateTime string               `json:"create_time,omitempty" example:"2022-10-30 14:48:48.712860"`
	Body       string               `json:"body,omitempty" example:"Столько крови и убийств нет ни в одном из сериалов, из 730 персонажей больше половины полегло"`
	CountLikes int                  `json:"count_likes,omitempty" example:"142"`
	Author     ReviewAuthorResponse `json:"author,omitempty"`
}

//easyjson:json
type ReviewList []ReviewResponse

func NewReviewsResponse(reviews *[]models.Review) ReviewList {
	res := make([]ReviewResponse, len(*reviews))

	for idx, value := range *reviews {
		res[idx] = ReviewResponse{
			Name:       security.Sanitize(value.Name),
			Type:       value.Type,
			CreateTime: value.CreateTime,
			Body:       security.Sanitize(value.Body),
			CountLikes: value.CountLikes,
			Author: ReviewAuthorResponse{
				ID:           value.Author.ID,
				Nickname:     security.Sanitize(value.Author.Nickname),
				CountReviews: value.Author.CountReviews,
				Avatar:       value.Author.Avatar,
			},
		}
	}

	return res
}
