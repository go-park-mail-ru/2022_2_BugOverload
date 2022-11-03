package models

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
)

type FilmActorResponse struct {
	ID        int    `json:"id,omitempty" example:"2132"`
	Name      string `json:"name,omitempty" example:"Питер Динклэйдж"`
	Avatar    string `json:"avatar,omitempty" example:"2132"`
	Character string `json:"character,omitempty" example:"Тирион Ланистер"`
}

type FilmPersonResponse struct {
	ID   int    `json:"id,omitempty" example:"123123"`
	Name string `json:"name,omitempty" example:"Стивен Спилберг"`
}

type FilmResponse struct {
	ID           int    `json:"id,omitempty" example:"23"`
	Name         string `json:"name,omitempty" example:"Игра престолов"`
	OriginalName string `json:"original_name,omitempty" example:"Game of Thrones"`
	ProdYear     int    `json:"prod_year,omitempty" example:"2011"`
	Slogan       string `json:"slogan,omitempty" example:"Победа или смерть"`
	Description  string `json:"description,omitempty" example:"Британская лингвистка Алетея прилетает из Лондона"`
	AgeLimit     int    `json:"age_limit,omitempty" example:"18"`
	Duration     int    `json:"duration,omitempty" example:"55"`
	PosterVer    string `json:"poster_ver,omitempty" example:"23"`

	Budget         int    `json:"budget,omitempty" example:"18323222"`
	BoxOffice      int    `json:"box_office,omitempty" example:"60000000"`
	CurrencyBudget string `json:"currency_budget,omitempty"  example:"USD"`

	CountSeasons int    `json:"count_seasons,omitempty" example:"8"`
	EndYear      int    `json:"end_year,omitempty" example:"2019"`
	Type         string `json:"type,omitempty" example:"serial"`

	Rating               float32 `json:"rating,omitempty" example:"9.0"`
	CountActors          int     `json:"count_actors,omitempty" example:"783"`
	CountScores          int     `json:"count_scores,omitempty" example:"786442"`
	CountNegativeReviews int     `json:"count_negative_reviews,omitempty" example:"373"`
	CountNeutralReviews  int     `json:"count_neutral_reviews,omitempty" example:"63"`
	CountPositiveReviews int     `json:"count_positive_reviews,omitempty" example:"65"`

	Genres        []string             `json:"genres,omitempty" example:"фантастика,боевик"`
	ProdCompanies []string             `json:"prod_companies,omitempty" example:"HBO"`
	ProdCountries []string             `json:"prod_countries,omitempty" example:"США,Великобритания"`
	Actors        []FilmActorResponse  `json:"actors,omitempty"`
	Artists       []FilmPersonResponse `json:"artists,omitempty"`
	Directors     []FilmPersonResponse `json:"directors,omitempty"`
	Writers       []FilmPersonResponse `json:"writers,omitempty"`
	Producers     []FilmPersonResponse `json:"producers,omitempty"`
	Operators     []FilmPersonResponse `json:"operators,omitempty"`
	Montage       []FilmPersonResponse `json:"montage,omitempty"`
	Composers     []FilmPersonResponse `json:"composers,omitempty"`
}

func NewFilmResponse(film *models.Film) *FilmResponse {
	actors := make([]FilmActorResponse, len(film.Actors))

	for idx, val := range film.Actors {
		actors[idx].ID = val.ID
		actors[idx].Name = val.Name
		actors[idx].Avatar = val.Avatar
		actors[idx].Character = val.Character
	}

	fillPersons := func(someStruct []models.FilmPerson) []FilmPersonResponse {
		persons := make([]FilmPersonResponse, len(film.Actors))

		for idx, val := range someStruct {
			persons[idx].ID = val.ID
			persons[idx].Name = val.Name
		}

		return persons
	}

	return &FilmResponse{
		ID:             film.ID,
		Name:           film.Name,
		OriginalName:   film.OriginalName,
		ProdYear:       film.ProdYear,
		Slogan:         film.Slogan,
		Description:    film.Description,
		AgeLimit:       film.AgeLimit,
		BoxOffice:      film.BoxOffice,
		Budget:         film.Budget,
		CurrencyBudget: film.CurrencyBudget,
		Duration:       film.Duration,
		PosterVer:      film.PosterVer,

		CountSeasons: film.CountSeasons,
		EndYear:      film.EndYear,
		Type:         film.Type,

		Rating:               film.Rating,
		CountActors:          film.CountActors,
		CountScores:          film.CountScores,
		CountNegativeReviews: film.CountNegativeReviews,
		CountNeutralReviews:  film.CountNeutralReviews,
		CountPositiveReviews: film.CountPositiveReviews,

		Genres:        film.Genres,
		ProdCompanies: film.ProdCompanies,
		ProdCountries: film.ProdCountries,
		Actors:        actors,
		Artists:       fillPersons(film.Artists),
		Directors:     fillPersons(film.Directors),
		Writers:       fillPersons(film.Writers),
		Producers:     fillPersons(film.Producers),
		Operators:     fillPersons(film.Operators),
		Montage:       fillPersons(film.Montage),
		Composers:     fillPersons(film.Composers),
	}
}
