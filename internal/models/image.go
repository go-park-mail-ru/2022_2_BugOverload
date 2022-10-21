package models

type Image struct {
	Key    string `json:"key"`
	Object string `json:"object"`
	Bytes  []byte `json:"-"`
}
