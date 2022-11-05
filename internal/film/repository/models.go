package repository

import (
	"context"
	"database/sql"
	"strconv"
	"strings"

	stdErrors "github.com/pkg/errors"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	innerPKG "go-park-mail-ru/2022_2_BugOverload/internal/pkg/sqltools"
)

type FilmSQL struct {
	ID          int
	Name        string
	ProdYear    int
	Description string
	Duration    int

	ShortDescription sql.NullString
	OriginalName     sql.NullString
	Slogan           sql.NullString
	AgeLimit         sql.NullInt32
	PosterHor        sql.NullString
	PosterVer        sql.NullString

	BoxOffice      sql.NullInt32
	Budget         sql.NullInt32
	CurrencyBudget sql.NullString

	CountSeasons sql.NullInt32
	EndYear      sql.NullInt32
	Type         sql.NullString

	Rating               sql.NullFloat64
	CountScores          sql.NullInt32
	CountActors          sql.NullInt32
	CountNegativeReviews sql.NullInt32
	CountNeutralReviews  sql.NullInt32
	CountPositiveReviews sql.NullInt32

	Genres []string
}

func NewFilmSQL() FilmSQL {
	return FilmSQL{}
}

func NewFilmSQLOnFilm(film models.Film) FilmSQL {
	return FilmSQL{
		ID:          film.ID,
		Name:        film.Name,
		ProdYear:    film.ProdYear,
		Description: film.Description,
		Duration:    film.Duration,

		ShortDescription: innerPKG.NewSQLNullString(film.ShortDescription),
		OriginalName:     innerPKG.NewSQLNullString(film.OriginalName),
		Slogan:           innerPKG.NewSQLNullString(film.Slogan),
		AgeLimit:         innerPKG.NewSQLNullInt32(film.AgeLimit),
		PosterHor:        innerPKG.NewSQLNullString(film.PosterHor),
		PosterVer:        innerPKG.NewSQLNullString(film.PosterVer),

		BoxOffice:      innerPKG.NewSQLNullInt32(film.BoxOffice),
		Budget:         innerPKG.NewSQLNullInt32(film.Budget),
		CurrencyBudget: innerPKG.NewSQLNullString(film.CurrencyBudget),

		CountSeasons: innerPKG.NewSQLNullInt32(film.CountSeasons),
		EndYear:      innerPKG.NewSQLNullInt32(film.EndYear),
		Type:         innerPKG.NewSQLNullString(film.Type),

		Rating:               innerPKG.NewSQLNullFloat64(film.Rating),
		CountScores:          innerPKG.NewSQLNullInt32(film.CountScores),
		CountActors:          innerPKG.NewSQLNullInt32(film.CountActors),
		CountNegativeReviews: innerPKG.NewSQLNullInt32(film.CountNegativeReviews),
		CountNeutralReviews:  innerPKG.NewSQLNullInt32(film.CountNeutralReviews),
		CountPositiveReviews: innerPKG.NewSQLNullInt32(film.CountPositiveReviews),
	}
}

func (f *FilmSQL) Convert() models.Film {
	res := models.Film{
		ID:          f.ID,
		Name:        f.Name,
		ProdYear:    f.ProdYear,
		Description: f.Description,
		Duration:    f.Duration,

		ShortDescription: f.ShortDescription.String,
		OriginalName:     f.OriginalName.String,
		Slogan:           f.Slogan.String,
		AgeLimit:         int(f.AgeLimit.Int32),
		PosterHor:        f.PosterHor.String,
		PosterVer:        f.PosterVer.String,

		BoxOffice:      int(f.BoxOffice.Int32),
		Budget:         int(f.Budget.Int32),
		CurrencyBudget: f.CurrencyBudget.String,

		CountSeasons: int(f.CountSeasons.Int32),
		EndYear:      int(f.EndYear.Int32),
		Type:         f.Type.String,

		Rating:               float32(f.Rating.Float64),
		CountScores:          int(f.CountScores.Int32),
		CountActors:          int(f.CountActors.Int32),
		CountNegativeReviews: int(f.CountNegativeReviews.Int32),
		CountNeutralReviews:  int(f.CountNeutralReviews.Int32),
		CountPositiveReviews: int(f.CountPositiveReviews.Int32),

		Genres: f.Genres,
	}

	return res
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

func GetGenresBatch(ctx context.Context, target []FilmSQL, tx *sql.Tx) ([]FilmSQL, error) {
	setID := make([]string, len(target))

	mapFilms := make(map[int]int, len(target))

	for idx := range target {
		setID[idx] = strconv.Itoa(target[idx].ID)

		mapFilms[target[idx].ID] = idx
	}

	setIDRes := strings.Join(setID, ",")

	rowsFilmsGenres, err := tx.QueryContext(ctx, getGenresFilmBatchBegin+setIDRes+getGenresFilmBatchEnd)
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

func GetShortFilmsBatch(ctx context.Context, tx *sql.Tx, query string, values []interface{}) ([]FilmSQL, error) {
	res := make([]FilmSQL, 0)

	rowsFilms, err := tx.QueryContext(ctx, query, values...)
	if stdErrors.Is(err, sql.ErrNoRows) {
		return []FilmSQL{}, errors.ErrNotFoundInDB
	}

	if err != nil {
		return []FilmSQL{}, err
	}

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

		res = append(res, film)
	}

	return res, nil
}
