package models

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	proto "go-park-mail-ru/2022_2_BugOverload/internal/warehouse/delivery/grpc/protobuf"
)

// User
func NewUserProto(user *models.User) *proto.User {
	return &proto.User{
		ID:               uint32(user.ID),
		Nickname:         user.Nickname,
		Email:            user.Email,
		Password:         user.Password,
		IsAdmin:          user.IsAdmin,
		Avatar:           user.Avatar,
		JoinedDate:       user.JoinedDate,
		CountViewsFilms:  uint32(user.CountViewsFilms),
		CountCollections: uint32(user.CountCollections),
		CountReviews:     uint32(user.CountReviews),
		CountRatings:     uint32(user.CountRatings),
	}
}

func NewUser(user *proto.User) models.User {
	return models.User{
		ID:               int(user.ID),
		Nickname:         user.Nickname,
		Email:            user.Email,
		Password:         user.Password,
		IsAdmin:          user.IsAdmin,
		Avatar:           user.Avatar,
		JoinedDate:       user.JoinedDate,
		CountViewsFilms:  int(user.CountViewsFilms),
		CountCollections: int(user.CountCollections),
		CountReviews:     int(user.CountReviews),
		CountRatings:     int(user.CountRatings),
	}
}

// Review
func NewReviewProto(review *models.Review) *proto.Review {
	author := NewUserProto(&review.Author)

	return &proto.Review{
		ID:         uint32(review.ID),
		Name:       review.Name,
		Type:       review.Type,
		Body:       review.Body,
		CountLikes: uint32(review.CountLikes),
		CreateTime: review.CreateTime,
		Author:     author,
	}
}

func NewReview(review *proto.Review) models.Review {
	return models.Review{
		ID:         int(review.ID),
		Name:       review.Name,
		Type:       review.Type,
		Body:       review.Body,
		CountLikes: int(review.CountLikes),
		CreateTime: review.CreateTime,
		Author:     NewUser(review.Author),
	}
}

func NewReviewsProto(reviews []models.Review) *proto.Reviews {
	res := make([]*proto.Review, len(reviews))

	for idx, value := range reviews {
		res[idx] = NewReviewProto(&value)
	}

	return &proto.Reviews{
		Reviews: res,
	}
}

func NewReviews(reviews *proto.Reviews) []models.Review {
	res := make([]models.Review, len(reviews.Reviews))

	for idx, value := range reviews.Reviews {
		res[idx] = NewReview(value)
	}

	return res
}

// Film
func NewFilmProto(film *models.Film) *proto.Film {
	actors := make([]*proto.FilmActor, len(film.Actors))

	for idx, val := range film.Actors {
		actors[idx] = &proto.FilmActor{
			ID:        uint32(val.ID),
			Name:      val.Name,
			Avatar:    val.Avatar,
			Character: val.Character,
		}
	}

	fillPersons := func(someStruct []models.FilmPerson) []*proto.FilmPerson {
		persons := make([]*proto.FilmPerson, len(someStruct))

		for idx, val := range someStruct {
			persons[idx] = &proto.FilmPerson{
				ID:   uint32(val.ID),
				Name: val.Name,
			}
		}

		return persons
	}

	return &proto.Film{
		ID:                   uint32(film.ID),
		Name:                 film.Name,
		OriginalName:         film.OriginalName,
		ProdDate:             film.ProdDate,
		Slogan:               film.Slogan,
		ShortDescription:     film.ShortDescription,
		Description:          film.Description,
		AgeLimit:             film.AgeLimit,
		DurationMinutes:      uint32(film.DurationMinutes),
		PosterHor:            film.PosterHor,
		PosterVer:            film.PosterVer,
		Ticket:               film.Ticket,
		Trailer:              film.Trailer,
		BoxOfficeDollars:     uint32(film.BoxOfficeDollars),
		Budget:               uint32(film.Budget),
		CurrencyBudget:       film.CurrencyBudget,
		CountSeasons:         uint32(film.CountSeasons),
		EndYear:              film.EndYear,
		Type:                 film.Type,
		Rating:               film.Rating,
		CountRatings:         uint32(film.CountRatings),
		CountActors:          uint32(film.CountActors),
		CountPositiveReviews: uint32(film.CountPositiveReviews),
		CountNegativeReviews: uint32(film.CountNegativeReviews),
		CountNeutralReviews:  uint32(film.CountNeutralReviews),
		Tags:                 film.Tags,
		Genres:               film.Genres,
		ProdCompanies:        film.ProdCompanies,
		ProdCountries:        film.ProdCountries,
		Actors:               actors,
		Artists:              fillPersons(film.Artists),
		Directors:            fillPersons(film.Directors),
		Writers:              fillPersons(film.Writers),
		Producers:            fillPersons(film.Producers),
		Operators:            fillPersons(film.Operators),
		Montage:              fillPersons(film.Montage),
		Composers:            fillPersons(film.Composers),
	}
}

