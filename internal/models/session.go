package models

type Session struct {
	ID   string `json:"session_id,omitempty"`
	User *User  `json:"user,omitempty"`
}
