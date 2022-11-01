package fillerdb

import (
	"context"
	"database/sql"
	"go-park-mail-ru/2022_2_BugOverload/pkg"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func (f *DBFiller) UploadFilms() (int, error) {
	countInserts := len(f.filmsSQL)

	insertStatement, countAttributes := GetBatchInsertFilms(countInserts)

	values := make([]interface{}, countAttributes*countInserts)

	for idx, value := range f.filmsSQL {
		posAttr := 0
		posValue := idx * countAttributes

		values[posValue+posAttr] = value.Name
		posAttr++
		values[posValue+posAttr] = value.ProdYear
		posAttr++
		values[posValue+posAttr] = value.PosterVer
		posAttr++
		values[posValue+posAttr] = value.PosterHor
		posAttr++
		values[posValue+posAttr] = value.Description
		posAttr++
		values[posValue+posAttr] = value.ShortDescription
		posAttr++
		values[posValue+posAttr] = value.OriginalName
		posAttr++
		values[posValue+posAttr] = value.Slogan
		posAttr++
		values[posValue+posAttr] = value.AgeLimit
		posAttr++
		values[posValue+posAttr] = value.BoxOffice
		posAttr++
		values[posValue+posAttr] = value.Budget
		posAttr++
		values[posValue+posAttr] = value.Duration
		posAttr++
		values[posValue+posAttr] = value.CurrencyBudget
		posAttr++
		values[posValue+posAttr] = value.Type
		posAttr++
		values[posValue+posAttr] = value.CountSeasons
		posAttr++
		values[posValue+posAttr] = value.EndYear
	}

	target := "films"

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(f.Config.Database.Timeout)*time.Second)
	defer cancelFunc()

	stmt, err := f.DB.Connection.PrepareContext(ctx, insertStatement)
	if err != nil {
		logrus.Errorf("Error [%s] when preparing SQL statement in [%s]", err, target)
		return 0, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, values...)
	if errors.Is(err, sql.ErrNoRows) {
		logrus.Infof("Info [%s] [%s]", err, target)
	}

	if err != nil {
		logrus.Errorf("Error [%s] when inserting row into [%s] table", err, target)
		return 0, err
	}
	defer rows.Close()

	counter := 0
	var insertID int64
	for rows.Next() {
		err = rows.Scan(&insertID)
		if err != nil {
			logrus.Errorf("Error [%s] when getting insertID [%s]", err, target)
			return 0, err
		}

		f.films[counter].ID = int(insertID)
		counter++
	}

	return countInserts, nil
}

func (f *DBFiller) LinkFilmsReviews() (int, error) {
	countInserts := len(f.faceReviews)

	insertStatement, countAttributes := GetBatchInsertFilmReviews(countInserts)

	values := make([]interface{}, countAttributes*countInserts)

	pos := 0
	i := 0

	sequenceReviews := pkg.CryptoRandSequence(f.faceReviews[len(f.faceReviews)-1].ID+1, f.faceReviews[0].ID)

	for _, value := range f.films {
		countBatch := pkg.Rand(f.Config.Volume.MaxReviewsOnFilm)
		if (countInserts - i) < countBatch {
			countBatch = countInserts - i
		}

		sequenceUsers := pkg.CryptoRandSequence(f.faceUsers[len(f.faceUsers)-1].ID+1, f.faceUsers[0].ID)

		for j := 0; j < countBatch; j++ {
			values[pos] = sequenceReviews[i:][j]
			pos++
			values[pos] = sequenceUsers[j]
			pos++
			values[pos] = value.ID
			pos++
		}

		i += countBatch
	}

	target := "film reviews"

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(f.Config.Database.Timeout)*time.Second)
	defer cancelFunc()

	stmt, err := f.DB.Connection.PrepareContext(ctx, insertStatement)
	if err != nil {
		logrus.Errorf("Error [%s] when preparing SQL statement in [%s]", err, target)
		return 0, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, values...)
	if err != nil {
		logrus.Errorf("Error [%s] when inserting row into [%s] table", err, target)
		return 0, err
	}
	defer rows.Close()

	return countInserts, nil
}

func (f *DBFiller) LinkFilmGenres() (int, error) {
	countInserts := 0

	for _, value := range f.films {
		countInserts += len(value.Genres)
	}

	insertStatement, countAttributes := GetBatchInsertFilmGenres(countInserts)

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

	target := "film genres"

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(f.Config.Database.Timeout)*time.Second)
	defer cancelFunc()

	stmt, err := f.DB.Connection.PrepareContext(ctx, insertStatement)
	if err != nil {
		logrus.Errorf("Error [%s] when preparing SQL statement in [%s]", err, target)
		return 0, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, values...)
	if errors.Is(err, sql.ErrNoRows) {
		logrus.Infof("Info [%s] [%s]", err, target)
	}

	if err != nil {
		logrus.Errorf("Error [%s] when inserting row into [%s] table", err, target)
		return 0, err
	}
	defer rows.Close()

	return countInserts, nil
}

func (f *DBFiller) LinkFilmCountries() (int, error) {
	countInserts := 0

	for _, value := range f.films {
		countInserts += len(value.ProdCountries)
	}

	insertStatement, countAttributes := GetBatchInsertFilmCountries(countInserts)

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

	target := "film countries"

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(f.Config.Database.Timeout)*time.Second)
	defer cancelFunc()

	stmt, err := f.DB.Connection.PrepareContext(ctx, insertStatement)
	if err != nil {
		logrus.Errorf("Error [%s] when preparing SQL statement in [%s]", err, target)
		return 0, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, values...)
	if errors.Is(err, sql.ErrNoRows) {
		logrus.Infof("Info [%s] [%s]", err, target)
	}

	if err != nil {
		logrus.Errorf("Error [%s] when inserting row into [%s] table", err, target)
		return 0, err
	}
	defer rows.Close()

	return countInserts, nil
}

func (f *DBFiller) LinkFilmCompanies() (int, error) {
	countInserts := 0

	for _, value := range f.films {
		countInserts += len(value.ProdCompanies)
	}

	insertStatement, countAttributes := GetBatchInsertFilmCompanies(countInserts)

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

	target := "film companies"

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(f.Config.Database.Timeout)*time.Second)
	defer cancelFunc()

	stmt, err := f.DB.Connection.PrepareContext(ctx, insertStatement)
	if err != nil {
		logrus.Errorf("Error [%s] when preparing SQL statement in [%s]", err, target)
		return 0, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, values...)
	if errors.Is(err, sql.ErrNoRows) {
		logrus.Infof("Info [%s] [%s]", err, target)
	}

	if err != nil {
		logrus.Errorf("Error [%s] when inserting row into [%s] table", err, target)
		return 0, err
	}
	defer rows.Close()

	return countInserts, nil
}
