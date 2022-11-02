package fillerdb

import (
	"context"
	"time"

	"github.com/pkg/errors"

	pkgInner "go-park-mail-ru/2022_2_BugOverload/internal/pkg"
)

func (f *DBFiller) uploadPersons() (int, error) {
	countInserts := len(f.personsSQL)

	insertStatement, countAttributes := pkgInner.CreateStatement(insertPersons, countInserts)

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

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(f.Config.Database.Timeout)*time.Second)
	defer cancelFunc()

	rows, err := pkgInner.SendQuery(ctx, f.DB.Connection, insertStatement, values)
	if err != nil {
		return 0, errors.Wrap(err, "uploadPersons")
	}
	defer rows.Close()

	count := 1
	for idx := range f.persons {
		f.persons[idx].ID = count
		count++
	}

	return countInserts, nil
}

func (f *DBFiller) linkPersonProfession() (int, error) {
	countInserts := 0

	for _, value := range f.persons {
		countInserts += len(value.Professions)
	}

	insertStatement, countAttributes := pkgInner.CreateStatement(insertPersonsProfessions, countInserts)

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

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(f.Config.Database.Timeout)*time.Second)
	defer cancelFunc()

	rows, err := pkgInner.SendQuery(ctx, f.DB.Connection, insertStatement, values)
	if err != nil {
		return 0, errors.Wrap(err, "linkPersonProfession")
	}
	defer rows.Close()

	return countInserts, nil
}

func (f *DBFiller) linkPersonGenres() (int, error) {
	countInserts := 0

	for _, value := range f.persons {
		countInserts += len(value.Genres)
	}

	insertStatement, countAttributes := pkgInner.CreateStatement(insertPersonsGenres, countInserts)

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

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(f.Config.Database.Timeout)*time.Second)
	defer cancelFunc()

	rows, err := pkgInner.SendQuery(ctx, f.DB.Connection, insertStatement, values)
	if err != nil {
		return 0, errors.Wrap(err, "linkPersonGenres")
	}
	defer rows.Close()

	return countInserts, nil
}
