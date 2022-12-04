package collection

import (
	"database/sql"
	innerPKG "go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	"go-park-mail-ru/2022_2_BugOverload/internal/warehouse/repository/film"
	"time"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
)

type AuthorSQL struct {
	ID       int
	Nickname sql.NullString
}

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

	Author AuthorSQL
}

func NewCollectionSQL() ModelSQL {
	return ModelSQL{}
}

func (c *ModelSQL) Convert() models.Collection {
	updateTime := ""

	if !c.UpdateTime.Equal(time.Time{}) {
		updateTime = c.UpdateTime.Format(innerPKG.DateFormat + " " + innerPKG.TimeFormat)
	}

	createTime := ""

	if !c.CreateTime.Equal(time.Time{}) {
		createTime = c.CreateTime.Format(innerPKG.DateFormat + " " + innerPKG.TimeFormat)
	}

	res := models.Collection{
		ID:         c.ID,
		Name:       c.Name,
		Time:       c.Time,
		UpdateTime: updateTime,
		CreateTime: createTime,

		Description: c.Description.String,
		Poster:      c.Poster.String,
		CountLikes:  int(c.CountLikes.Int32),
		CountFilms:  int(c.CountFilms.Int32),
		Films:       make([]models.Film, len(c.Films)),

		Author: models.User{
			ID:       c.Author.ID,
			Nickname: c.Author.Nickname.String,
		},
	}

	for idx := range res.Films {
		res.Films[idx] = c.Films[idx].Convert()
	}

	return res
}
