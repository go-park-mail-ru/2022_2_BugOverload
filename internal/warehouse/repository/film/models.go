package film

import (
	"context"
	"database/sql"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	"strconv"
	"strings"
	"time"

	stdErrors "github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
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
		CreateTime: r.CreateTime.Format(constparams.DateFormat + " " + constparams.TimeFormat),
		CountLikes: int(r.CountLikes.Int32),
		Author: models.User{
			ID:           r.Author.ID,
			Nickname:     r.Author.Nickname,
			Avatar:       r.Author.Avatar.String,
			CountReviews: int(r.Author.CountReviews.Int32),
		},
	}
}

type ModelActorSQL struct {
	ID        int
	Name      string
	Avatar    sql.NullString
	Character string
}

type ModelPersonSQL struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ModelSQL struct {
	ID              int
	Name            string
	ProdDate        string
	Description     string
	DurationMinutes int

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

	Actors    []ModelActorSQL
	Artists   []ModelPersonSQL
	Directors []ModelPersonSQL
	Writers   []ModelPersonSQL
	Producers []ModelPersonSQL
	Operators []ModelPersonSQL
	Montage   []ModelPersonSQL
	Composers []ModelPersonSQL
}

func NewFilmSQL() ModelSQL {
	return ModelSQL{}
}

func (f *ModelSQL) Convert() models.Film {
	endYear := ""

	if f.EndYear.Valid {
		endYear = f.EndYear.Time.Format(constparams.OnlyDate)
	}

	res := models.Film{
		ID:              f.ID,
		Name:            f.Name,
		ProdDate:        f.ProdDate,
		Description:     f.Description,
		DurationMinutes: f.DurationMinutes,

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
	GetGenresFilmBatchBegin = `
SELECT f.film_id,
       g.name
FROM genres g
         JOIN film_genres fg ON g.genre_id = fg.fk_genre_id
         JOIN films f ON fg.fk_film_id = f.film_id
WHERE f.film_id IN (`

	GetGenresFilmBatchEnd = `) ORDER BY f.film_id, fg.weight DESC`
)

func GetGenresBatch(ctx context.Context, target []ModelSQL, conn *sql.Conn) ([]ModelSQL, error) {
	setID := make([]string, len(target))

	mapFilms := make(map[int]int, len(target))

	for idx := range target {
		setID[idx] = strconv.Itoa(target[idx].ID)

		mapFilms[target[idx].ID] = idx
	}

	setIDRes := strings.Join(setID, ",")

	query := GetGenresFilmBatchBegin + setIDRes + GetGenresFilmBatchEnd

	rowsFilmsGenres, err := conn.QueryContext(ctx, query)
	if err != nil {
		return []ModelSQL{}, stdErrors.WithMessagef(errors.ErrWorkDatabase,
			"GetGenresBatch: Err: params input: query - [%s]. Special Error [%s]",
			query, err)
	}
	defer rowsFilmsGenres.Close()

	for rowsFilmsGenres.Next() {
		var filmID int
		var genre sql.NullString

		err = rowsFilmsGenres.Scan(&filmID, &genre)
		if err != nil {
			return []ModelSQL{}, stdErrors.WithMessagef(errors.ErrWorkDatabase,
				"GetGenresBatch: Err: Scan: params input: query - [%s]. Special Error [%s]",
				query, err)
		}

		if len(target[mapFilms[filmID]].Genres) < constparams.MaxCountAttrInCollection {
			target[mapFilms[filmID]].Genres = append(target[mapFilms[filmID]].Genres, genre.String)
		}
	}

	return target, nil
}

const (
	GetProdCountriesFilmBatchBegin = `
SELECT f.film_id,
       countries.name
FROM countries
         JOIN film_countries fc on countries.country_id = fc.fk_country_id
         JOIN films f ON fc.fk_film_id = f.film_id
WHERE f.film_id IN (`

	GetProdCountriesBatchEnd = `) ORDER BY f.film_id, fc.weight DESC`
)

func GetProdCountriesBatch(ctx context.Context, target []ModelSQL, conn *sql.Conn) ([]ModelSQL, error) {
	setID := make([]string, len(target))

	mapFilms := make(map[int]int, len(target))

	for idx := range target {
		setID[idx] = strconv.Itoa(target[idx].ID)

		mapFilms[target[idx].ID] = idx
	}

	setIDRes := strings.Join(setID, ",")

	query := GetProdCountriesFilmBatchBegin + setIDRes + GetProdCountriesBatchEnd

	rowsFilmsProdCountries, err := conn.QueryContext(ctx, query)
	if err != nil {
		return []ModelSQL{}, stdErrors.WithMessagef(errors.ErrWorkDatabase,
			"Err: params input: query - [%s]. Special Error [%s]",
			query, err)
	}
	defer rowsFilmsProdCountries.Close()

	for rowsFilmsProdCountries.Next() {
		var filmID int
		var country sql.NullString

		err = rowsFilmsProdCountries.Scan(&filmID, &country)
		if err != nil {
			return []ModelSQL{}, stdErrors.WithMessagef(errors.ErrWorkDatabase,
				"Err Scan: params input: query - [%s]. Special Error [%s]",
				query, err)
		}

		if len(target[mapFilms[filmID]].ProdCountries) < constparams.MaxCountAttrInCollection {
			target[mapFilms[filmID]].ProdCountries = append(target[mapFilms[filmID]].ProdCountries, country.String)
		}
	}

	return target, nil
}

const (
	GetDirectorsFilmBatchBegin = `
SELECT DISTINCT
        f.film_id,
        persons.person_id,
        persons.name
FROM films f
    JOIN film_persons fp on fp.fk_film_id = f.film_id
    JOIN persons ON persons.person_id = fp.fk_person_id
WHERE persons.person_id IN (
    SELECT
        p.person_id
    FROM persons p
        JOIN person_professions pp on p.person_id = pp.fk_person_id
        JOIN professions on professions.profession_id = pp.fk_profession_id
    WHERE professions.name = 'режиссер'
    ORDER BY pp.weight DESC
) AND f.film_id IN (`

	GetDirectorsFilmBatchEnd = `) ORDER BY f.film_id`
)

func GetDirectorsBatch(ctx context.Context, target []ModelSQL, conn *sql.Conn) ([]ModelSQL, error) {
	setID := make([]string, len(target))

	mapFilms := make(map[int]int, len(target))

	for idx := range target {
		setID[idx] = strconv.Itoa(target[idx].ID)

		mapFilms[target[idx].ID] = idx
	}

	setIDRes := strings.Join(setID, ",")

	query := GetDirectorsFilmBatchBegin + setIDRes + GetDirectorsFilmBatchEnd

	rowsFilmsDirectors, err := conn.QueryContext(ctx, query)
	if err != nil {
		return []ModelSQL{}, stdErrors.WithMessagef(errors.ErrWorkDatabase,
			"Err: params input: query - [%s]. Special Error [%s]",
			query, err)
	}
	defer rowsFilmsDirectors.Close()

	for rowsFilmsDirectors.Next() {
		var filmID int
		var directorID int
		var directorName sql.NullString

		err = rowsFilmsDirectors.Scan(&filmID, &directorID, &directorName)
		if err != nil {
			return []ModelSQL{}, stdErrors.WithMessagef(errors.ErrWorkDatabase,
				"Err Scan: params input: query - [%s]. Special Error [%s]",
				query, err)
		}

		if len(target[mapFilms[filmID]].Directors) < constparams.MaxCountAttrInCollection {
			target[mapFilms[filmID]].Directors = append(target[mapFilms[filmID]].Directors,
				ModelPersonSQL{ID: directorID, Name: directorName.String})
		}
	}

	return target, nil
}

func GetShortFilmsBatch(ctx context.Context, conn *sql.Conn, query string, args ...any) ([]ModelSQL, error) {
	res := make([]ModelSQL, 0)

	//  Тут какой то жесткий баг. sql.ErrNoRows не возвращается
	rowsFilms, err := conn.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Info("NeededCondition ", err)
		return []ModelSQL{}, err
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
			return []ModelSQL{}, err
		}

		if !film.PosterVer.Valid {
			film.PosterVer.String = constparams.DefFilmPosterVer
		}

		res = append(res, film)
	}

	//  Это какой то треш, запрос не отдает sql.ErrNoRows
	if len(res) == 0 {
		logrus.Info("BadCondition")
		return []ModelSQL{}, sql.ErrNoRows
	}

	for idx := range res {
		if res[idx].Type.String == constparams.DefTypeSerial {
			rowSerial := conn.QueryRowContext(ctx, GetShortSerialByID, res[idx].ID)
			if rowSerial.Err() != nil {
				return []ModelSQL{}, rowSerial.Err()
			}

			err = rowSerial.Scan(&res[idx].EndYear)
			if err != nil {
				return []ModelSQL{}, err
			}
		}
	}

	return res, nil
}

