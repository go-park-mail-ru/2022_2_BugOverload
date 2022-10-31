package fillerdb

import (
	"context"
	"database/sql"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

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

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(f.Config.Database.Timeout)*time.Second)
	defer cancelFunc()

	stmt, err := f.DB.Connection.PrepareContext(ctx, insertStatement)
	if err != nil {
		logrus.Errorf("Error [%s] when preparing SQL statement", err)
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

		f.persons[counter].ID = int(insertID)
		counter++
	}

	return countInserts, nil
}

func (f *DBFiller) LinkPersonProfession() (int, error) {
	countInserts := 0

	for _, value := range f.persons {
		countInserts += len(value.Professions)
	}

	insertStatement, countAttributes := GetBatchInsertPersonProfessions(countInserts)

	values := make([]interface{}, countAttributes*countInserts)

	offset := 0
	posValue := 0

	for _, value := range f.persons {
		posValue += offset
		offset = 0
		weight := len(value.Professions)

		for _, profession := range value.Professions {
			values[posValue+offset] = value.ID
			offset++
			values[posValue+offset] = f.professions[profession]
			weight--
			offset++
			values[posValue+offset] = weight
			offset++
		}
	}

	target := "persons professions"

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(f.Config.Database.Timeout)*time.Second)
	defer cancelFunc()

	stmt, err := f.DB.Connection.PrepareContext(ctx, insertStatement)
	if err != nil {
		logrus.Errorf("Error [%s] when preparing SQL statement", err)
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
