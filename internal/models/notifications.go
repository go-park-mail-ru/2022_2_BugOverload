package models

type Notification struct {
	Action  string      `json:"action,omitempty"`
	Payload interface{} `json:"payload,omitempty"`
}
