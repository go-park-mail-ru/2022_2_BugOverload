package models

//go:generate easyjson -all -disallow_unknown_fields session.go

type Session struct {
	ID   string `json:"session_id,omitempty"`
	User *User  `json:"user,omitempty"`
}
