package fillerdb

import (
	"context"
	"fmt"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/pkg/errors"

	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/sqltools"
	"go-park-mail-ru/2022_2_BugOverload/pkg"
)

func (f *DBFiller) uploadUsers() (int, error) {
	countInserts := len(f.faceUsers)

	insertStatement, countAttributes := sqltools.CreateFullQuery(insertUsers, countInserts)

	values := make([]interface{}, countAttributes*countInserts)

	pos := 0

	for _, value := range f.faceUsers {
		values[pos] = value.Nickname
		pos++
		values[pos] = value.Email
		pos++
		values[pos] = value.Password
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
		f.faceUsers[i].ID = i + 1
	}

	return countInserts, nil
}

func (f *DBFiller) linkUsersProfiles() (int, error) {
	countInserts := len(f.faceUsers)

	insertStatement, countAttributes := sqltools.CreateFullQuery(insertUsersProfiles, countInserts)

	values := make([]interface{}, countAttributes*countInserts)

	pos := 0
	for _, value := range f.faceUsers {
		values[pos] = value.ID
		pos++
		values[pos] = faker.Timestamp()
		pos++
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(f.Config.Database.Timeout)*time.Second)
	defer cancelFunc()

	_, err := sqltools.InsertBatch(ctx, f.DB.Connection, insertStatement, values)
	if err != nil {
		return 0, errors.Wrap(err, "linkUsersProfiles")
	}

	return countInserts, nil
}

func (f *DBFiller) linkProfileViews() (int, error) {
	countInserts := f.Config.Volume.CountViews

	insertStatement, countAttributes := sqltools.CreateFullQuery(insertProfileViews, countInserts)

	values := make([]interface{}, countAttributes*countInserts)

	pos := 0
	appended := 0

	for _, value := range f.faceUsers {
		count := pkg.RandMaxInt(f.Config.Volume.MaxViewOnFilm)
		if (countInserts - appended) < count {
			count = countInserts - appended
		}

		faker.Word()

		sequence := pkg.CryptoRandSequence(len(f.films)+1, 1)

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

const offset = 4.8

func (f *DBFiller) linkProfileRatings() (int, error) {
	countInserts := f.Config.Volume.CountRatings

	insertStatement, countAttributes := sqltools.CreateFullQuery(insertProfileRatings, countInserts)

	values := make([]interface{}, countAttributes*countInserts)

	pos := 0
	appended := 0

	for _, value := range f.faceUsers {
		count := pkg.RandMaxInt(f.Config.Volume.MaxCountRatingsOnFilm)
		if (countInserts - appended) < count {
			count = countInserts - appended
		}

		sequence := pkg.CryptoRandSequence(len(f.films)+1, 1)

		for j := 0; j < count; j++ {
			values[pos] = value.ID
			pos++
			values[pos] = sequence[j]
			pos++
			values[pos] = pkg.RandMaxFloat64(f.Config.Volume.MaxRatings-offset, 1) + offset
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
		return 0, errors.Wrap(err, "linkProfileRatings")
	}

	return countInserts, nil
}

func (f *DBFiller) UpdateProfiles() (int, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(f.Config.Database.Timeout)*time.Second)
	defer cancelFunc()

	rows, err := f.DB.Connection.ExecContext(ctx, updateProfiles)
	if err != nil {
		return 0, fmt.Errorf("UpdateFilms: [%w] when inserting row into [%s] table", err, updateProfiles)
	}

	affected, err := rows.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "UpdateProfiles")
	}

	return int(affected), nil
}
