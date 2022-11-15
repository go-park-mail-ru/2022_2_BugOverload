package fillerdb

import (
	"database/sql"
	"time"

	"go-park-mail-ru/2022_2_BugOverload/internal/film/repository"
	innerPKG "go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/sqltools"
)

type CollectionFiller struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type PersonFiller struct {
	ID           int      `json:"id,omitempty"`
	Name         string   `json:"name,omitempty"`
	OriginalName string   `json:"original_name,omitempty"`
	Birthday     string   `json:"birthday,omitempty"`
	Avatar       string   `json:"avatar,omitempty"`
	Death        string   `json:"death,omitempty"`
	Growth       float32  `json:"growth,omitempty"`
	Gender       string   `json:"gender,omitempty"`
	Professions  []string `json:"professions,omitempty"`
	Genres       []string `json:"genres,omitempty"`

	Images []string `json:"images,omitempty"`
}

type FilmFiller struct {
	ID               int    `json:"id,omitempty"`
	Name             string `json:"name,omitempty"`
	OriginalName     string `json:"original_name,omitempty"`
	ProdYear         string `json:"prod_year,omitempty"`
	Slogan           string `json:"slogan,omitempty"`
	ShortDescription string `json:"short_description,omitempty"`
	Description      string `json:"description,omitempty"`
	AgeLimit         string `json:"age_limit,omitempty"`
	Duration         int    `json:"duration,omitempty"`
	PosterHor        string `json:"poster_hor,omitempty"`
	PosterVer        string `json:"poster_ver,omitempty"`

	BoxOffice      int    `json:"box_office,omitempty"`
	Budget         int    `json:"budget,omitempty"`
	CurrencyBudget string `json:"currency_budget,omitempty"`

	CountSeasons int    `json:"count_seasons,omitempty"`
	EndYear      string `json:"end_year,omitempty"`
	Type         string `json:"type,omitempty"`

	Tags          []string `json:"tags,omitempty"`
	Genres        []string `json:"genres,omitempty"`
	ProdCompanies []string `json:"prod_companies,omitempty"`
	ProdCountries []string `json:"prod_countries,omitempty"`
	Images        []string `json:"images,omitempty"`
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
}

func NewFilmSQLFillerOnFilm(film FilmFiller) FilmSQLFiller {
	date, err := time.Parse(innerPKG.OnlyDate, film.ProdYear)
	if err != nil {
		date = time.Time{}
	}

	return FilmSQLFiller{
		ID:          film.ID,
		Name:        film.Name,
		ProdYear:    date,
		Description: film.Description,
		Duration:    film.Duration,

		ShortDescription: sqltools.NewSQLNullString(film.ShortDescription),
		OriginalName:     sqltools.NewSQLNullString(film.OriginalName),
		Slogan:           sqltools.NewSQLNullString(film.Slogan),
		AgeLimit:         sqltools.NewSQLNullString(film.AgeLimit),
		PosterHor:        sqltools.NewSQLNullString(film.PosterHor),
		PosterVer:        sqltools.NewSQLNullString(film.PosterVer),

		BoxOffice:      sqltools.NewSQLNullInt32(film.BoxOffice),
		Budget:         sqltools.NewSQLNullInt32(film.Budget),
		CurrencyBudget: sqltools.NewSQLNullString(film.CurrencyBudget),

		CountSeasons: sqltools.NewSQLNullInt32(film.CountSeasons),
		EndYear:      sqltools.NewSQLNNullDate(film.EndYear, innerPKG.OnlyDate),
		Type:         sqltools.NewSQLNullString(film.Type),
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

	BestFilms []repository.FilmSQL

	Images      []string
	Professions []string
	Genres      []string
}

func NewPersonSQLFillerOnPerson(person PersonFiller) PersonSQLFiller {
	birthday := time.Time{}

	if person.Birthday != "" {
		var err error
		birthday, err = time.Parse(innerPKG.DateFormat, person.Birthday)
		if err != nil {
			birthday = time.Time{}
		}
	}

	return PersonSQLFiller{
		ID:       person.ID,
		Name:     person.Name,
		Birthday: birthday,
		Growth:   person.Growth,

		Avatar:       sqltools.NewSQLNullString(person.Avatar),
		OriginalName: sqltools.NewSQLNullString(person.OriginalName),
		Gender:       sqltools.NewSQLNullString(person.Gender),
		Death:        sqltools.NewSQLNNullDate(person.Death, innerPKG.DateFormat),
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

	Films []repository.FilmSQL
}

func NewCollectionSQLFilmOnCollection(collection CollectionFiller) CollectionSQLFiller {
	return CollectionSQLFiller{
		ID:          collection.ID,
		Name:        collection.Name,
		Description: sqltools.NewSQLNullString(collection.Description),
	}
}
