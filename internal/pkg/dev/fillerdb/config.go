package fillerdb

type Volume struct {
	CountUser            int  `toml:"count_users"`
	CountViews           int  `toml:"count_views"`
	CountRatings         int  `toml:"count_ratings"`
	CountCollections     int  `toml:"count_collections"`
	CountReviews         int  `toml:"count_reviews"`
	MaxLengthReviewsBody uint `toml:"max_length_review_body"`
	MaxLikesOnReview     int  `toml:"max_likes_on_review"`
	MaxReviewsOnFilm     int  `toml:"max_reviews_on_film"`
	CountReviewsLikes    int  `toml:"count_reviews_likes"`
}

type Database struct {
	URL     string `toml:"URL"`
	Timeout int    `toml:"timeout"`
}

type Config struct {
	Volume   Volume   `toml:"volume"`
	Database Database `toml:"database"`
}

func NewConfig() *Config {
	return &Config{}
}
