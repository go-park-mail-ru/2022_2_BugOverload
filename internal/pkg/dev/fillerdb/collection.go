package fillerdb

import (
	"context"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/pkg/errors"

	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/sqltools"
)

func (f *DBFiller) uploadCollections() (int, error) {
	countInserts := len(f.collections) * len(f.faceUsers)

	insertStatement, countAttributes := sqltools.CreateFullQuery(insertCollections, countInserts)

	values := make([]interface{}, countAttributes*countInserts)

	pos := 0

	for _, value := range f.collectionsSQL {
		for range f.faceUsers {
			values[pos] = value.Name
			pos++
			values[pos] = value.Description
			pos++
			values[pos] = value.Poster
			pos++
			values[pos] = faker.Timestamp()
			pos++
		}
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(f.Config.Database.Timeout)*time.Second)
	defer cancelFunc()

	_, err := sqltools.InsertBatch(ctx, f.DB.Connection, insertStatement, values)
	if err != nil {
		return 0, errors.Wrap(err, "insertCollections")
	}

	return countInserts, nil
}

func (f *DBFiller) linkCollectionProfile() (int, error) {
	countInserts := len(f.collections) * len(f.faceUsers)

	insertStatement, countAttributes := sqltools.CreateFullQuery(insertProfileCollections, countInserts)

	values := make([]interface{}, countAttributes*countInserts)

	pos := 0
	collectionID := 1

	for range f.collectionsSQL {
		for _, user := range f.faceUsers {
			values[pos] = collectionID
			collectionID++
			pos++
			values[pos] = user.ID
			pos++
		}
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(f.Config.Database.Timeout)*time.Second)
	defer cancelFunc()

	_, err := sqltools.InsertBatch(ctx, f.DB.Connection, insertStatement, values)
	if err != nil {
		return 0, errors.Wrap(err, "linkCollectionProfile")
	}

	return countInserts, nil
}
