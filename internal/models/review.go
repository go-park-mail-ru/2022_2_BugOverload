package models

type Review struct {
	ID         int    `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	Type       string `json:"type,omitempty"`
	Time       string `json:"time,omitempty"`
	Body       string `json:"body,omitempty"`
	CountLikes int    `json:"count_likes,omitempty"`
	Author     User   `json:"author,omitempty"`
}
