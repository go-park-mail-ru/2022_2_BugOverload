package generatordatadb

type UserFace struct {
	ID       int
	Nickname string `faker:"username"`
	Email    string `faker:"email"`
	Password string `faker:"password"`
}

type ProfileFake struct {
	Avatar           string
	JoinedDate       string
	CountViewsFilms  int
	CountCollections int
	CountReviews     int
	CountRatings     int
}

type ReviewFace struct {
	ID   int
	Name string `faker:"word"`
	Type string
	Time string `faker:"timestamp"`
	Body string `faker:"lang=rus"`
}

type FilmFace struct {
	ID               int
	Name             string
	OriginalName     string
	ProdDate         string
	Slogan           string
	ShortDescription string
	Description      string
	AgeLimit         string
	DurationMinutes  int
	PosterHor        string
	PosterVer        string

	Ticket  string
	Trailer string

	BoxOfficeDollars int
	Budget           int
	CurrencyBudget   string

	CountSeasons int
	EndYear      string
	Type         string

	Genres        []string
	ProdCompanies []string
	ProdCountries []string

	Images []string
}
