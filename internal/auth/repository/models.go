package repository

import (
	"database/sql"
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
)

type UserSQL struct {
	ID       int
	nickname sql.NullString
	email    sql.NullString
	password sql.NullString
	avatar   sql.NullString
}

func NewUserSQL() UserSQL {
	return UserSQL{}
}

func (u *UserSQL) Convert() models.User {
	return models.User{
		ID:       u.ID,
		Nickname: u.nickname.String,
		Email:    u.email.String,
		Password: u.password.String,
		Profile: models.Profile{
			Avatar: u.avatar.String,
		},
	}
}
