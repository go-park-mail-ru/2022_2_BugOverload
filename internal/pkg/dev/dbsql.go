package dev

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const countAttributesFilms = 6

func (f *DBFiller) UploadFilms() (int, error) {
	var values []interface{}

	countInserts := len(f.films)

	placeholders := CreatePlaceholders(countAttributesFilms, countInserts)

	for _, value := range f.films {
		values = append(values,
			value.Name,
			value.ProdYear,
			value.PosterVer,
			value.PosterHor,
			value.Description,
			value.ShortDescription)
	}

	queryBegin := "INSERT INTO films(name, prod_year, poster_ver, poster_hor, description, short_description) VALUES "

	queryEnd := "RETURNING film_id;"

	insertStatement := fmt.Sprintf("%s %s \n %s", queryBegin, placeholders, queryEnd)

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
		logrus.Infof("Info [%s] films", err)
	}

	if err != nil {
		logrus.Errorf("Error [%s] when inserting row into films table", err)
		return 0, err
	}
	defer rows.Close()

	counter := 0

	for rows.Next() {
		var insertID int64

		err = rows.Scan(&insertID)
		if err != nil {
			logrus.Errorf("Error [%s] when getting insertID films", err)
			return 0, err
		}

		f.films[counter].ID = int(insertID)
		counter++
	}

	return countInserts, nil
}
