package models

type Review struct {
	ID         uint   `json:"film_id,omitempty"`
	Name       string `json:"film_name,omitempty"`
	Score      int    `json:"score,omitempty"`
	Time       string `json:"time,omitempty"`
	CountLikes int    `json:"count_likes,omitempty"`
	Author     User   `json:"author,omitempty"`
}
