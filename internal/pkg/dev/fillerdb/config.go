package fillerdb

type Volume struct {
	CountUser               int  `toml:"count_users"`
	CountProfileViews       int  `toml:"count_views"`
	CountProfileRatings     int  `toml:"count_ratings"`
	CountProfileCollections int  `toml:"count_collections"`
	CountReviews            int  `toml:"count_reviews"`
	MaxLengthReviewsBody    uint `toml:"max_length_review_body"`
	CountReviewsLikes       int  `toml:"count_reviews_likes"`
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
