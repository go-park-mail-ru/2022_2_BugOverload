package repository

import (
	"database/sql"
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	innerPKG "go-park-mail-ru/2022_2_BugOverload/internal/pkg/sqltools"
)

type FilmSQL struct {
	ID               int
	Name             string
	ProdYear         int
	ShortDescription string
	Description      string
	Duration         int

	OriginalName sql.NullString
	Slogan       sql.NullString
	AgeLimit     sql.NullInt32
	PosterHor    sql.NullString
	PosterVer    sql.NullString

	BoxOffice      sql.NullInt32
	Budget         sql.NullInt32
	CurrencyBudget sql.NullString

	CountSeasons sql.NullInt32
	EndYear      sql.NullInt32
	Type         sql.NullString

	Rating               sql.NullFloat64
	CountScores          sql.NullInt32
	CountActors          sql.NullInt32
	CountNegativeReviews sql.NullInt32
	CountNeutralReviews  sql.NullInt32
	CountPositiveReviews sql.NullInt32
}

func NewFilmSQL(film models.Film) FilmSQL {
	return FilmSQL{
		ID:               film.ID,
		Name:             film.Name,
		ProdYear:         film.ProdYear,
		ShortDescription: film.ShortDescription,
		Description:      film.Description,
		Duration:         film.Duration,

		OriginalName: innerPKG.NewSQLNullString(film.OriginalName),
		Slogan:       innerPKG.NewSQLNullString(film.Slogan),
		AgeLimit:     innerPKG.NewSQLNullInt32(film.AgeLimit),
		PosterHor:    innerPKG.NewSQLNullString(film.PosterHor),
		PosterVer:    innerPKG.NewSQLNullString(film.PosterVer),

		BoxOffice:      innerPKG.NewSQLNullInt32(film.BoxOffice),
		Budget:         innerPKG.NewSQLNullInt32(film.Budget),
		CurrencyBudget: innerPKG.NewSQLNullString(film.CurrencyBudget),

		CountSeasons: innerPKG.NewSQLNullInt32(film.CountSeasons),
		EndYear:      innerPKG.NewSQLNullInt32(film.EndYear),
		Type:         innerPKG.NewSQLNullString(film.Type),

		Rating:               innerPKG.NewSQLNullFloat64(film.Rating),
		CountScores:          innerPKG.NewSQLNullInt32(film.CountScores),
		CountActors:          innerPKG.NewSQLNullInt32(film.CountActors),
		CountNegativeReviews: innerPKG.NewSQLNullInt32(film.CountNegativeReviews),
		CountNeutralReviews:  innerPKG.NewSQLNullInt32(film.CountNeutralReviews),
		CountPositiveReviews: innerPKG.NewSQLNullInt32(film.CountPositiveReviews),
	}
}
