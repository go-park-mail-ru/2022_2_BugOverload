package dev

type Volume struct {
	CountUser               int `toml:"count_users"`
	CountProfileViews       int `toml:"count_profile_views"`
	CountProfileRatings     int `toml:"count_profile_ratings"`
	CountProfileCollections int `toml:"count_profile_collections"`
	CountReviews            int `toml:"count_reviews"`
	CountReviewsLikes       int `toml:"count_reviews_likes"`
}

type Database struct {
	URL string `toml:"URL"`
}

type Config struct {
	Volume   Volume   `toml:"volume"`
	Database Database `toml:"database"`
}

func NewConfig() *Config {
	return &Config{}
}
