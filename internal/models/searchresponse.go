package models

type SearchResponse struct {
	Films   []Film   `json:"films,omitempty"`
	Series  []Film   `json:"series,omitempty"`
	Persons []Person `json:"persons,omitempty"`
}
