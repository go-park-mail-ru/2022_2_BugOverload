package repository

import (
	"database/sql"
	"time"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
)

type AuthorSQL struct {
	ID           int
	Nickname     string
	CountReviews sql.NullInt32
	Avatar       sql.NullString
}

type ReviewSQL struct {
	Name       string
	Type       string
	Body       string
	CountLikes sql.NullInt32
	CreateTime time.Time
	Author     AuthorSQL
}

func (r *ReviewSQL) Convert() models.Review {
	return models.Review{
		Name:       r.Name,
		Type:       r.Type,
		Body:       r.Body,
		CreateTime: r.CreateTime.Format("2006-01-02 15:04:05"),
		CountLikes: int(r.CountLikes.Int32),
		Author: models.User{
			ID:       r.Author.ID,
			Nickname: r.Author.Nickname,
			Profile: models.Profile{
				Avatar:       r.Author.Avatar.String,
				CountReviews: int(r.Author.CountReviews.Int32),
			},
		},
	}
}
