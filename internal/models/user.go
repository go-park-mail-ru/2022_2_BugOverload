package models

type User struct {
	ID       uint    `json:"user_id,omitempty"`
	Nickname string  `json:"nickname,omitempty"`
	Email    string  `json:"email,omitempty"`
	Password string  `json:"password,omitempty"`
	Profile  Profile `json:"profile,omitempty"`
}

type Profile struct {
	Avatar           string `json:"avatar,omitempty"`
	JoinedDate       string `json:"joined_date,omitempty"`
	CountViewsFilms  int    `json:"count_views_films,omitempty"`
	CountCollections int    `json:"count_collections,omitempty"`
	CountReviews     int    `json:"count_reviews,omitempty"`
	CountRatings     int    `json:"count_ratings,omitempty"`
}
