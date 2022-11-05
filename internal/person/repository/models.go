package repository

import (
	"database/sql"
	"go-park-mail-ru/2022_2_BugOverload/internal/film/repository"
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/sqltools"
	"time"
)

type PersonSQL struct {
	ID       int
	Name     string
	Birthday time.Time
	Growth   float32

	Avatar       sql.NullString
	Gender       sql.NullString
	CountFilms   sql.NullInt32
	OriginalName sql.NullString
	Death        sql.NullTime

	BestFilms []repository.FilmSQL

	Images []string
}

func NewPersonSQL() PersonSQL {
	return PersonSQL{}
}

func NewPersonSQLOnPerson(person models.Person) PersonSQL {
	birthday := time.Time{}

	if person.Birthday != "" {
		var err error
		birthday, err = time.Parse("2006-04-02", person.Birthday)
		if err != nil {
			birthday = time.Time{}
		}
	}

	return PersonSQL{
		ID:       person.ID,
		Name:     person.Name,
		Birthday: birthday,
		Growth:   person.Growth,

		Avatar:       sqltools.NewSQLNullString(person.Avatar),
		OriginalName: sqltools.NewSQLNullString(person.OriginalName),
		Gender:       sqltools.NewSQLNullString(person.Gender),
		Death:        sqltools.NewSQLNNullDate(person.Death),
		CountFilms:   sqltools.NewSQLNullInt32(person.CountFilms),
	}
}

func (p *PersonSQL) Convert() models.Person {
	death := ""
	if p.Death.Valid {
		death = p.Death.Time.Format("2006-01-02")
	}

	res := models.Person{
		ID:       p.ID,
		Name:     p.Name,
		Birthday: p.Birthday.Format("2006-01-02"),
		Growth:   p.Growth,

		Avatar:       p.Avatar.String,
		Gender:       p.Gender.String,
		CountFilms:   int(p.CountFilms.Int32),
		OriginalName: p.OriginalName.String,
		Death:        death,
		BestFilms:    make([]models.Film, len(p.BestFilms)),

		Images: p.Images,
	}

	for idx := range res.BestFilms {
		res.BestFilms[idx] = p.BestFilms[idx].Convert()
	}

	return res
}
