package models

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/security"
)

//go:generate easyjson  -disallow_unknown_fields userprofile.go

type UserProfileRequest struct {
	ID int
}

func NewUserProfileRequest() *UserProfileRequest {
	return &UserProfileRequest{}
}

func (u *UserProfileRequest) Bind(r *http.Request) error {
	vars := mux.Vars(r)

	var err error

	u.ID, err = strconv.Atoi(vars["id"])
	if err != nil {
		return errors.ErrConvertQueryType
	}

	return nil
}

func (u *UserProfileRequest) GetUser() *models.User {
	return &models.User{
		ID: u.ID,
	}
}

//easyjson:json
type UserProfileResponse struct {
	Nickname string `json:"nickname,omitempty" example:"Калыван"`
	Avatar   string `json:"avatar,omitempty" example:"23"`

	CountViewsFilms  int    `json:"count_views_films,omitempty" example:"23"`
	CountCollections int    `json:"count_collections,omitempty" example:"3"`
	CountReviews     int    `json:"count_reviews,omitempty" example:"8"`
	CountRatings     int    `json:"count_ratings,omitempty" example:"20"`
	JoinedDate       string `json:"joined_date,omitempty" example:"2022.10.12"`
}

func NewUserProfileResponse(user *models.User) *UserProfileResponse {
	return &UserProfileResponse{
		Nickname:         security.Sanitize(user.Nickname),
		Avatar:           user.Avatar,
		JoinedDate:       user.JoinedDate,
		CountViewsFilms:  user.CountViewsFilms,
		CountCollections: user.CountCollections,
		CountReviews:     user.CountReviews,
		CountRatings:     user.CountRatings,
	}
}
