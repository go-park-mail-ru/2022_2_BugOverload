package fillerdb

import innerPKG "go-park-mail-ru/2022_2_BugOverload/internal/pkg"

type Volume struct {
	CountUser int `toml:"count_users"`

	CountRatings          int `toml:"count_ratings"`
	MaxRatings            int `toml:"max_rating"`
	MaxCountRatingsOnFilm int `toml:"max_count_ratings_on_film"`

	CountViews    int `toml:"count_views"`
	MaxViewOnFilm int `toml:"max_views_on_film"`

	CountReviews         int `toml:"count_reviews"`
	MaxLengthReviewsBody int `toml:"max_length_review_body"`
	CountReviewsLikes    int `toml:"count_reviews_likes"`
	MaxLikesOnReview     int `toml:"max_likes_on_review"`
	MaxReviewsOnFilm     int `toml:"max_reviews_on_film"`

	TypeFilmPersonLinks string `toml:"type_film_persons_links"`
	MaxFilmActors       int    `toml:"max_film_actors"`
	MaxFilmPersons      int    `toml:"max_film_persons"`
	MaxFilmsInTag       int    `toml:"max_films_in_tag"`

	CountCollections int `toml:"count_collections"`
}

type Database struct {
	Timeout int `toml:"timeout"`
}

type Config struct {
	Volume         Volume                  `toml:"volume"`
	Database       Database                `toml:"database"`
	DatabaseParams innerPKG.DatabaseParams `toml:"database_params"`
}

func NewConfig() *Config {
	return &Config{}
}
