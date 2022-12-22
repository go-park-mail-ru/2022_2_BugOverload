package fillerdb

import (
	"database/sql"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	"go-park-mail-ru/2022_2_BugOverload/internal/warehouse/repository/film"
	"time"

	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/sqltools"
)

type CollectionFiller struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Poster      string `json:"poster,omitempty"`
	Public      bool   `json:"public,omitempty"`
}

type PersonFiller struct {
	ID           int      `json:"id,omitempty"`
	Name         string   `json:"name,omitempty"`
	OriginalName string   `json:"original_name,omitempty"`
	Birthday     string   `json:"birthday,omitempty"`
	Avatar       string   `json:"avatar,omitempty"`
	Death        string   `json:"death,omitempty"`
	GrowthMeters float32  `json:"growth_meters,omitempty"`
	Gender       string   `json:"gender,omitempty"`
	Professions  []string `json:"professions,omitempty"`
	Genres       []string `json:"genres,omitempty"`

	Images []string `json:"images,omitempty"`
}

type FilmFiller struct {
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

	BoxOfficeDollars int    `json:"box_office_dollars,omitempty"`
	Budget           int    `json:"budget,omitempty"`
	CurrencyBudget   string `json:"currency_budget,omitempty"`

	CountSeasons int    `json:"count_seasons,omitempty"`
	EndYear      string `json:"end_year,omitempty"`
	Type         string `json:"type,omitempty"`

	Tags          []string `json:"tags,omitempty"`
	Genres        []string `json:"genres,omitempty"`
	ProdCompanies []string `json:"prod_companies,omitempty"`
	ProdCountries []string `json:"prod_countries,omitempty"`
	Images        []string `json:"images,omitempty"`

	Ticket  string `json:"ticket,omitempty"`
	Trailer string `json:"trailer,omitempty"`
}

type FilmSQLFiller struct {
	ID          int
	Name        string
	ProdYear    time.Time
	Description string
	Duration    int

	ShortDescription sql.NullString
	OriginalName     sql.NullString
	Slogan           sql.NullString
	AgeLimit         sql.NullString
	PosterHor        sql.NullString
	PosterVer        sql.NullString

	BoxOffice      sql.NullInt32
	Budget         sql.NullInt32
	CurrencyBudget sql.NullString

	CountSeasons sql.NullInt32
	EndYear      sql.NullTime
	Type         sql.NullString

	Rating               sql.NullFloat64
	CountScores          sql.NullInt32
	CountActors          sql.NullInt32
	CountNegativeReviews sql.NullInt32
	CountNeutralReviews  sql.NullInt32
	CountPositiveReviews sql.NullInt32

	Tags          []string
	Genres        []string
	ProdCompanies []string
	ProdCountries []string
	Images        []string

	Ticket  sql.NullString
	Trailer sql.NullString
}

func NewFilmSQLFillerOnFilm(film FilmFiller) FilmSQLFiller {
	var date time.Time

	var err error

	if len(film.ProdDate) == len(constparams.OnlyDate) {
		date, err = time.Parse(constparams.OnlyDate, film.ProdDate)
		if err != nil {
			date = time.Time{}
		}
	} else {
		date, err = time.Parse(constparams.DateFormat, film.ProdDate)
		if err != nil {
			date = time.Time{}
		}
	}

	return FilmSQLFiller{
		ID:          film.ID,
		Name:        film.Name,
		ProdYear:    date,
		Description: film.Description,
		Duration:    film.DurationMinutes,

		ShortDescription: sqltools.NewSQLNullString(film.ShortDescription),
		OriginalName:     sqltools.NewSQLNullString(film.OriginalName),
		Slogan:           sqltools.NewSQLNullString(film.Slogan),
		AgeLimit:         sqltools.NewSQLNullString(film.AgeLimit),
		PosterHor:        sqltools.NewSQLNullString(film.PosterHor),
		PosterVer:        sqltools.NewSQLNullString(film.PosterVer),

		BoxOffice:      sqltools.NewSQLNullInt32(film.BoxOfficeDollars),
		Budget:         sqltools.NewSQLNullInt32(film.Budget),
		CurrencyBudget: sqltools.NewSQLNullString(film.CurrencyBudget),

		CountSeasons: sqltools.NewSQLNullInt32(film.CountSeasons),
		EndYear:      sqltools.NewSQLNNullDate(film.EndYear, constparams.OnlyDate),
		Type:         sqltools.NewSQLNullString(film.Type),

		Ticket:  sqltools.NewSQLNullString(film.Ticket),
		Trailer: sqltools.NewSQLNullString(film.Trailer),
	}
}

type PersonSQLFiller struct {
	ID       int
	Name     string
	Birthday time.Time
	Growth   float32

	Avatar       sql.NullString
	Gender       sql.NullString
	CountFilms   sql.NullInt32
	OriginalName sql.NullString
	Death        sql.NullTime

	BestFilms []film.ModelSQL

	Images      []string
	Professions []string
	Genres      []string
}

func NewPersonSQLFillerOnPerson(person PersonFiller) PersonSQLFiller {
	birthday := time.Time{}

	if person.Birthday != "" {
		var err error
		birthday, err = time.Parse(constparams.DateFormat, person.Birthday)
		if err != nil {
			birthday = time.Time{}
		}
	}

	return PersonSQLFiller{
		ID:       person.ID,
		Name:     person.Name,
		Birthday: birthday,
		Growth:   person.GrowthMeters,

		Avatar:       sqltools.NewSQLNullString(person.Avatar),
		OriginalName: sqltools.NewSQLNullString(person.OriginalName),
		Gender:       sqltools.NewSQLNullString(person.Gender),
		Death:        sqltools.NewSQLNNullDate(person.Death, constparams.DateFormat),
	}
}

type CollectionSQLFiller struct {
	ID   int
	Name string
	Time string

	Description sql.NullString
	Poster      sql.NullString
	CountLikes  sql.NullInt32
	CountFilms  sql.NullInt32
	Public      bool

	Films []film.ModelSQL
}

func NewCollectionSQLFilmOnCollection(collection CollectionFiller) CollectionSQLFiller {
	return CollectionSQLFiller{
		ID:          collection.ID,
		Name:        collection.Name,
		Description: sqltools.NewSQLNullString(collection.Description),
		Poster:      sqltools.NewSQLNullString(collection.Poster),
		Public:      collection.Public,
	}
}
