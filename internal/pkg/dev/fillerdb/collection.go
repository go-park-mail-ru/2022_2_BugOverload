package fillerdb

import (
	"context"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/dev/devpkg"
	"strings"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/pkg/errors"

	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/sqltools"
)

func (f *DBFiller) uploadCollections() (int, error) {
	// Defining sending parameters
	query := insertCollections
	message := "uploadCollections"

	globalCountInserts := len(f.collections) * len(f.Users)

	countAttributes := strings.Count(query, ",") + 1

	maxInsertValues := devpkg.MaxInsertValuesSQL / countAttributes

	action := func(countInserts int, values []interface{}) error {
		ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(f.Config.Database.Timeout)*time.Second)
		defer cancelFunc()

		insertStatement := sqltools.CreateFullQuery(query, countInserts, countAttributes)

		_, err := sqltools.InsertBatch(ctx, f.DB.Connection, insertStatement, values)
		if err != nil {
			return errors.Wrap(err, message)
		}

		return nil
	}

	pos := 0

	values := make([]interface{}, maxInsertValues*countAttributes)

	countInserts := 0

	for i := 0; i < len(f.collections); i++ {
		for j := 0; j < len(f.Users); j++ {
			values[pos] = f.collectionsSQL[i].Name
			pos++
			values[pos] = f.collectionsSQL[i].Description
			pos++
			values[pos] = f.collectionsSQL[i].Poster
			pos++
			values[pos] = faker.Timestamp()
			pos++
			values[pos] = f.collectionsSQL[i].Public
			pos++

			countInserts++

			if devpkg.MaxInsertValuesSQL-pos < 20 || j == countInserts-1 {
				err := action(countInserts, values[:pos])
				if err != nil {
					return 0, errors.Wrap(err, message)
				}

				pos = 0

				countInserts = 0
			}
		}
	}

	return globalCountInserts, nil
}

func (f *DBFiller) linkCollectionProfile() (int, error) {
	// Defining sending parameters
	query := insertProfileCollections
	message := "linkCollectionProfile"

	globalCountInserts := len(f.collections) * len(f.Users)

	countAttributes := strings.Count(query, ",") + 1

	maxInsertValues := devpkg.MaxInsertValuesSQL / countAttributes

	action := func(countInserts int, values []interface{}) error {
		ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(f.Config.Database.Timeout)*time.Second)
		defer cancelFunc()

		insertStatement := sqltools.CreateFullQuery(query, countInserts, countAttributes)

		_, err := sqltools.InsertBatch(ctx, f.DB.Connection, insertStatement, values)
		if err != nil {
			return errors.Wrap(err, message)
		}

		return nil
	}

	pos := 0

	values := make([]interface{}, maxInsertValues*countAttributes)

	countInserts := 0

	collectionID := 1

	for i := 0; i < len(f.collections); i++ {
		for j := 0; j < len(f.Users); j++ {
			values[pos] = collectionID
			collectionID++
			pos++
			values[pos] = f.Users[i].ID
			pos++

			countInserts++

			if devpkg.MaxInsertValuesSQL-pos < 20 || j == countInserts-1 {
				err := action(countInserts, values[:pos])
				if err != nil {
					return 0, errors.Wrap(err, message)
				}

				pos = 0

				countInserts = 0
			}
		}
	}

	return globalCountInserts, nil
}
