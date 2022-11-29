package models

type Search struct {
	Films   []Film   `json:"films,omitempty"`
	Serials []Film   `json:"serials,omitempty"`
	Persons []Person `json:"persons,omitempty"`
}
