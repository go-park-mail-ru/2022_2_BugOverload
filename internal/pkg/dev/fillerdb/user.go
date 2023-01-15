package fillerdb

import (
	"context"
	"fmt"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/security"
	"strings"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/pkg/errors"

	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/sqltools"
	"go-park-mail-ru/2022_2_BugOverload/pkg"
)

func (f *DBFiller) uploadUsers() (int, error) {
	countInserts := len(f.Users)

	countAttributes := strings.Count(insertUsers, ",") + 1

	insertStatement := sqltools.CreateFullQuery(insertUsers, countInserts, countAttributes)

	values := make([]interface{}, countAttributes*countInserts)

	pos := 0

	for _, value := range f.Users {
		values[pos] = value.Nickname
		pos++

		values[pos] = value.Email
		pos++

		hash, err := security.HashPassword(value.Password)
		if err != nil {
			return 0, errors.Wrap(err, "uploadUsers")
		}

		values[pos] = []byte(hash)
		pos++
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(f.Config.Database.Timeout)*time.Second)
	defer cancelFunc()

	rows, err := sqltools.InsertBatch(ctx, f.DB.Connection, insertStatement, values)
	if err != nil {
		return 0, errors.Wrap(err, "uploadUsers")
	}

	affected, err := rows.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "uploadUsers")
	}

	for i := 0; i < int(affected); i++ {
		f.Users[i].ID = i + 1
	}

	return countInserts, nil
}

func (f *DBFiller) linkProfileViews() (int, error) {
	countInserts := f.Config.Volume.CountViews

	countAttributes := strings.Count(insertProfileViews, ",") + 1

	insertStatement := sqltools.CreateFullQuery(insertProfileViews, countInserts, countAttributes)

	values := make([]interface{}, countAttributes*countInserts)

	pos := 0
	appended := 0

	for _, value := range f.Users {
		count := pkg.RandMaxInt(f.Config.Volume.MaxViewOnFilm)
		if (countInserts - appended) < count {
			count = countInserts - appended
		}

		if count == 0 {
			continue
		}

		faker.Word()

		sequence := pkg.CryptoRandSequence(count+1, 1)

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
		return 0, errors.Wrap(err, "linkProfileViews")
	}

	return countInserts, nil
}

const offset = 3

func (f *DBFiller) linkProfileRatings() (int, error) {
	countInserts := f.Config.Volume.CountRatings

	countAttributes := strings.Count(insertProfileRatings, ",") + 1

	insertStatement := sqltools.CreateFullQuery(insertProfileRatings, countInserts, countAttributes)

	values := make([]interface{}, countAttributes*countInserts)

	pos := 0
	appended := 0

	for _, value := range f.Users {
		count := pkg.RandMaxInt(f.Config.Volume.MaxCountRatingsOnFilm)
		if (countInserts - appended) < count {
			count = countInserts - appended
		}

		if count == 0 {
			continue
		}

		sequence := pkg.CryptoRandSequence(count+1, 1)

		for j := 0; j < count; j++ {
			filmID := sequence[j]

			score := pkg.RandMaxInt(f.Config.Volume.MaxRatings-offset) + 1 + offset

			values[pos] = value.ID
			pos++
			values[pos] = filmID
			pos++
			values[pos] = score
			pos++
			values[pos] = faker.Timestamp()
			pos++

			// For update denormal fields
			f.films[filmID-1].Rating += float64(score)

			f.films[filmID-1].CountScores++
		}

		appended += count
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(f.Config.Database.Timeout)*time.Second)
	defer cancelFunc()

	_, err := sqltools.InsertBatch(ctx, f.DB.Connection, insertStatement, values)
	if err != nil {
		return 0, errors.Wrap(err, "linkProfileRatings")
	}

	return countInserts, nil
}

func (f *DBFiller) UpdateProfiles() (int, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(f.Config.Database.Timeout)*time.Second)
	defer cancelFunc()

	rows, err := f.DB.Connection.ExecContext(ctx, updateProfiles)
	if err != nil {
		return 0, fmt.Errorf("UpdateProfiles: [%w] when inserting row into [%s] table", err, updateProfiles)
	}

	affected, err := rows.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "UpdateProfiles")
	}

	return int(affected), nil
}
