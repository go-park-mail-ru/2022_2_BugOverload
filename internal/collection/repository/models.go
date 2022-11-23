package repository

import (
	"database/sql"
	"time"

	"go-park-mail-ru/2022_2_BugOverload/internal/film/repository"
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	innerPKG "go-park-mail-ru/2022_2_BugOverload/internal/pkg"
)

type CollectionSQL struct {
	ID         int
	Name       string
	Time       string
	UpdateTime time.Time
	CreateTime time.Time

	Description sql.NullString
	Poster      sql.NullString
	CountLikes  sql.NullInt32
	CountFilms  sql.NullInt32

	Films []repository.FilmSQL
}

func NewCollectionSQL() CollectionSQL {
	return CollectionSQL{}
}

func (c *CollectionSQL) Convert() models.Collection {
	res := models.Collection{
		ID:         c.ID,
		Name:       c.Name,
		Time:       c.Time,
		UpdateTime: c.UpdateTime.Format(innerPKG.DateFormat + " " + innerPKG.TimeFormat),
		CreateTime: c.CreateTime.Format(innerPKG.DateFormat + " " + innerPKG.TimeFormat),

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
