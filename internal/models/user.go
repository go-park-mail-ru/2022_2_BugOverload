package models

//go:generate easyjson -all -disallow_unknown_fields user.go

type User struct {
	ID               int    `json:"user_id,omitempty"`
	Nickname         string `json:"nickname,omitempty"`
	Email            string `json:"email,omitempty"`
	Password         string `json:"password,omitempty"`
	IsAdmin          bool   `json:"-"`
	Avatar           string `json:"avatar,omitempty"`
	JoinedDate       string `json:"joined_date,omitempty"`
	CountViewsFilms  int    `json:"count_views_films,omitempty"`
	CountCollections int    `json:"count_collections,omitempty"`
	CountReviews     int    `json:"count_reviews,omitempty"`
	CountRatings     int    `json:"count_ratings,omitempty"`
}

type NodeInUserCollection struct {
	ID     int    `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	IsUsed bool   `json:"is_used,omitempty"`
}

type UserActivity struct {
	CountReviews int                    `json:"count_reviews,omitempty"`
	Rating       int                    `json:"rating,omitempty"`
	DateRating   string                 `json:"date_rating,omitempty"`
	Collections  []NodeInUserCollection `json:"collections,omitempty"`
}
