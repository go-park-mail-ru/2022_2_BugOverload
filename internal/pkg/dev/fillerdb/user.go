package fillerdb

import (
	"github.com/sirupsen/logrus"

	"go-park-mail-ru/2022_2_BugOverload/pkg"
)

func (f *DBFiller) uploadUsers() (int, error) {
	countInserts := len(f.faceUsers)

	insertStatement, countAttributes := pkg.CreateStatement(insertUsers, countInserts)

	insertStatement += insertUsersEnd

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

	stmt, rows, cancelFunc, err := pkg.SendQuery(f.DB.Connection, f.Config.Database.Timeout, insertStatement, target, values)
	if err != nil {
		return 0, err
	}
	defer cancelFunc()
	defer stmt.Close()
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

	insertStatement, countAttributes := pkg.CreateStatement(insertUsersProfiles, countInserts)

	values := make([]interface{}, countAttributes*countInserts)

	for idx, value := range f.faceUsers {
		values[idx] = value.ID
	}

	stmt, rows, cancelFunc, err := pkg.SendQuery(f.DB.Connection, f.Config.Database.Timeout, insertStatement, "profiles", values)
	if err != nil {
		return 0, err
	}
	defer cancelFunc()
	defer stmt.Close()
	defer rows.Close()

	return countInserts, nil
}

func (f *DBFiller) linkProfileViews() (int, error) {
	countInserts := f.Config.Volume.CountViews

	insertStatement, countAttributes := pkg.CreateStatement(insertProfileViews, countInserts)

	values := make([]interface{}, countAttributes*countInserts)

	pos := 0
	appended := 0

	for _, value := range f.faceUsers {
		count := pkg.RandMaxInt(f.Config.Volume.MaxViewOnFilm)
		if (countInserts - appended) < count {
			count = countInserts - appended
		}

		sequence := pkg.CryptoRandSequence(f.films[len(f.films)-1].ID+1, f.films[0].ID)

		for j := 0; j < count; j++ {
			values[pos] = value.ID
			pos++
			values[pos] = sequence[j]
			pos++
		}

		appended += count
	}

	stmt, rows, cancelFunc, err := pkg.SendQuery(f.DB.Connection, f.Config.Database.Timeout, insertStatement, "profile views", values)
	if err != nil {
		return 0, err
	}
	defer cancelFunc()
	defer stmt.Close()
	defer rows.Close()

	return countInserts, nil
}

func (f *DBFiller) linkProfileRatings() (int, error) {
	countInserts := f.Config.Volume.CountRatings

	insertStatement, countAttributes := pkg.CreateStatement(insertProfileRatings, countInserts)

	values := make([]interface{}, countAttributes*countInserts)

	pos := 0
	appended := 0

	for _, value := range f.faceUsers {
		count := pkg.RandMaxInt(f.Config.Volume.MaxCountRatingsOnFilm)
		if (countInserts - appended) < count {
			count = countInserts - appended
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

		appended += count
	}

	stmt, rows, cancelFunc, err := pkg.SendQuery(f.DB.Connection, f.Config.Database.Timeout, insertStatement, "profile ratings", values)
	if err != nil {
		return 0, err
	}
	defer cancelFunc()
	defer stmt.Close()
	defer rows.Close()

	return countInserts, nil
}
