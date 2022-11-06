package fillerdb

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/sqltools"
	"go-park-mail-ru/2022_2_BugOverload/pkg"
)

func (f *DBFiller) uploadFilms() (int, error) {
	countInserts := len(f.filmsSQL)

	insertStatement, countAttributes := sqltools.CreateFullQuery(insertFilms, countInserts)

	values := make([]interface{}, countAttributes*countInserts)

	pos := 0

	for _, value := range f.filmsSQL {
		values[pos] = value.Name
		pos++
		values[pos] = value.ProdYear
		pos++
		values[pos] = value.PosterVer
		pos++
		values[pos] = value.PosterHor
		pos++
		values[pos] = value.Description
		pos++
		values[pos] = value.ShortDescription
		pos++
		values[pos] = value.OriginalName
		pos++
		values[pos] = value.Slogan
		pos++
		values[pos] = value.AgeLimit
		pos++
		values[pos] = value.BoxOffice
		pos++
		values[pos] = value.Budget
		pos++
		values[pos] = value.Duration
		pos++
		values[pos] = value.CurrencyBudget
		pos++
		values[pos] = value.Type
		pos++
		values[pos] = value.CountSeasons
		pos++
		values[pos] = value.EndYear
		pos++
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(f.Config.Database.Timeout)*time.Second)
	defer cancelFunc()

	rows, err := sqltools.InsertBatch(ctx, f.DB.Connection, insertStatement, values)
	if err != nil {
		return 0, errors.Wrap(err, "uploadFilms")
	}

	affected, err := rows.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "uploadFilms")
	}

	for i := 0; i < int(affected); i++ {
		f.films[i].ID = i + 1
	}

	return countInserts, nil
}

func (f *DBFiller) linkFilmsReviews() (int, error) {
	countInserts := len(f.faceReviews)

	insertStatement, countAttributes := sqltools.CreateFullQuery(insertFilmsReviews, countInserts)

	values := make([]interface{}, countAttributes*countInserts)

	pos := 0
	appended := 0

	sequenceReviews := pkg.CryptoRandSequence(len(f.faceReviews)+1, 1)

	for _, value := range f.films {
		countPartBatch := pkg.RandMaxInt(f.Config.Volume.MaxReviewsOnFilm)
		if (countInserts - appended) < countPartBatch {
			countPartBatch = countInserts - appended
		}

		sequenceUsers := pkg.CryptoRandSequence(len(f.faceUsers)+1, 1)

		for j := 0; j < countPartBatch; j++ {
			values[pos] = sequenceReviews[appended:][j]
			pos++
			values[pos] = sequenceUsers[j]
			pos++
			values[pos] = value.ID
			pos++
		}

		appended += countPartBatch
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(f.Config.Database.Timeout)*time.Second)
	defer cancelFunc()

	_, err := sqltools.InsertBatch(ctx, f.DB.Connection, insertStatement, values)
	if err != nil {
		return 0, errors.Wrap(err, "linkFilmsReviews")
	}

	return countInserts, nil
}

func (f *DBFiller) linkFilmGenres() (int, error) {
	countInserts := 0

	for _, value := range f.films {
		countInserts += len(value.Genres)
	}

	insertStatement, countAttributes := sqltools.CreateFullQuery(insertFilmsGenres, countInserts)

	values := make([]interface{}, countAttributes*countInserts)

	offset := 0
	posValue := 0

	for _, value := range f.films {
		posValue += offset
		offset = 0
		weight := len(value.Genres)

		for _, genre := range value.Genres {
			values[posValue+offset] = value.ID
			offset++
			values[posValue+offset] = f.genres[genre]
			weight--
			offset++
			values[posValue+offset] = weight
			offset++
		}
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(f.Config.Database.Timeout)*time.Second)
	defer cancelFunc()

	_, err := sqltools.InsertBatch(ctx, f.DB.Connection, insertStatement, values)
	if err != nil {
		return 0, errors.Wrap(err, "linkFilmGenres")
	}

	return countInserts, nil
}

func (f *DBFiller) linkFilmCountries() (int, error) {
	countInserts := 0

	for _, value := range f.films {
		countInserts += len(value.ProdCountries)
	}

	insertStatement, countAttributes := sqltools.CreateFullQuery(insertFilmsCountries, countInserts)

	values := make([]interface{}, countAttributes*countInserts)

	offset := 0
	posValue := 0

	for _, value := range f.films {
		posValue += offset
		offset = 0
		weight := len(value.ProdCountries)

		for _, country := range value.ProdCountries {
			values[posValue+offset] = value.ID
			offset++
			_, ok := f.countries[country]
			if !ok {
				logrus.Error(country)
			}
			values[posValue+offset] = f.countries[country]
			weight--
			offset++
			values[posValue+offset] = weight
			offset++
		}
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(f.Config.Database.Timeout)*time.Second)
	defer cancelFunc()

	_, err := sqltools.InsertBatch(ctx, f.DB.Connection, insertStatement, values)
	if err != nil {
		return 0, errors.Wrap(err, "linkFilmCountries")
	}

	return countInserts, nil
}

