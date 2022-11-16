package repository

import (
	"database/sql"
	"time"

	"go-park-mail-ru/2022_2_BugOverload/internal/film/repository"
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	innerPKG "go-park-mail-ru/2022_2_BugOverload/internal/pkg"
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

	Images      []string
	Professions []string
	Genres      []string
}

func NewPersonSQL() PersonSQL {
	return PersonSQL{}
}

func (p *PersonSQL) Convert() models.Person {
	death := ""
	if p.Death.Valid {
		death = p.Death.Time.Format(innerPKG.DateFormat)
	}

	res := models.Person{
		ID:       p.ID,
		Name:     p.Name,
		Birthday: p.Birthday.Format(innerPKG.DateFormat),
		Growth:   p.Growth,

		Avatar:       p.Avatar.String,
		Gender:       p.Gender.String,
		CountFilms:   int(p.CountFilms.Int32),
		OriginalName: p.OriginalName.String,
		Death:        death,
		BestFilms:    make([]models.Film, len(p.BestFilms)),

		Images:      p.Images,
		Professions: p.Professions,
		Genres:      p.Genres,
	}

	for idx := range res.BestFilms {
		res.BestFilms[idx] = p.BestFilms[idx].Convert()
	}

	return res
}
