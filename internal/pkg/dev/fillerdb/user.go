package fillerdb

import (
	"context"
	"database/sql"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func (f *DBFiller) UploadUsers() (int, error) {
	countInserts := len(f.faceUsers)

	insertStatement, countAttributes := GetBatchInsertUsers(countInserts)

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

func (f *DBFiller) LinkUsersProfiles() (int, error) {
	countInserts := len(f.faceUsers)

	insertStatement, countAttributes := GetBatchInsertProfiles(countInserts)

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
