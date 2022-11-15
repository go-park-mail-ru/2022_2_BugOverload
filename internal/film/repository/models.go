package repository

import (
	"context"
	"database/sql"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	innerPKG "go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/sqltools"
)

type FilmActorSQL struct {
	ID        int
	Name      string
	Avatar    sql.NullString
	Character string
}

type FilmPersonSQL struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type FilmSQL struct {
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

	Actors    []FilmActorSQL
	Artists   []FilmPersonSQL
	Directors []FilmPersonSQL
	Writers   []FilmPersonSQL
	Producers []FilmPersonSQL
	Operators []FilmPersonSQL
	Montage   []FilmPersonSQL
	Composers []FilmPersonSQL
}

func NewFilmSQL() FilmSQL {
	return FilmSQL{}
}

func (f *FilmSQL) Convert() models.Film {
	res := models.Film{
		ID:          f.ID,
		Name:        f.Name,
		ProdYear:    f.ProdYear.Format(innerPKG.OnlyDate),
		Description: f.Description,
		Duration:    f.Duration,

		ShortDescription: f.ShortDescription.String,
		OriginalName:     f.OriginalName.String,
		Slogan:           f.Slogan.String,
		AgeLimit:         f.AgeLimit.String,
		PosterHor:        f.PosterHor.String,
		PosterVer:        f.PosterVer.String,

		BoxOffice:      int(f.BoxOffice.Int32),
		Budget:         int(f.Budget.Int32),
		CurrencyBudget: f.CurrencyBudget.String,

		CountSeasons: int(f.CountSeasons.Int32),
		EndYear:      f.EndYear.Time.Format(innerPKG.OnlyDate),
		Type:         f.Type.String,

		Rating:               float32(f.Rating.Float64),
		CountScores:          int(f.CountScores.Int32),
		CountActors:          int(f.CountActors.Int32),
		CountNegativeReviews: int(f.CountNegativeReviews.Int32),
		CountNeutralReviews:  int(f.CountNeutralReviews.Int32),
		CountPositiveReviews: int(f.CountPositiveReviews.Int32),

		Genres:        f.Genres,
		Tags:          f.Tags,
		ProdCountries: f.ProdCountries,
		ProdCompanies: f.ProdCompanies,
		Images:        f.Images,

		Actors:    make([]models.FilmActor, len(f.Actors)),
		Artists:   make([]models.FilmPerson, len(f.Artists)),
		Directors: make([]models.FilmPerson, len(f.Directors)),
		Writers:   make([]models.FilmPerson, len(f.Writers)),
		Producers: make([]models.FilmPerson, len(f.Producers)),
		Operators: make([]models.FilmPerson, len(f.Operators)),
		Montage:   make([]models.FilmPerson, len(f.Montage)),
		Composers: make([]models.FilmPerson, len(f.Composers)),
	}

	for idx, value := range f.Actors {
		res.Actors[idx].ID = value.ID
		res.Actors[idx].Name = value.Name
		res.Actors[idx].Character = value.Character
		res.Actors[idx].Avatar = value.Avatar.String
	}

	for idx, value := range f.Artists {
		res.Artists[idx].ID = value.ID
		res.Artists[idx].Name = value.Name
	}

	for idx, value := range f.Directors {
		res.Directors[idx].ID = value.ID
		res.Directors[idx].Name = value.Name
	}

	for idx, value := range f.Writers {
		res.Writers[idx].ID = value.ID
		res.Writers[idx].Name = value.Name
	}

	for idx, value := range f.Producers {
		res.Producers[idx].ID = value.ID
		res.Producers[idx].Name = value.Name
	}

	for idx, value := range f.Operators {
		res.Operators[idx].ID = value.ID
		res.Operators[idx].Name = value.Name
	}

	for idx, value := range f.Artists {
		res.Artists[idx].ID = value.ID
		res.Artists[idx].Name = value.Name
	}

	for idx, value := range f.Composers {
		res.Composers[idx].ID = value.ID
		res.Composers[idx].Name = value.Name
	}

	return res
}

func (f *FilmSQL) GetMainInfo(ctx context.Context, db *sql.DB, query string, args ...any) error {
	return sqltools.RunQuery(ctx, db, func(ctx context.Context, conn *sql.Conn) error {
		rowFilm := conn.QueryRowContext(ctx, query, args...)
		if rowFilm.Err() != nil {
			return rowFilm.Err()
		}

		err := rowFilm.Scan(
			&f.Name,
			&f.OriginalName,
			&f.ProdYear,
			&f.Slogan,
			&f.Description,
			&f.ShortDescription,
			&f.AgeLimit,
			&f.Duration,
			&f.PosterHor,
			&f.Budget,
			&f.BoxOffice,
			&f.CurrencyBudget,
			&f.CountSeasons,
			&f.EndYear,
			&f.Type,
			&f.Rating,
			&f.CountActors,
			&f.CountScores,
			&f.CountNegativeReviews,
			&f.CountNeutralReviews,
			&f.CountPositiveReviews)
		if err != nil {
			return err
		}

		if !f.Type.Valid {
			f.Type.String = innerPKG.DefTypeFilm
		}

		if !f.PosterHor.Valid {
			f.PosterHor.String = innerPKG.DefFilmPosterHor
		}

		return nil
	})
}