func NewFilm(film *proto.Film) models.Film {
	actors := make([]models.FilmActor, len(film.Actors))

	for idx, val := range film.Actors {
		actors[idx].ID = int(val.ID)
		actors[idx].Name = val.Name
		actors[idx].Avatar = val.Avatar
		actors[idx].Character = val.Character
	}

	fillPersons := func(someStruct []*proto.FilmPerson) []models.FilmPerson {
		persons := make([]models.FilmPerson, len(someStruct))

		for idx, val := range someStruct {
			persons[idx].ID = int(val.ID)
			persons[idx].Name = val.Name
		}

		return persons
	}

	return models.Film{
		ID:                   int(film.ID),
		Name:                 film.Name,
		OriginalName:         film.OriginalName,
		ProdDate:             film.ProdDate,
		Slogan:               film.Slogan,
		ShortDescription:     film.ShortDescription,
		Description:          film.Description,
		AgeLimit:             film.AgeLimit,
		DurationMinutes:      int(film.DurationMinutes),
		PosterHor:            film.PosterHor,
		PosterVer:            film.PosterVer,
		Ticket:               film.Ticket,
		Trailer:              film.Trailer,
		BoxOfficeDollars:     int(film.BoxOfficeDollars),
		Budget:               int(film.Budget),
		CurrencyBudget:       film.CurrencyBudget,
		CountSeasons:         int(film.CountSeasons),
		EndYear:              film.EndYear,
		Type:                 film.Type,
		Rating:               film.Rating,
		CountRatings:         int(film.CountRatings),
		CountActors:          int(film.CountActors),
		CountPositiveReviews: int(film.CountPositiveReviews),
		CountNegativeReviews: int(film.CountNegativeReviews),
		CountNeutralReviews:  int(film.CountNeutralReviews),
		Tags:                 film.Tags,
		Genres:               film.Genres,
		ProdCompanies:        film.ProdCompanies,
		ProdCountries:        film.ProdCountries,
		Actors:               actors,
		Artists:              fillPersons(film.Artists),
		Directors:            fillPersons(film.Directors),
		Writers:              fillPersons(film.Writers),
		Producers:            fillPersons(film.Producers),
		Operators:            fillPersons(film.Operators),
		Montage:              fillPersons(film.Montage),
		Composers:            fillPersons(film.Composers),
	}
}

func NewFilmsProto(films []models.Film) []*proto.Film {
	res := make([]*proto.Film, len(films))

	for idx, value := range films {
		res[idx] = NewFilmProto(&value)
	}

	return res
}

func NewFilms(films []*proto.Film) []models.Film {
	res := make([]models.Film, len(films))

	for idx, value := range films {
		res[idx] = NewFilm(value)
	}

	return res
}

func NewGetFilmParamsProto(film *models.Film, params *constparams.GetFilmParams) *proto.GetFilmParams {
	return &proto.GetFilmParams{
		FilmID:      uint32(film.ID),
		CountImages: uint32(params.CountImages),
		CountActors: uint32(params.CountActors),
	}
}

func NewGetFilmParams(params *proto.GetFilmParams) (*models.Film, *constparams.GetFilmParams) {
	return &models.Film{
			ID: int(params.FilmID),
		},
		&constparams.GetFilmParams{
			CountImages: int(params.CountImages),
			CountActors: int(params.CountActors),
		}
}

