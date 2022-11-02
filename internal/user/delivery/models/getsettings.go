package models

import (
	"context"
	"net/http"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
)

type UserGetSettingsRequest struct{}

func NewGetUserSettingsRequest() *UserGetSettingsRequest {
	return &UserGetSettingsRequest{}
}

func (u *UserGetSettingsRequest) Bind(r *http.Request) (context.Context, error) {
	if r.Header.Get("Cookie") == "" {
		return nil, errors.NewErrAuth(errors.ErrNoCookie)
	}

	cookie := r.Cookies()[0]

	ctx := context.WithValue(r.Context(), pkg.SessionKey, cookie.Value)

	return ctx, nil
}

type GetUserSettingsResponse struct {
	CountViewsFilms  int    `json:"count_views_films,omitempty" example:"23"`
	CountCollections int    `json:"count_collections,omitempty" example:"3"`
	CountReviews     int    `json:"count_reviews,omitempty" example:"8"`
	CountRatings     int    `json:"count_ratings,omitempty" example:"20"`
	JoinedDate       string `json:"joined_date,omitempty" example:"2022-10-12"`
}

func NewGetUserSettingsResponse(user *models.User) *GetUserSettingsResponse {
	return &GetUserSettingsResponse{
		JoinedDate:       user.Profile.JoinedDate,
		CountViewsFilms:  user.Profile.CountViewsFilms,
		CountCollections: user.Profile.CountCollections,
		CountReviews:     user.Profile.CountReviews,
		CountRatings:     user.Profile.CountRatings,
	}
}
