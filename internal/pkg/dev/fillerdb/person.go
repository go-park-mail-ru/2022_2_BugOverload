package fillerdb

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/pkg/errors"

	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/sqltools"
)

func (f *DBFiller) uploadPersons() (int, error) {
	countInserts := len(f.personsSQL)

	countAttributes := strings.Count(insertPersons, ",") + 1

	insertStatement := sqltools.CreateFullQuery(insertPersons, countInserts, countAttributes)

	values := make([]interface{}, countAttributes*countInserts)

	pos := 0

	for _, value := range f.personsSQL {
		values[pos] = value.Name
		pos++
		values[pos] = value.OriginalName
		pos++
		values[pos] = value.Birthday
		pos++
		values[pos] = value.Growth
		pos++
		values[pos] = value.Avatar
		pos++
		values[pos] = value.Gender
		pos++
		values[pos] = value.Death
		pos++
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(f.Config.Database.Timeout)*time.Second)
	defer cancelFunc()

	rows, err := sqltools.InsertBatch(ctx, f.DB.Connection, insertStatement, values)
	if err != nil {
		return 0, errors.Wrap(err, "uploadPersons")
	}

	affected, err := rows.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "uploadPersons")
	}

	for i := 0; i < int(affected); i++ {
		f.persons[i].ID = i + 1

		f.mapPersons[f.persons[i].Name] = f.persons[i].ID
	}

	return countInserts, nil
}

func (f *DBFiller) linkPersonProfession() (int, error) {
	countInserts := 0

	for _, value := range f.persons {
		countInserts += len(value.Professions)
	}

	countAttributes := strings.Count(insertPersonsProfessions, ",") + 1

	insertStatement := sqltools.CreateFullQuery(insertPersonsProfessions, countInserts, countAttributes)

	values := make([]interface{}, countAttributes*countInserts)

	pos := 0

	for _, value := range f.persons {
		weight := len(value.Professions)

		for _, profession := range value.Professions {
			values[pos] = value.ID
			pos++
			values[pos] = f.guides.Professions[profession]
			weight--
			pos++
			values[pos] = weight
			pos++
		}
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(f.Config.Database.Timeout)*time.Second)
	defer cancelFunc()

	_, err := sqltools.InsertBatch(ctx, f.DB.Connection, insertStatement, values)
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

	countAttributes := strings.Count(insertPersonsGenres, ",") + 1

	insertStatement := sqltools.CreateFullQuery(insertPersonsGenres, countInserts, countAttributes)

	values := make([]interface{}, countAttributes*countInserts)

	pos := 0

	for _, value := range f.persons {
		weight := len(value.Genres)

		for _, genre := range value.Genres {
			values[pos] = value.ID
			pos++
			values[pos] = f.guides.Genres[genre]
			weight--
			pos++
			values[pos] = weight
			pos++
		}
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(f.Config.Database.Timeout)*time.Second)
	defer cancelFunc()

	_, err := sqltools.InsertBatch(ctx, f.DB.Connection, insertStatement, values)
	if err != nil {
		return 0, errors.Wrap(err, "linkPersonGenres")
	}

	return countInserts, nil
}

func (f *DBFiller) linkPersonImages() (int, error) {
	countInserts := 0

	for _, value := range f.persons {
		countInserts += len(value.Images)
	}

	countAttributes := strings.Count(insertPersonsImages, ",") + 1

	insertStatement := sqltools.CreateFullQuery(insertPersonsImages, countInserts, countAttributes)

	values := make([]interface{}, countAttributes*countInserts)

	pos := 0

	for _, value := range f.persons {
		weight := len(value.Images)

		for _, image := range value.Images {
			values[pos] = value.ID
			pos++
			values[pos] = image
			weight--
			pos++
			values[pos] = weight
			pos++
		}
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(f.Config.Database.Timeout)*time.Second)
	defer cancelFunc()

	_, err := sqltools.InsertBatch(ctx, f.DB.Connection, insertStatement, values)
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
