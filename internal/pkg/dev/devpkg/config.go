package devpkg

import innerPKG "go-park-mail-ru/2022_2_BugOverload/internal/pkg"

type Volume struct {
	// Users
	CountUser int `toml:"count_users"`

	// Films
	TypeFilms                     string `toml:"type_film"`
	CountFilms                    int    `toml:"count_films"`
	MaxFilmsNameLength            int    `toml:"max_films_name_length"`
	MaxFilmDescriptionLength      int    `toml:"max_films_description_length"`
	MaxFilmShortDescriptionLength int    `toml:"max_films_short_description_length"`
	MaxFilmsOriginalNameLength    int    `toml:"max_films_original_name_length"`
	MaxFilmsSloganLength          int    `toml:"max_films_slogan_length"`
	MaxFilmsPosterHorLength       int    `toml:"max_films_poster_hor_length"`
	MaxFilmsPosterVerLength       int    `toml:"max_films_poster_ver_length"`
	MaxFilmsBoxOfficeDollars      int    `toml:"max_films_box_office_dollars"`
	MaxFilmsBudget                int    `toml:"max_films_budget"`
	MaxFilmsDurationMinutes       int    `toml:"max_films_duration_minutes"`
	MaxFilmsCountSeasons          int    `toml:"max_films_count_seasons"`
	MaxFilmsTicketLength          int    `toml:"max_films_ticket_length"`
	MaxFilmsTrailerLength         int    `toml:"max_films_trailer_length"`
	MaxFilmsProdCountriesCount    int    `toml:"max_films_prod_countries_count"`
	MaxFilmsProdCompaniesCount    int    `toml:"max_films_prod_companies_count"`
	MaxFilmsPGenresCount          int    `toml:"max_films_genres_fields"`

	CountRatings          int `toml:"count_ratings"`
	MaxRatings            int `toml:"max_rating"`
	MaxCountRatingsOnFilm int `toml:"max_count_ratings_on_film"`

	CountViews    int `toml:"count_views"`
	MaxViewOnFilm int `toml:"max_views_on_film"`

	// Reviews
	CountReviews         int `toml:"count_reviews"`
	MaxLengthReviewsBody int `toml:"max_length_review_body"`
	CountReviewsLikes    int `toml:"count_reviews_likes"`
	MaxLikesOnReview     int `toml:"max_likes_on_review"`
	MaxReviewsOnFilm     int `toml:"max_reviews_on_film"`

	// FilmLinkPersons
	TypeFilmsPersonLinks string `toml:"type_film_persons_links"`
	MaxFilmsActors       int    `toml:"max_film_actors"`
	MaxFilmsPersons      int    `toml:"max_film_persons"`
	MaxFilmsInTag        int    `toml:"max_films_in_tag"`
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