func (f *FilmSQL) GetPersons(ctx context.Context, db *sql.DB, query string, args ...any) error {
	return sqltools.RunQuery(ctx, db, func(ctx context.Context, conn *sql.Conn) error {
		rowsFilmActors, err := conn.QueryContext(ctx, query, args...)
		if err != nil {
			return err
		}
		defer rowsFilmActors.Close()

		for rowsFilmActors.Next() {
			var person FilmPersonSQL
			var professionID int

			err = rowsFilmActors.Scan(
				&person.ID,
				&person.Name,
				&professionID)
			if err != nil {
				return err
			}

			switch professionID {
			case innerPKG.Artist:
				f.Artists = append(f.Artists, person)
			case innerPKG.Director:
				f.Directors = append(f.Directors, person)
			case innerPKG.Writer:
				f.Writers = append(f.Writers, person)
			case innerPKG.Producer:
				f.Producers = append(f.Producers, person)
			case innerPKG.Operator:
				f.Operators = append(f.Operators, person)
			case innerPKG.Montage:
				f.Montage = append(f.Montage, person)
			case innerPKG.Composer:
				f.Composers = append(f.Composers, person)
			}
		}

		return nil
	})
}

func (f *FilmSQL) GetActors(ctx context.Context, db *sql.DB, query string, args ...any) error {
	return sqltools.RunQuery(ctx, db, func(ctx context.Context, conn *sql.Conn) error {
		rowsFilmActors, err := conn.QueryContext(ctx, query, args...)
		if err != nil {
			return err
		}
		defer rowsFilmActors.Close()

		for rowsFilmActors.Next() {
			var actor FilmActorSQL

			err = rowsFilmActors.Scan(
				&actor.ID,
				&actor.Name,
				&actor.Avatar,
				&actor.Character)
			if err != nil {
				return err
			}

			f.Actors = append(f.Actors, actor)
		}

		return nil
	})
}

const (
	getGenresFilmBatchBegin = `
SELECT f.film_id,
       g.name
FROM genres g
         JOIN film_genres fg ON g.genre_id = fg.fk_genre_id
         JOIN films f ON fg.fk_film_id = f.film_id
WHERE f.film_id IN (`

	getGenresFilmBatchEnd = `) ORDER BY f.film_id, fg.weight DESC`
)

func GetGenresBatch(ctx context.Context, target []FilmSQL, conn *sql.Conn) ([]FilmSQL, error) {
	setID := make([]string, len(target))

	mapFilms := make(map[int]int, len(target))

	for idx := range target {
		setID[idx] = strconv.Itoa(target[idx].ID)

		mapFilms[target[idx].ID] = idx
	}

	setIDRes := strings.Join(setID, ",")

	rowsFilmsGenres, err := conn.QueryContext(ctx, getGenresFilmBatchBegin+setIDRes+getGenresFilmBatchEnd)
	if err != nil {
		return []FilmSQL{}, err
	}
	defer rowsFilmsGenres.Close()

	for rowsFilmsGenres.Next() {
		var filmID int
		var genre sql.NullString

		err = rowsFilmsGenres.Scan(
			&filmID,
			&genre)
		if err != nil {
			return []FilmSQL{}, err
		}

		target[mapFilms[filmID]].Genres = append(target[mapFilms[filmID]].Genres, genre.String)
	}

	return target, nil
}

func GetShortFilmsBatch(ctx context.Context, conn *sql.Conn, query string, args ...any) ([]FilmSQL, error) {
	res := make([]FilmSQL, 0)

	//  Тут какой то жесткий баг. sql.ErrNoRows не возвращается
	rowsFilms, err := conn.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Info("NeededCondition ", err)
		return []FilmSQL{}, err
	}
	defer rowsFilms.Close()

	for rowsFilms.Next() {
		film := NewFilmSQL()

		err = rowsFilms.Scan(
			&film.ID,
			&film.Name,
			&film.OriginalName,
			&film.ProdYear,
			&film.PosterVer,
			&film.EndYear,
			&film.Rating)
		if err != nil {
			return []FilmSQL{}, err
		}

		if !film.PosterVer.Valid {
			film.PosterVer.String = innerPKG.DefFilmPosterVer
		}

		res = append(res, film)
	}

	//  Это какой то треш, запрос не отдает sql.ErrNoRows
	if len(res) == 0 {
		logrus.Info("BadCondition")
		return []FilmSQL{}, sql.ErrNoRows
	}

	return res, nil
}
