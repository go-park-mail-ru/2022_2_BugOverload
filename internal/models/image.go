package models

//go:generate easyjson -all -disallow_unknown_fields image.go

type Image struct {
	Key    string `json:"key"`
	Object string `json:"object"`
	Bytes  []byte `json:"-"`
}
