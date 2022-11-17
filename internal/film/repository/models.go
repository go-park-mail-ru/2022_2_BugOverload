package repository

import (
	"context"
	"database/sql"
	"strconv"
	"strings"
	"time"

	stdErrors "github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	innerPKG "go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
)

type AuthorSQL struct {
	ID           int
	Nickname     string
	CountReviews sql.NullInt32
	Avatar       sql.NullString
}

type ReviewSQL struct {
	Name       string
	Type       string
	Body       string
	CountLikes sql.NullInt32
	CreateTime time.Time
	Author     AuthorSQL
}

func (r *ReviewSQL) Convert() models.Review {
	return models.Review{
		Name:       r.Name,
		Type:       r.Type,
		Body:       r.Body,
		CreateTime: r.CreateTime.Format(innerPKG.DateFormat + " " + innerPKG.TimeFormat),
		CountLikes: int(r.CountLikes.Int32),
		Author: models.User{
			ID:           r.Author.ID,
			Nickname:     r.Author.Nickname,
			Avatar:       r.Author.Avatar.String,
			CountReviews: int(r.Author.CountReviews.Int32),
		},
	}
}

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
	ProdDate    string
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
	CountRatings         sql.NullInt32
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
	endYear := ""

	if f.EndYear.Valid {
		endYear = f.EndYear.Time.Format(innerPKG.OnlyDate)
	}

	res := models.Film{
		ID:              f.ID,
		Name:            f.Name,
		ProdDate:        f.ProdDate,
		Description:     f.Description,
		DurationMinutes: f.Duration,

		ShortDescription: f.ShortDescription.String,
		OriginalName:     f.OriginalName.String,
		Slogan:           f.Slogan.String,
		AgeLimit:         f.AgeLimit.String,
		PosterHor:        f.PosterHor.String,
		PosterVer:        f.PosterVer.String,

		BoxOfficeDollars: int(f.BoxOffice.Int32),
		Budget:           int(f.Budget.Int32),
		CurrencyBudget:   f.CurrencyBudget.String,

		CountSeasons: int(f.CountSeasons.Int32),
		EndYear:      endYear,
		Type:         f.Type.String,

		Rating:               float32(f.Rating.Float64),
		CountRatings:         int(f.CountRatings.Int32),
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
		return []FilmSQL{}, stdErrors.WithMessagef(errors.ErrWorkDatabase,
			"Err: params input: query - [%s]. Special Error [%s]",
			getGenresFilmBatchBegin+setIDRes+getGenresFilmBatchEnd, err)
	}
	defer rowsFilmsGenres.Close()

	for rowsFilmsGenres.Next() {
		var filmID int
		var genre sql.NullString

		err = rowsFilmsGenres.Scan(&filmID, &genre)
		if err != nil {
			return []FilmSQL{}, stdErrors.WithMessagef(errors.ErrWorkDatabase,
				"Err Scan: params input: query - [%s]. Special Error [%s]",
				getGenresFilmBatchBegin+setIDRes+getGenresFilmBatchEnd, err)
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
			&film.ProdDate,
			&film.PosterVer,
			&film.Type,
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

	for idx := range res {
		if res[idx].Type.String == innerPKG.DefTypeSerial {
			rowSerial := conn.QueryRowContext(ctx, getShortSerialByID, res[idx].ID)
			if rowSerial.Err() != nil {
				return []FilmSQL{}, rowSerial.Err()
			}

			err = rowSerial.Scan(&res[idx].EndYear)
			if err != nil {
				return []FilmSQL{}, err
			}
		}
	}

	return res, nil
}
