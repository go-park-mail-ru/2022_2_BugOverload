package pkg

// GetPersonParams in struct for GetPersonParams in personHandler
type GetPersonParams struct {
	CountFilms  int
	CountImages int
}

// GetCollectionTagParams in struct for GetPersonParams in tagCollectionHandler
type GetCollectionTagParams struct {
	Tag        string
	CountFilms int
	Delimiter  string
}

// GetReviewsFilmParams in struct for GetReviewsParamsKey in reviewHandler
type GetReviewsFilmParams struct {
	FilmID int
	Count  int
	Offset int
}

// GetFilmParams in struct for GetFilmParams in filmHandler
type GetFilmParams struct {
	CountImages int
}
