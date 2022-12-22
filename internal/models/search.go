package models

//go:generate easyjson -all -disallow_unknown_fields search.go

type Search struct {
	Films   []Film   `json:"films,omitempty"`
	Serials []Film   `json:"serials,omitempty"`
	Persons []Person `json:"persons,omitempty"`
}
