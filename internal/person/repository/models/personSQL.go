package models

import (
	"database/sql"
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/pkg"
)

type PersonSQL struct {
	ID       int
	Name     string
	Birthday string
	Growth   float32

	Avatar       sql.NullString
	Gender       sql.NullString
	CountFilms   sql.NullInt32
	OriginalName sql.NullString
	Death        sql.NullString
}

func NewPersonSQL(person models.Person) PersonSQL {
	return PersonSQL{
		ID:       person.ID,
		Name:     person.Name,
		Birthday: person.Birthday,
		Growth:   person.Growth,

		Avatar:       pkg.NewSQLNullString(person.Avatar),
		OriginalName: pkg.NewSQLNullString(person.OriginalName),
		Gender:       pkg.NewSQLNullString(person.Gender),
		Death:        pkg.NewSQLNullString(person.Death),
		CountFilms:   pkg.NewSQLNullInt32(person.CountFilms),
	}
}
