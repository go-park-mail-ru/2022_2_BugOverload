package models

//go:generate easyjson -all -disallow_unknown_fields film.go

type FilmActor struct {
	ID        int    `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Avatar    string `json:"avatar,omitempty"`
	Character string `json:"character,omitempty"`
}

type FilmPerson struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Film struct {
	ID               int    `json:"id,omitempty"`
	Name             string `json:"name,omitempty"`
	OriginalName     string `json:"original_name,omitempty"`
	ProdDate         string `json:"prod_date,omitempty"`
	Slogan           string `json:"slogan,omitempty"`
	ShortDescription string `json:"short_description,omitempty"`
	Description      string `json:"description,omitempty"`
	AgeLimit         string `json:"age_limit,omitempty"`
	DurationMinutes  int    `json:"duration_minutes,omitempty"`
	PosterHor        string `json:"poster_hor,omitempty"`
	PosterVer        string `json:"poster_ver,omitempty"`
	Ticket           string `json:"ticket,omitempty"`
	Trailer          string `json:"trailer,omitempty"`

	BoxOfficeDollars int    `json:"box_office_dollars,omitempty"`
	Budget           int    `json:"budget,omitempty"`
	CurrencyBudget   string `json:"currency_budget,omitempty"`

	CountSeasons int    `json:"count_seasons,omitempty"`
	EndYear      string `json:"end_year,omitempty"`
	Type         string `json:"type,omitempty"`

	Rating               float32 `json:"rating,omitempty"`
	CountRatings         int     `json:"count_ratings,omitempty"`
	CountActors          int     `json:"count_actors,omitempty"`
	CountNegativeReviews int     `json:"count_negative_reviews,omitempty"`
	CountNeutralReviews  int     `json:"count_neutral_reviews,omitempty"`
	CountPositiveReviews int     `json:"count_positive_reviews,omitempty"`

	Tags          []string     `json:"tags,omitempty"`
	Genres        []string     `json:"genres,omitempty"`
	ProdCompanies []string     `json:"prod_companies,omitempty"`
	ProdCountries []string     `json:"prod_countries,omitempty"`
	Actors        []FilmActor  `json:"actors,omitempty"`
	Artists       []FilmPerson `json:"artists,omitempty"`
	Directors     []FilmPerson `json:"directors,omitempty"`
	Writers       []FilmPerson `json:"writers,omitempty"`
	Producers     []FilmPerson `json:"producers,omitempty"`
	Operators     []FilmPerson `json:"operators,omitempty"`
	Montage       []FilmPerson `json:"montage,omitempty"`
	Composers     []FilmPerson `json:"composers,omitempty"`

	Images []string `json:"images,omitempty"`
}
