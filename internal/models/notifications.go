package models

//go:generate easyjson -all -disallow_unknown_fields notifications.go

type Notification struct {
	Action  string      `json:"action,omitempty"`
	Payload interface{} `json:"payload,omitempty"`
}
