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
	countInserts := f.Config.Volume.CountReviewsLikes

	insertStatement, countAttributes := GetBatchInsertReviewsLikes(countInserts)

	values := make([]interface{}, countAttributes*countInserts)

	offset := 0
	posValue := 0

	for i := 0; i < countInserts; {
		posValue += offset
		offset = 0

		for _, value := range f.faceReviews {
			countBatchLikes := pkg.Rand(f.Config.Volume.MaxLikesOnReview)
			if (countInserts - i) < countBatchLikes {
				countBatchLikes = countInserts - i
			}

			if countBatchLikes == 0 {
				break
			}

			sequence := pkg.CryptoRandSequence(len(f.faceUsers)+1, 1)

			logrus.Info(countBatchLikes, sequence, len(sequence))

			for j := 0; j < countBatchLikes; j++ {
				values[posValue+offset] = value.ID
				offset++
				values[posValue+offset] = sequence[j]
				offset++
			}

			i += countBatchLikes
		}
	}

	target := "reviews likes"

	logrus.Info(insertStatement, "\n\n", values)

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
