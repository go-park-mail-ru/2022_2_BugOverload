package constparams

// GetPersonParams in struct for personHandler
type GetPersonParams struct {
	CountFilms  int
	CountImages int
}

// GetStdCollectionParams in struct for getStdCollectionHandler
type GetStdCollectionParams struct {
	Target     string
	Key        string
	SortParam  string
	CountFilms int
	Delimiter  string
}

// PremiersCollectionParams in struct for getStdCollectionHandler
type PremiersCollectionParams struct {
	CountFilms int
	Delimiter  int
}

// CollectionFilmsUpdateParams in struct for premieresCollectionHandler
type CollectionFilmsUpdateParams struct {
	CollectionID int
	FilmID       int
}

// GetUserCollectionsParams in struct for addFilmHandler and dropFilmHandler
type GetUserCollectionsParams struct {
	SortParam        string
	CountCollections int
	Delimiter        string
}

// GetReviewsFilmParams in struct for reviewHandler
type GetReviewsFilmParams struct {
	FilmID       int
	CountReviews int
	Offset       int
}

// GetFilmParams in struct for filmHandler
type GetFilmParams struct {
	CountImages int
}

// ChangeUserSettings in struct for changeUserSettings
type ChangeUserSettings struct {
	CurPassword string
	NewPassword string
	Nickname    string
}

// FilmRateParams in struct for filmRateHandler
type FilmRateParams struct {
	FilmID int
	Score  int
}

// FilmRateDropParams in struct for filmRateDropHandler
type FilmRateDropParams struct {
	FilmID int
}

// NewFilmReviewParams in struct for newFilmReviewHandler
type NewFilmReviewParams struct {
	FilmID int
}

// GetUserActivityOnFilmParams in struct for getUserActivityOnFilmHandler
type GetUserActivityOnFilmParams struct {
	FilmID int
}
