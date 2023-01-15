package fillerdb

import (
	"context"
	"strings"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/pkg/errors"

	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/sqltools"
)

func (f *DBFiller) uploadCollections() (int, error) {
	countInserts := len(f.collections) * len(f.Users)

	countAttributes := strings.Count(insertCollections, ",") + 1

	insertStatement := sqltools.CreateFullQuery(insertCollections, countInserts, countAttributes)

	values := make([]interface{}, countAttributes*countInserts)

	pos := 0

	for _, value := range f.collectionsSQL {
		for range f.Users {
			values[pos] = value.Name
			pos++
			values[pos] = value.Description
			pos++
			values[pos] = value.Poster
			pos++
			values[pos] = faker.Timestamp()
			pos++
			values[pos] = value.Public
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
	countInserts := len(f.collections) * len(f.Users)

	countAttributes := strings.Count(insertProfileCollections, ",") + 1

	insertStatement := sqltools.CreateFullQuery(insertProfileCollections, countInserts, countAttributes)

	values := make([]interface{}, countAttributes*countInserts)

	pos := 0
	collectionID := 1

	for range f.collectionsSQL {
		for _, user := range f.Users {
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
