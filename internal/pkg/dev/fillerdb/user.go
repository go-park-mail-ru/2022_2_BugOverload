package fillerdb

import (
	"context"
	"database/sql"
	"go-park-mail-ru/2022_2_BugOverload/pkg"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func (f *DBFiller) uploadUsers() (int, error) {
	countInserts := len(f.faceUsers)

	insertStatement, countAttributes := getBatchInsertUsers(countInserts)

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

	target := "users"

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(f.Config.Database.Timeout)*time.Second)
	defer cancelFunc()

	stmt, err := f.DB.Connection.PrepareContext(ctx, insertStatement)
	if err != nil {
		logrus.Errorf("Error [%s] when preparing SQL statement in [%s]", err, target)
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
	for rows.Next() {
		var insertID int64
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

func (f *DBFiller) linkUsersProfiles() (int, error) {
	countInserts := len(f.faceUsers)

	insertStatement, countAttributes := getBatchInsertProfiles(countInserts)

	values := make([]interface{}, countAttributes*countInserts)

	for idx, value := range f.faceUsers {
		values[idx] = value.ID
	}

	target := "profiles"

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(f.Config.Database.Timeout)*time.Second)
	defer cancelFunc()

	stmt, err := f.DB.Connection.PrepareContext(ctx, insertStatement)
	if err != nil {
		logrus.Errorf("Error [%s] when preparing SQL statement in [%s]", err, target)
		return 0, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, values...)
	if err != nil {
		logrus.Errorf("Error [%s] when inserting row into [%s] table", err, target)
		return 0, err
	}
	defer rows.Close()

	return countInserts, nil
}

func (f *DBFiller) linkProfileViews() (int, error) {
	countInserts := f.Config.Volume.CountViews

	insertStatement, countAttributes := getBatchInsertProfileViews(countInserts)

	values := make([]interface{}, countAttributes*countInserts)

	pos := 0
	i := 0

	for _, value := range f.faceUsers {
		count := pkg.RandMaxInt(f.Config.Volume.MaxViewOnFilm)
		if (countInserts - i) < count {
			count = countInserts - i
		}

		sequence := pkg.CryptoRandSequence(f.films[len(f.films)-1].ID+1, f.films[0].ID)

		for j := 0; j < count; j++ {
			values[pos] = value.ID
			pos++
			values[pos] = sequence[j]
			pos++
		}

		i += count
	}

	target := "profile views"

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(f.Config.Database.Timeout)*time.Second)
	defer cancelFunc()

	stmt, err := f.DB.Connection.PrepareContext(ctx, insertStatement)
	if err != nil {
		logrus.Errorf("Error [%s] when preparing SQL statement in [%s]", err, target)
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

func (f *DBFiller) linkProfileRatings() (int, error) {
	countInserts := f.Config.Volume.CountRatings

	insertStatement, countAttributes := getBatchInsertProfileRatings(countInserts)

	values := make([]interface{}, countAttributes*countInserts)

	pos := 0
	i := 0

	for _, value := range f.faceUsers {
		count := pkg.RandMaxInt(f.Config.Volume.MaxCountRatingsOnFilm)
		if (countInserts - i) < count {
			count = countInserts - i
		}

		sequence := pkg.CryptoRandSequence(f.films[len(f.films)-1].ID+1, f.films[0].ID)

		for j := 0; j < count; j++ {
			values[pos] = value.ID
			pos++
			values[pos] = sequence[j]
			pos++
			values[pos] = pkg.RandMaxFloat64(f.Config.Volume.MaxRatings, 1)
			pos++
		}

		i += count
	}

	target := "profile ratings"

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(f.Config.Database.Timeout)*time.Second)
	defer cancelFunc()

	stmt, err := f.DB.Connection.PrepareContext(ctx, insertStatement)
	if err != nil {
		logrus.Errorf("Error [%s] when preparing SQL statement in [%s]", err, target)
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
