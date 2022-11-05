package fillerdb

import (
	"context"
	"fmt"
	pkgInner "go-park-mail-ru/2022_2_BugOverload/internal/pkg/sqltools"
	"strings"
	"time"

	"github.com/pkg/errors"
)

func (f *DBFiller) uploadPersons() (int, error) {
	countInserts := len(f.personsSQL)

	insertStatement, countAttributes := pkgInner.CreateFullQuery(insertPersons, countInserts)

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

	rows, err := pkgInner.InsertBatch(ctx, f.DB.Connection, insertStatement, values)
	if err != nil {
		return 0, errors.Wrap(err, "uploadPersons")
	}

	affected, err := rows.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "uploadPersons")
	}

	for i := 0; i < int(affected); i++ {
		f.persons[i].ID = i + 1
	}

	return countInserts, nil
}

func (f *DBFiller) linkPersonProfession() (int, error) {
	countInserts := 0

	for _, value := range f.persons {
		countInserts += len(value.Professions)
	}

	insertStatement, countAttributes := pkgInner.CreateFullQuery(insertPersonsProfessions, countInserts)

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

	_, err := pkgInner.InsertBatch(ctx, f.DB.Connection, insertStatement, values)
	if err != nil {
		return 0, errors.Wrap(err, "linkPersonProfession")
	}

	return countInserts, nil
}

func (f *DBFiller) linkPersonGenres() (int, error) {
	countInserts := 0

	for _, value := range f.persons {
		countInserts += len(value.Genres)
	}

	insertStatement, countAttributes := pkgInner.CreateFullQuery(insertPersonsGenres, countInserts)

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

	_, err := pkgInner.InsertBatch(ctx, f.DB.Connection, insertStatement, values)
	if err != nil {
		return 0, errors.Wrap(err, "linkPersonGenres")
	}

	return countInserts, nil
}

func (f *DBFiller) linkPersonImages() (int, error) {
	countInserts := len(f.persons)

	insertStatement, countAttributes := pkgInner.CreateFullQuery(insertPersonsImages, countInserts)

	values := make([]interface{}, countAttributes*countInserts)

	pos := 0
	for _, value := range f.persons {
		values[pos] = value.ID
		pos++

		imagesList := strings.Join(value.Images, "_")
		values[pos] = imagesList
		pos++
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(f.Config.Database.Timeout)*time.Second)
	defer cancelFunc()

	_, err := pkgInner.InsertBatch(ctx, f.DB.Connection, insertStatement, values)
	if err != nil {
		return 0, errors.Wrap(err, "linkPersonImages")
	}

	return countInserts, nil
}

func (f *DBFiller) UpdatePersons() (int, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(f.Config.Database.Timeout)*time.Second)
	defer cancelFunc()

	rows, err := f.DB.Connection.ExecContext(ctx, updatePersons)
	if err != nil {
		return 0, fmt.Errorf("UpdateFilms: [%w] when inserting row into [%s] table", err, updatePersons)
	}

	affected, err := rows.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "UpdatePersons")
	}

	return int(affected), nil
}