func GetFilmsPremieresBatch(ctx context.Context, conn *sql.Conn, args ...any) ([]ModelSQL, error) {
	res := make([]ModelSQL, 0)

	rowsFilms, err := conn.QueryContext(ctx, GetNewFilms, args...)
	if err != nil {
		logrus.Info("NeededCondition ", err)
		return []ModelSQL{}, err
	}
	defer rowsFilms.Close()

	for rowsFilms.Next() {
		film := NewFilmSQL()

		err = rowsFilms.Scan(
			&film.ID,
			&film.Name,
			&film.ProdDate,
			&film.PosterVer,
			&film.Rating,
			&film.DurationMinutes,
			&film.Description)
		if err != nil {
			return []ModelSQL{}, err
		}

		if !film.PosterVer.Valid {
			film.PosterVer.String = constparams.DefFilmPosterVer
		}

		res = append(res, film)
	}

	if len(res) == 0 {
		logrus.Info("BadCondition")
		return []ModelSQL{}, sql.ErrNoRows
	}

	return res, nil
}

func GetFilmsRealisesBatch(ctx context.Context, conn *sql.Conn, args ...any) ([]ModelSQL, error) {
	res := make([]ModelSQL, 0)

	rowsFilms, err := conn.QueryContext(ctx, getRelease, args...)
	if err != nil {
		logrus.Info("NeededCondition ", err)
		return []ModelSQL{}, err
	}
	defer rowsFilms.Close()

	for rowsFilms.Next() {
		film := NewFilmSQL()

		err = rowsFilms.Scan(
			&film.ID,
			&film.Name,
			&film.ProdDate,
			&film.PosterVer,
			&film.Rating)
		if err != nil {
			return []ModelSQL{}, err
		}

		if !film.PosterVer.Valid {
			film.PosterVer.String = constparams.DefFilmPosterVer
		}

		res = append(res, film)
	}

	if len(res) == 0 {
		logrus.Info("BadCondition")
		return []ModelSQL{}, sql.ErrNoRows
	}

	return res, nil
}
