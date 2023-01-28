package fillerdb

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/pkg/errors"

	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/sqltools"
	"go-park-mail-ru/2022_2_BugOverload/pkg"
)

func (f *DBFiller) uploadReviews() (int, error) {
	countInserts := len(f.Reviews)

	countAttributes := strings.Count(insertReviews, ",") + 1

	insertStatement := sqltools.CreateFullQuery(insertReviews, countInserts, countAttributes)

	values := make([]interface{}, countAttributes*countInserts)

	pos := 0

	for _, value := range f.Reviews {
		values[pos] = value.Name
		pos++
		values[pos] = value.Type
		pos++
		values[pos] = value.Time
		pos++
		values[pos] = value.Body
		pos++
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(f.Config.Database.Timeout)*time.Second)
	defer cancelFunc()

	rows, err := sqltools.InsertBatch(ctx, f.DB.Connection, insertStatement, values)
	if err != nil {
		return 0, errors.Wrap(err, "insertCollections")
	}

	affected, err := rows.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "uploadUsers")
	}

	for i := 0; i < int(affected); i++ {
		f.Reviews[i].ID = i + 1
	}

	return countInserts, nil
}

func (f *DBFiller) linkReviewsLikes() (int, error) {
	countInserts := f.Config.Volume.CountReviewsLikes

	countAttributes := strings.Count(insertReviewsLikes, ",") + 1

	insertStatement := sqltools.CreateFullQuery(insertReviewsLikes, countInserts, countAttributes)

	values := make([]interface{}, countAttributes*countInserts)

	pos := 0
	appended := 0

	for _, value := range f.Reviews {
		count := pkg.RandMaxInt(f.Config.Volume.MaxLikesOnReview)
		if (countInserts - appended) < count {
			count = countInserts - appended
		}

		sequence := pkg.CryptoRandSequence(len(f.Users)+1, 1, count)

		for j := 0; j < count; j++ {
			values[pos] = value.ID
			pos++
			values[pos] = sequence[j]
			pos++
			values[pos] = faker.Timestamp()
			pos++
		}

		appended += count
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(f.Config.Database.Timeout)*time.Second)
	defer cancelFunc()

	_, err := sqltools.InsertBatch(ctx, f.DB.Connection, insertStatement, values)
	if err != nil {
		return 0, errors.Wrap(err, "linkReviewsLikes")
	}

	return countInserts, nil
}

func (f *DBFiller) UpdateReviews() (int, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(f.Config.Database.Timeout)*time.Second)
	defer cancelFunc()

	rows, err := f.DB.Connection.ExecContext(ctx, updateReviews)
	if err != nil {
		return 0, fmt.Errorf("UpdateReviews: [%w] when inserting row into [%s] table", err, updateReviews)
	}

	affected, err := rows.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "UpdateReviews")
	}

	return int(affected), nil
}
