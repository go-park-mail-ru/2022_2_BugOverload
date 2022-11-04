package repository

import (
	"context"
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/sqltools"
)

// CollectionRepository provides the versatility of collection repositories.
type CollectionRepository interface {
	GetCollectionByTag(ctx context.Context) (models.Collection, error)
}

// collectionPostgres is implementation repository of collection
// Postgres DB corresponding to the CollectionService interface.
type collectionPostgres struct {
	database *sqltools.Database
}

// NewCollectionCache is constructor for collectionPostgres.
func NewCollectionCache(database *sqltools.Database) CollectionRepository {
	return &collectionPostgres{
		database,
	}
}

// GetCollectionByTag it gives away movies by tag from the repository.
func (c *collectionPostgres) GetCollectionByTag(ctx context.Context) (models.Collection, error) {
	response := NewCollectionSQL()

	//err := sqltools.RunTx(ctx, innerPKG.TxDefaultOptions, c.database.Connection, func(tx *sql.Tx) error {
	//	rowPerson := tx.QueryRowContext(ctx, getPerson, person.ID)
	//	if stdErrors.Is(rowPerson.Err(), sql.ErrNoRows) {
	//		return errors.ErrNotFoundInDB
	//	}
	//
	//	if rowPerson.Err() != nil {
	//		return rowPerson.Err()
	//	}
	//
	//	err := rowPerson.Scan(
	//		&response.Name,
	//		&response.Birthday,
	//		&response.Growth,
	//		&response.OriginalName,
	//		&response.Avatar,
	//		&response.Death,
	//		&response.Gender,
	//		&response.CountFilms)
	//	if err != nil {
	//		return err
	//	}
	//
	//	return nil
	//})
	//
	//// the main entity is not found
	//if stdErrors.Is(err, errors.ErrNotFoundInDB) {
	//	return models.Collection{}, errors.ErrNotFoundInDB
	//}
	//
	//// execution error
	//if err != nil {
	//	return models.Collection{}, errors.ErrPostgresRequest
	//}

	return response.Convert(), nil
}
