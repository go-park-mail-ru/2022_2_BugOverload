package models

import (
	innerPKG "go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
)

type FilmRequest struct {
	FilmID      int
	CountImages int
}

func NewFilmRequest() FilmRequest {
	return FilmRequest{}
}

func (f *FilmRequest) Bind(r *http.Request) error {
	var err error

	vars := mux.Vars(r)

	f.FilmID, _ = strconv.Atoi(vars["id"])

	f.CountImages, err = strconv.Atoi(r.FormValue("count_images"))
	if err != nil {
		return errors.ErrConvertQueryType
	}

	if f.CountImages < 0 {
		return errors.ErrBadQueryParams
	}

	return nil
}

func (f *FilmRequest) GetParams() *innerPKG.GetFilmParams {
	return &innerPKG.GetFilmParams{
		CountImages: f.CountImages,
	}
}

func (f *FilmRequest) GetFilm() *models.Film {
	return &models.Film{
		ID: f.FilmID,
	}
}

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
	Name             string `json:"name,omitempty" example:"Игра престолов"`
	OriginalName     string `json:"original_name,omitempty" example:"Game of Thrones"`
	ProdYear         string `json:"prod_year,omitempty" example:"2011"`
	Slogan           string `json:"slogan,omitempty" example:"Победа или смерть"`
	Description      string `json:"description,omitempty" example:"Британская лингвистка Алетея прилетает из Лондона"`
	ShortDescription string `json:"short_description,omitempty" example:"Что вы знаете о джинах кроме желайний?"`
<<<<<<< HEAD
	AgeLimit         string `json:"age_limit,omitempty" example:"18+"`
	DurationMinutes  int    `json:"duration_minutes,omitempty" example:"55"`
=======
	AgeLimit         string `json:"age_limit,omitempty" example:"18"`
	Duration         int    `json:"duration,omitempty" example:"55"`
>>>>>>> main
	PosterHor        string `json:"poster_hor,omitempty" example:"23"`

	Budget           int    `json:"budget,omitempty" example:"18323222"`
	BoxOfficeDollars int    `json:"box_office_dollars,omitempty" example:"60000000"`
	CurrencyBudget   string `json:"currency_budget,omitempty"  example:"USD"`

	CountSeasons int    `json:"count_seasons,omitempty" example:"8"`
	EndYear      string `json:"end_year,omitempty" example:"2019"`
	Type         string `json:"type,omitempty" example:"serial"`

	Rating               float32 `json:"rating,omitempty" example:"9.2"`
	CountActors          int     `json:"count_actors,omitempty" example:"783"`
	CountRatings         int     `json:"count_ratings,omitempty" example:"786442"`
	CountNegativeReviews int     `json:"count_negative_reviews,omitempty" example:"373"`
	CountNeutralReviews  int     `json:"count_neutral_reviews,omitempty" example:"63"`
	CountPositiveReviews int     `json:"count_positive_reviews,omitempty" example:"65"`

	Images        []string             `json:"images,omitempty" example:"1,2,3,4"`
	Tags          []string             `json:"tags,omitempty" example:"популярное,сейчас в кино"`
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
		persons := make([]FilmPersonResponse, len(someStruct))

		for idx, val := range someStruct {
			persons[idx].ID = val.ID
			persons[idx].Name = val.Name
		}

		return persons
	}

	return &FilmResponse{
		Name:             film.Name,
		OriginalName:     film.OriginalName,
		ProdYear:         film.ProdYear,
		Slogan:           film.Slogan,
		Description:      film.Description,
		ShortDescription: film.ShortDescription,
		AgeLimit:         film.AgeLimit,
		BoxOfficeDollars: film.BoxOfficeDollars,
		Budget:           film.Budget,
		CurrencyBudget:   film.CurrencyBudget,
		DurationMinutes:  film.DurationMinutes,
		PosterHor:        film.PosterHor,

		CountSeasons: film.CountSeasons,
		EndYear:      film.EndYear,
		Type:         film.Type,

		Rating:               film.Rating,
		CountActors:          film.CountActors,
<<<<<<< HEAD
		CountRatings:         film.CountRatings,
=======
		CountScores:          film.CountRatings,
>>>>>>> main
		CountNegativeReviews: film.CountNegativeReviews,
		CountNeutralReviews:  film.CountNeutralReviews,
		CountPositiveReviews: film.CountPositiveReviews,

		Images:        film.Images,
		Tags:          film.Tags,
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
