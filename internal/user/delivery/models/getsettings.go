package models

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
)

type GetUserSettingsResponse struct {
	CountViewsFilms  int    `json:"count_views_films,omitempty" example:"23"`
	CountCollections int    `json:"count_collections,omitempty" example:"3"`
	CountReviews     int    `json:"count_reviews,omitempty" example:"8"`
	CountRatings     int    `json:"count_ratings,omitempty" example:"20"`
	JoinedDate       string `json:"joined_date,omitempty" example:"2022.10.12"`
}

func NewGetUserSettingsResponse(user *models.User) *GetUserSettingsResponse {
	return &GetUserSettingsResponse{
		JoinedDate:       user.JoinedDate,
		CountViewsFilms:  user.CountViewsFilms,
		CountCollections: user.CountCollections,
		CountReviews:     user.CountReviews,
		CountRatings:     user.CountRatings,
	}
}
