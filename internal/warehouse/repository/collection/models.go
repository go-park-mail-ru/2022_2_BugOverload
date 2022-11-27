package collection

import (
	"database/sql"
	innerPKG "go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	"go-park-mail-ru/2022_2_BugOverload/internal/warehouse/repository/film"
	"time"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
)

type ModelSQL struct {
	ID         int
	Name       string
	Time       string
	UpdateTime time.Time
	CreateTime time.Time

	Description sql.NullString
	Poster      sql.NullString
	CountLikes  sql.NullInt32
	CountFilms  sql.NullInt32

	Films []film.ModelSQL
}

func NewCollectionSQL() ModelSQL {
	return ModelSQL{}
}

func (c *ModelSQL) Convert() models.Collection {
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
