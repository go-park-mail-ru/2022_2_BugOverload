package pkg

// GetPersonParams in struct for personHandler
type GetPersonParams struct {
	CountFilms  int
	CountImages int
}

// GetCollectionTagParams in struct for tagCollectionHandler
type GetCollectionTagParams struct {
	Tag        string
	CountFilms int
	Delimiter  string
}

// GetReviewsFilmParams in struct for reviewHandler
type GetReviewsFilmParams struct {
	FilmID int
	Count  int
	Offset int
}

// GetFilmParams in struct for filmHandler
type GetFilmParams struct {
	CountImages int
}

// ChangeUserSettings in struct for changeUserSettings
type ChangeUserSettings struct {
	CurPassword string
	NewPassword string
}

// FilmRateParams in struct for filmHandler
type FilmRateParams struct {
	FilmID int
}
