package structs

// Film collection is a structure for response to getPopularFilms/etc requests
type FilmCollection struct {
	Title string `json:"title,omitempty"`
	Films []Film `json:"films,omitempty"`
}

// Create new film collection
func CreateFilmCollection(title string, films []Film) FilmCollection {
	return FilmCollection{
		Title: title,
		Films: films,
	}
}
