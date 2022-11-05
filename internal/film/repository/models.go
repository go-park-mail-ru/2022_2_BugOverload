package repository

import (
	"context"
	"database/sql"
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	innerPKG "go-park-mail-ru/2022_2_BugOverload/internal/pkg/sqltools"
	"strconv"
	"strings"
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
	IDSet := make([]string, len(target))

	for idx := range target {
		IDSet[idx] = strconv.Itoa(target[idx].ID)
	}

	IDSetResult := strings.Join(IDSet, ",")

	rowsFilmsGenres, err := tx.QueryContext(ctx, getGenresFilmBatchBegin+IDSetResult+getGenresFilmBatchEnd)
	if err != nil {
		return []FilmSQL{}, err
	}

	defer rowsFilmsGenres.Close()

	counter := 0

	for rowsFilmsGenres.Next() {
		var filmID int
		var genre sql.NullString

		err = rowsFilmsGenres.Scan(
			&filmID,
			&genre)
		if err != nil {
			return []FilmSQL{}, err
		}

		if filmID != target[counter].ID {
			counter++
		}

		target[counter].Genres = append(target[counter].Genres, genre.String)
	}

	return target, nil
}
