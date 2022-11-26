package repository

import (
	"database/sql"
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
)

type UserSQL struct {
	ID       int
	nickname sql.NullString
	email    sql.NullString
	password []byte
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
		Password: string(u.password),
		Avatar:   u.avatar.String,
	}
}