func NewGetFilmReviewsParamsProto(params *constparams.GetFilmReviewsParams) *proto.GetFilmReviewsParams {
	return &proto.GetFilmReviewsParams{
		FilmID:       uint32(params.FilmID),
		CountReviews: uint32(params.CountReviews),
		Offset:       uint32(params.Offset),
	}
}

func NewGetFilmReviewsParams(params *proto.GetFilmReviewsParams) *constparams.GetFilmReviewsParams {
	return &constparams.GetFilmReviewsParams{
		FilmID:       int(params.FilmID),
		CountReviews: int(params.CountReviews),
		Offset:       int(params.Offset),
	}
}

// Person
func NewPersonProto(person *models.Person) *proto.Person {
	return &proto.Person{
		ID:           uint32(person.ID),
		Name:         person.Name,
		OriginalName: person.OriginalName,
		Birthday:     person.Birthday,
		Avatar:       person.Avatar,
		Death:        person.Death,
		GrowthMeters: person.GrowthMeters,
		Gender:       person.Gender,
		CountFilms:   uint32(person.CountFilms),
		Professions:  person.Professions,
		Genres:       person.Genres,
		BestFilms:    NewFilmsProto(person.BestFilms),
		Images:       person.Images,
	}
}

func NewPerson(person *proto.Person) models.Person {
	return models.Person{
		ID:           int(person.ID),
		Name:         person.Name,
		OriginalName: person.OriginalName,
		Birthday:     person.Birthday,
		Avatar:       person.Avatar,
		Death:        person.Death,
		GrowthMeters: person.GrowthMeters,
		Gender:       person.Gender,
		CountFilms:   int(person.CountFilms),
		Professions:  person.Professions,
		Genres:       person.Genres,
		BestFilms:    NewFilms(person.BestFilms),
		Images:       person.Images,
	}
}

func NewGetPersonParamsProto(person *models.Person, params *constparams.GetPersonParams) *proto.GetPersonParams {
	return &proto.GetPersonParams{
		PersonID:    uint32(person.ID),
		CountImages: uint32(params.CountImages),
		CountFilms:  uint32(params.CountFilms),
	}
}

func NewGetPersonParams(params *proto.GetPersonParams) (*models.Person, *constparams.GetPersonParams) {
	return &models.Person{
			ID: int(params.PersonID),
		},
		&constparams.GetPersonParams{
			CountImages: int(params.CountImages),
			CountFilms:  int(params.CountFilms),
		}
}

// Collection
func NewCollectionProto(collection *models.Collection) *proto.Collection {
	return &proto.Collection{
		ID:          uint32(collection.ID),
		Name:        collection.Name,
		Description: collection.Description,
		Poster:      collection.Poster,
		Time:        collection.Time,
		UpdateTime:  collection.UpdateTime,
		CreateTime:  collection.CreateTime,
		CountFilms:  uint32(collection.CountFilms),
		CountLikes:  uint32(collection.CountLikes),
		Films:       NewFilmsProto(collection.Films),
		Author:      NewUserProto(&collection.Author),
	}
}

func NewCollection(collection *proto.Collection) models.Collection {
	return models.Collection{
		ID:          int(collection.ID),
		Name:        collection.Name,
		Description: collection.Description,
		Poster:      collection.Poster,
		Time:        collection.Time,
		UpdateTime:  collection.UpdateTime,
		CreateTime:  collection.CreateTime,
		CountFilms:  int(collection.CountFilms),
		CountLikes:  int(collection.CountLikes),
		Films:       NewFilms(collection.Films),
		Author:      NewUser(collection.Author),
	}
}

func NewGetStdCollectionParamsProto(params *constparams.GetStdCollectionParams) *proto.GetStdCollectionParams {
	return &proto.GetStdCollectionParams{
		Target:     params.Target,
		Key:        params.Key,
		SortParam:  params.SortParam,
		Delimiter:  params.Delimiter,
		CountFilms: uint32(params.CountFilms),
	}
}

