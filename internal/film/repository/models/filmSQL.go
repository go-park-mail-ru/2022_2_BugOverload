package models

import (
	"database/sql"
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	sql2 "go-park-mail-ru/2022_2_BugOverload/internal/pkg"
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

		OriginalName: sql2.NewSQLNullString(film.OriginalName),
		Slogan:       sql2.NewSQLNullString(film.Slogan),
		AgeLimit:     sql2.NewSQLNullInt32(film.AgeLimit),
		PosterHor:    sql2.NewSQLNullString(film.PosterHor),
		PosterVer:    sql2.NewSQLNullString(film.PosterVer),

		BoxOffice:      sql2.NewSQLNullInt32(film.BoxOffice),
		Budget:         sql2.NewSQLNullInt32(film.Budget),
		CurrencyBudget: sql2.NewSQLNullString(film.CurrencyBudget),

		CountSeasons: sql2.NewSQLNullInt32(film.CountSeasons),
		EndYear:      sql2.NewSQLNullInt32(film.EndYear),
		Type:         sql2.NewSQLNullString(film.Type),

		Rating:               sql2.NewSQLNullFloat64(film.Rating),
		CountScores:          sql2.NewSQLNullInt32(film.CountScores),
		CountActors:          sql2.NewSQLNullInt32(film.CountActors),
		CountNegativeReviews: sql2.NewSQLNullInt32(film.CountNegativeReviews),
		CountNeutralReviews:  sql2.NewSQLNullInt32(film.CountNeutralReviews),
		CountPositiveReviews: sql2.NewSQLNullInt32(film.CountPositiveReviews),
	}
}
