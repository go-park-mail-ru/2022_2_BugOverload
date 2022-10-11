package models

type FilmCollection struct {
	Title string `json:"title,omitempty"`
	Films []Film `json:"films,omitempty"`
}

func NewFilmCollection(title string, films []Film) *FilmCollection {
	return &FilmCollection{
		Title: title,
		Films: films,
	}
}
