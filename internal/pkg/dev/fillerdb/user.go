package fillerdb

import (
	"context"
	pkgInner "go-park-mail-ru/2022_2_BugOverload/internal/pkg/sqltools"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/pkg/errors"

	"go-park-mail-ru/2022_2_BugOverload/pkg"
)

func (f *DBFiller) uploadUsers() (int, error) {
	countInserts := len(f.faceUsers)

	insertStatement, countAttributes := pkgInner.CreateFullQuery(insertUsers, countInserts)

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

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(f.Config.Database.Timeout)*time.Second)
	defer cancelFunc()

	rows, err := pkgInner.InsertBatch(ctx, f.DB.Connection, insertStatement, values)
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

	insertStatement, countAttributes := pkgInner.CreateFullQuery(insertUsersProfiles, countInserts)

	values := make([]interface{}, countAttributes*countInserts)

	for idx, value := range f.faceUsers {
		values[idx] = value.ID
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(f.Config.Database.Timeout)*time.Second)
	defer cancelFunc()

	_, err := pkgInner.InsertBatch(ctx, f.DB.Connection, insertStatement, values)
	if err != nil {
		return 0, errors.Wrap(err, "linkUsersProfiles")
	}

	return countInserts, nil
}

func (f *DBFiller) linkProfileViews() (int, error) {
	countInserts := f.Config.Volume.CountViews

	insertStatement, countAttributes := pkgInner.CreateFullQuery(insertProfileViews, countInserts)

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

	_, err := pkgInner.InsertBatch(ctx, f.DB.Connection, insertStatement, values)
	if err != nil {
		return 0, errors.Wrap(err, "linkProfileViews")
	}

	return countInserts, nil
}

func (f *DBFiller) linkProfileRatings() (int, error) {
	countInserts := f.Config.Volume.CountRatings

	insertStatement, countAttributes := pkgInner.CreateFullQuery(insertProfileRatings, countInserts)

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
			values[pos] = pkg.RandMaxFloat64(f.Config.Volume.MaxRatings, 1)
			pos++
			values[pos] = faker.Timestamp()
			pos++
		}

		appended += count
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(f.Config.Database.Timeout)*time.Second)
	defer cancelFunc()

	_, err := pkgInner.InsertBatch(ctx, f.DB.Connection, insertStatement, values)
	if err != nil {
		return 0, errors.Wrap(err, "linkProfileRatings")
	}

	return countInserts, nil
}
