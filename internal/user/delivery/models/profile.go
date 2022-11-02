package models

import (
	"net/http"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
)

type UserProfileRequest struct{}

func NewUserProfileRequest() *UserProfileRequest {
	return &UserProfileRequest{}
}

func (u *UserProfileRequest) Bind(r *http.Request) error {
	if r.Header.Get("Cookie") == "" {
		return errors.NewErrAuth(errors.ErrNoCookie)
	}

	return nil
}

type UserProfileResponse struct {
	Nickname string `json:"nickname,omitempty" example:"Калыван"`
	Email    string `json:"email,omitempty" example:"iamturckishbaby@gmail.com"`
	Avatar   string `json:"avatar,omitempty" example:"23"`

	CountViewsFilms  int    `json:"count_views_films,omitempty" example:"23"`
	CountCollections int    `json:"count_collections,omitempty" example:"3"`
	CountReviews     int    `json:"count_reviews,omitempty" example:"8"`
	CountRatings     int    `json:"count_ratings,omitempty" example:"20"`
	JoinedDate       string `json:"joined_date,omitempty" example:"2022-10-12"`
}

func NewUserProfileResponse(user *models.User) *UserProfileResponse {
	return &UserProfileResponse{
		Nickname:         user.Nickname,
		Email:            user.Email,
		Avatar:           user.Profile.Avatar,
		JoinedDate:       user.Profile.JoinedDate,
		CountViewsFilms:  user.Profile.CountViewsFilms,
		CountCollections: user.Profile.CountCollections,
		CountReviews:     user.Profile.CountReviews,
		CountRatings:     user.Profile.CountRatings,
	}
}