func (f *DBFiller) linkFilmCompanies() (int, error) {
	countInserts := 0

	for _, value := range f.films {
		countInserts += len(value.ProdCompanies)
	}

	insertStatement, countAttributes := sqltools.CreateFullQuery(insertFilmsCompanies, countInserts)

	values := make([]interface{}, countAttributes*countInserts)

	offset := 0
	posValue := 0

	for _, value := range f.films {
		posValue += offset
		offset = 0
		weight := len(value.ProdCompanies)

		for _, company := range value.ProdCompanies {
			values[posValue+offset] = value.ID
			offset++
			values[posValue+offset] = f.companies[company]
			weight--
			offset++
			values[posValue+offset] = weight
			offset++
		}
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(f.Config.Database.Timeout)*time.Second)
	defer cancelFunc()

	_, err := sqltools.InsertBatch(ctx, f.DB.Connection, insertStatement, values)
	if err != nil {
		return 0, errors.Wrap(err, "linkFilmCompanies")
	}

	return countInserts, nil
}

func (f *DBFiller) linkFilmPersons() (int, error) {
	values := make([]interface{}, 0)

	countInserts := 0

	for _, value := range f.films {
		countActors := pkg.RandMaxInt(f.Config.Volume.MaxFilmActors) + 1
		weightActors := countActors - 1

		sequenceActors := pkg.CryptoRandSequence(len(f.persons)+1, 1)

		for i := 0; i < countActors; i++ {
			values = append(values, sequenceActors[i], value.ID, f.professions["актер"], sqltools.NewSQLNullString(faker.Word()), weightActors)
			weightActors--
		}

		for profession := 2; profession < len(f.professions); profession++ {
			countPersons := pkg.RandMaxInt(f.Config.Volume.MaxFilmPersons) + 1
			weightPersons := countPersons - 1

			sequencePersons := pkg.CryptoRandSequence(len(f.persons)+1, 1)

			for i := 0; i < countPersons; i++ {
				values = append(values, sequencePersons[i], value.ID, profession, sqltools.NewSQLNullString(""), weightPersons)
				weightPersons--
			}

			countInserts += countPersons
		}

		countInserts += countActors
	}

	insertStatement, _ := sqltools.CreateFullQuery(insertFilmsPersons, countInserts)

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(f.Config.Database.Timeout)*time.Second)
	defer cancelFunc()

	_, err := sqltools.InsertBatch(ctx, f.DB.Connection, insertStatement, values)
	if err != nil {
		return 0, errors.Wrap(err, "linkFilmPersons")
	}

	return countInserts, nil
}

func (f *DBFiller) linkFilmTags() (int, error) {
	countInserts := 0

	values := make([]interface{}, 0)

	for _, value := range f.tags {
		count := pkg.RandMaxInt(f.Config.Volume.MaxFilmsInTag) + 1

		sequence := pkg.CryptoRandSequence(len(f.films)+1, 1)

		for i := 0; i < count; i++ {
			values = append(values, sequence[i], value)
		}

		countInserts += count
	}

	insertStatement, _ := sqltools.CreateFullQuery(insertFilmsTags, countInserts)

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(f.Config.Database.Timeout)*time.Second)
	defer cancelFunc()

	_, err := sqltools.InsertBatch(ctx, f.DB.Connection, insertStatement, values)
	if err != nil {
		return 0, errors.Wrap(err, "linkFilmTags")
	}

	return countInserts, nil
}

func (f *DBFiller) linkFilmImages() (int, error) {
	countInserts := len(f.films)

	insertStatement, countAttributes := sqltools.CreateFullQuery(insertFilmsImages, countInserts)

	values := make([]interface{}, countAttributes*countInserts)

	pos := 0
	for _, value := range f.films {
		values[pos] = value.ID
		pos++

		imagesList := strings.Join(value.Images, "_")
		values[pos] = imagesList
		pos++
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(f.Config.Database.Timeout)*time.Second)
	defer cancelFunc()

	_, err := sqltools.InsertBatch(ctx, f.DB.Connection, insertStatement, values)
	if err != nil {
		return 0, errors.Wrap(err, "linkPersonImages")
	}

	return countInserts, nil
}

func (f *DBFiller) UpdateFilms() (int, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(f.Config.Database.Timeout)*time.Second)
	defer cancelFunc()

	rows, err := f.DB.Connection.ExecContext(ctx, updateFilms)
	if err != nil {
		return 0, fmt.Errorf("UpdateFilms: [%w] when inserting row into [%s] table", err, updateFilms)
	}

	affected, err := rows.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "UpdateFilms")
	}

	return int(affected), nil
}
