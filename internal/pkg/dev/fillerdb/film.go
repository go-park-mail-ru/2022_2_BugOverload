package fillerdb

import (
	"context"
	"database/sql"
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

		f.films[counter].ID = int(insertID)
		counter++
	}

	return countInserts, nil
}
