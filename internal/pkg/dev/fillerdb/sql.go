package fillerdb

import (
	"context"
	"database/sql"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func (f *DBFiller) GetQueryRes(insertStatement string, values []interface{}, target string) (*sql.Rows, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(f.Config.Database.Timeout)*time.Second)
	defer cancelFunc()

	stmt, err := f.DB.Connection.PrepareContext(ctx, insertStatement)
	if err != nil {
		logrus.Errorf("Error [%s] when preparing SQL statement", err)
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, values...)
	if errors.Is(err, sql.ErrNoRows) {
		logrus.Infof("Info [%s] [%s]", err, target)
	}

	if err != nil {
		logrus.Errorf("Error [%s] when inserting row into [%s] table", err, target)
		return nil, err
	}

	return rows, nil
}

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

	rows, err := f.GetQueryRes(insertStatement, values, target)
	if err != nil {
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

func (f *DBFiller) UploadPersons() (int, error) {
	countInserts := len(f.personsSQL)

	insertStatement, countAttributes := GetBatchInsertPersons(countInserts)

	values := make([]interface{}, countAttributes*countInserts)

	for idx, value := range f.personsSQL {
		posAttr := 0
		posValue := idx * countAttributes

		values[posValue+posAttr] = value.Name
		posAttr++
		values[posValue+posAttr] = value.OriginalName
		posAttr++
		values[posValue+posAttr] = value.Birthday
		posAttr++
		values[posValue+posAttr] = value.Growth
		posAttr++
		values[posValue+posAttr] = value.Avatar
		posAttr++
		values[posValue+posAttr] = value.Gender
		posAttr++
		values[posValue+posAttr] = value.Death
	}

	target := "persons"

	rows, err := f.GetQueryRes(insertStatement, values, target)
	if err != nil {
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

		f.persons[counter].ID = int(insertID)
		counter++
	}

	return countInserts, nil
}

func (f *DBFiller) UploadUsers() (int, error) {
	countInserts := len(f.faceUsers)

	insertStatement, countAttributes := GetBatchInsertUsers(countInserts)

	values := make([]interface{}, countAttributes*countInserts)

	for idx, value := range f.faceUsers {
		posAttr := 0
		posValue := idx * countAttributes

		values[posValue+posAttr] = value.Nickname
		posAttr++
		values[posValue+posAttr] = value.Email
		posAttr++
		values[posValue+posAttr] = value.Password
	}

	target := "persons"

	rows, err := f.GetQueryRes(insertStatement, values, target)
	if err != nil {
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

		f.faceUsers[counter].ID = int(insertID)
		counter++
	}

	return countInserts, nil
}

func (f *DBFiller) LinkUsersProfiles() (int, error) {
	countInserts := len(f.faceUsers)

	insertStatement, countAttributes := GetBatchInsertProfiles(countInserts)

	values := make([]interface{}, countAttributes*countInserts)

	for idx, value := range f.faceUsers {
		values[idx] = value.ID
	}

	target := "profiles"

	rows, err := f.GetQueryRes(insertStatement, values, target)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	return countInserts, nil
}

func (f *DBFiller) UploadReviews() (int, error) {
	countInserts := len(f.faceReviews)

	insertStatement, countAttributes := GetBatchInsertReviews(countInserts)

	values := make([]interface{}, countAttributes*countInserts)

	for idx, value := range f.faceReviews {
		posAttr := 0
		posValue := idx * countAttributes

		values[posValue+posAttr] = value.Name
		posAttr++
		values[posValue+posAttr] = value.Type
		posAttr++
		values[posValue+posAttr] = value.Time
		posAttr++
		values[posValue+posAttr] = value.Body
	}

	target := "reviews"

	rows, err := f.GetQueryRes(insertStatement, values, target)
	if err != nil {
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

		f.faceReviews[counter].ID = int(insertID)
		counter++
	}

	return countInserts, nil
}
