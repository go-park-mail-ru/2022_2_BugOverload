package models

// FilmCollection is a structure for response to getPopularFilms/etc requests
type FilmCollection struct {
	Title string `json:"title,omitempty"`
	Films []Film `json:"films,omitempty"`
}

// CreateFilmCollection returns collection
func CreateFilmCollection(title string, films []Film) FilmCollection {
	return FilmCollection{
		Title: title,
		Films: films,
	}
}
