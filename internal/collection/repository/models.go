package repository

import (
	"database/sql"
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/sqltools"

	"go-park-mail-ru/2022_2_BugOverload/internal/film/repository"
)

type CollectionSQL struct {
	ID   int
	Name string
	Time string

	Description sql.NullString
	Poster      sql.NullString
	CountLikes  sql.NullInt32
	CountFilms  sql.NullInt32

	Films []repository.FilmSQL
}

func NewCollectionSQL() CollectionSQL {
	return CollectionSQL{}
}

func NewCollectionSQLOnCollection(collection models.Collection) CollectionSQL {
	return CollectionSQL{
		ID:   collection.ID,
		Name: collection.Name,
		Time: collection.Time,

		Description: sqltools.NewSQLNullString(collection.Description),
		Poster:      sqltools.NewSQLNullString(collection.Poster),
		CountLikes:  sqltools.NewSQLNullInt32(collection.CountLikes),
		CountFilms:  sqltools.NewSQLNullInt32(collection.CountFilms),
	}
}

func (c *CollectionSQL) Convert() models.Collection {
	res := models.Collection{
		ID:   c.ID,
		Name: c.Name,
		Time: c.Time,

		Description: c.Description.String,
		Poster:      c.Poster.String,
		CountLikes:  int(c.CountLikes.Int32),
		CountFilms:  int(c.CountFilms.Int32),
		Films:       make([]models.Film, len(c.Films)),
	}

	for idx := range res.Films {
		res.Films[idx] = c.Films[idx].Convert()
	}

	return res
}
