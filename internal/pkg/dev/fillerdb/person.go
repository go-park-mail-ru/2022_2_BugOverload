package fillerdb

import (
	"github.com/sirupsen/logrus"

	"go-park-mail-ru/2022_2_BugOverload/pkg"
)

func (f *DBFiller) uploadPersons() (int, error) {
	countInserts := len(f.personsSQL)

	insertStatement, countAttributes := pkg.CreateStatement(insertPersons, countInserts)

	insertStatement += insertPersonsEnd

	values := make([]interface{}, countAttributes*countInserts)

	for idx, value := range f.personsSQL {
		posAttr := 0
		posValue := idx * countAttributes

		values[posValue+posAttr] = value.Name
		posAttr++
		values[posValue+posAttr] = value.OriginalName
		posAttr++
		values[posValue+posAttr] = value.Birthday
		posAttr++
		values[posValue+posAttr] = value.Growth
		posAttr++
		values[posValue+posAttr] = value.Avatar
		posAttr++
		values[posValue+posAttr] = value.Gender
		posAttr++
		values[posValue+posAttr] = value.Death
	}

	target := "persons"

	stmt, rows, cancelFunc, err := pkg.SendQuery(f.DB.Connection, f.Config.Database.Timeout, insertStatement, target, values)
	if err != nil {
		return 0, err
	}
	defer cancelFunc()
	defer stmt.Close()
	defer rows.Close()

	counter := 0
	var insertID int64
	for rows.Next() {
		err = rows.Scan(&insertID)
		if err != nil {
			logrus.Errorf("Error [%s] when getting insertID [%s]", err, target)
			return 0, err
		}

		f.persons[counter].ID = int(insertID)
		counter++
	}

	return countInserts, nil
}

func (f *DBFiller) linkPersonProfession() (int, error) {
	countInserts := 0

	for _, value := range f.persons {
		countInserts += len(value.Professions)
	}

	insertStatement, countAttributes := pkg.CreateStatement(insertPersonsProfessions, countInserts)

	values := make([]interface{}, countAttributes*countInserts)

	offset := 0
	posValue := 0

	for _, value := range f.persons {
		posValue += offset
		offset = 0
		weight := len(value.Professions)

		for _, profession := range value.Professions {
			values[posValue+offset] = value.ID
			offset++
			values[posValue+offset] = f.professions[profession]
			weight--
			offset++
			values[posValue+offset] = weight
			offset++
		}
	}

	stmt, rows, cancelFunc, err := pkg.SendQuery(f.DB.Connection, f.Config.Database.Timeout, insertStatement, "person professions", values)
	if err != nil {
		return 0, err
	}
	defer cancelFunc()
	defer stmt.Close()
	defer rows.Close()

	return countInserts, nil
}

func (f *DBFiller) linkPersonGenres() (int, error) {
	countInserts := 0

	for _, value := range f.persons {
		countInserts += len(value.Genres)
	}

	insertStatement, countAttributes := pkg.CreateStatement(insertPersonsGenres, countInserts)

	values := make([]interface{}, countAttributes*countInserts)

	offset := 0
	posValue := 0

	for _, value := range f.persons {
		posValue += offset
		offset = 0
		weight := len(value.Genres)

		for _, genre := range value.Genres {
			values[posValue+offset] = value.ID
			offset++
			values[posValue+offset] = f.genres[genre]
			weight--
			offset++
			values[posValue+offset] = weight
			offset++
		}
	}

	stmt, rows, cancelFunc, err := pkg.SendQuery(f.DB.Connection, f.Config.Database.Timeout, insertStatement, "person genres", values)
	if err != nil {
		return 0, err
	}
	defer cancelFunc()
	defer stmt.Close()
	defer rows.Close()
	return countInserts, nil
}