func NewGetStdCollectionParams(params *proto.GetStdCollectionParams) *constparams.GetStdCollectionParams {
	return &constparams.GetStdCollectionParams{
		Target:     params.Target,
		Key:        params.Key,
		SortParam:  params.SortParam,
		Delimiter:  params.Delimiter,
		CountFilms: int(params.CountFilms),
	}
}

func NewPremiersCollectionParamsProto(params *constparams.GetPremiersCollectionParams) *proto.PremiersCollectionParams {
	return &proto.PremiersCollectionParams{
		Delimiter:  uint32(params.Delimiter),
		CountFilms: uint32(params.CountFilms),
	}
}

func NewPremiersCollectionParams(params *proto.PremiersCollectionParams) *constparams.GetPremiersCollectionParams {
	return &constparams.GetPremiersCollectionParams{
		Delimiter:  int(params.Delimiter),
		CountFilms: int(params.CountFilms),
	}
}

func NewCollectionGetFilmsAuthParams(params *proto.CollectionGetFilmsAuthParams) (*models.User, *constparams.CollectionGetFilmsRequestParams) {
	return &models.User{
			ID: int(params.RequestedUser.ID),
		}, &constparams.CollectionGetFilmsRequestParams{
			CollectionID: int(params.CollectionID),
			SortParam:    params.SortParam,
		}
}

func NewCollectionGetFilmsAuthParamsProto(user *models.User, params *constparams.CollectionGetFilmsRequestParams) *proto.CollectionGetFilmsAuthParams {
	return &proto.CollectionGetFilmsAuthParams{
		CollectionID: uint32(params.CollectionID),
		SortParam:    params.SortParam,
		RequestedUser: &proto.User{
			ID: uint32(user.ID),
		},
	}
}

func NewCollectionGetFilmsNotAuthParams(params *proto.CollectionGetFilmsNotAuthParams) *constparams.CollectionGetFilmsRequestParams {
	return &constparams.CollectionGetFilmsRequestParams{
		CollectionID: int(params.CollectionID),
		SortParam:    params.SortParam,
	}
}

func NewCollectionGetFilmsNotAuthParamsProto(params *constparams.CollectionGetFilmsRequestParams) *proto.CollectionGetFilmsNotAuthParams {
	return &proto.CollectionGetFilmsNotAuthParams{
		CollectionID: uint32(params.CollectionID),
		SortParam:    params.SortParam,
	}
}

func NewSearchParamsProto(params *constparams.SearchParams) *proto.SearchParams {
	return &proto.SearchParams{
		Query: params.Query,
	}
}

func NewSearchParams(params *proto.SearchParams) *constparams.SearchParams {
	return &constparams.SearchParams{
		Query: params.Query,
	}
}

func NewSearchResponseProto(response *models.Search) *proto.SearchResponse {
	res := proto.SearchResponse{
		Films:   NewFilmsProto(response.Films),
		Series:  NewFilmsProto(response.Serials),
		Persons: make([]*proto.Person, len(response.Persons)),
	}
	for idx := range response.Persons {
		res.Persons[idx] = NewPersonProto(&response.Persons[idx])
	}
	return &res
}

func NewSearchResponse(response *proto.SearchResponse) models.Search {
	res := models.Search{
		Films:   NewFilms(response.Films),
		Serials: NewFilms(response.Series),
		Persons: make([]models.Person, len(response.Persons)),
	}
	for idx := range response.Persons {
		res.Persons[idx] = NewPerson(response.Persons[idx])
	}
	return res
}

func NewGetSimilarFilmsParamsProto(params *constparams.GetSimilarFilmsParams) *proto.GetSimilarFilmsParams {
	return &proto.GetSimilarFilmsParams{
		FilmID: uint32(params.FilmID),
	}
}

func NewGetSimilarFilmsParams(params *proto.GetSimilarFilmsParams) *constparams.GetSimilarFilmsParams {
	return &constparams.GetSimilarFilmsParams{
		FilmID: int(params.FilmID),
	}
}
