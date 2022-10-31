package fillerdb

import (
	"context"
	"database/sql"
	"go-park-mail-ru/2022_2_BugOverload/pkg"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

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

		f.faceReviews[counter].ID = int(insertID)
		counter++
	}

	return countInserts, nil
}

func (f *DBFiller) LinkReviewsLikes() (int, error) {
	countInserts := len(f.faceReviews)

	insertStatement, countAttributes := GetBatchInsertReviewsLikes(countInserts)

	values := make([]interface{}, countAttributes*countInserts)

	for idx, value := range f.faceReviews {
		posAttr := 0
		posValue := idx * countAttributes

		values[posValue+posAttr] = value.ID
		posAttr++
		values[posValue+posAttr] = f.faceUsers[pkg.Rand(len(f.faceUsers))].ID
	}

	target := "reviews likes"

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
