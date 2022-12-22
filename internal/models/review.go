package models

//go:generate easyjson -all -disallow_unknown_fields review.go

type Review struct {
	ID         int    `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	Type       string `json:"type,omitempty"`
	Body       string `json:"body,omitempty"`
	CountLikes int    `json:"count_likes,omitempty"`
	CreateTime string `json:"create_time,omitempty"`
	Author     User   `json:"author,omitempty"`
}
