package repository

import "go-park-mail-ru/2022_2_BugOverload/internal/models"

type UserSQL struct {
	userID   int
	nickname string
	email    string
	password string
}

func NewUserSQL() UserSQL {
	return UserSQL{}
}

func NewUser(user UserSQL, avatar string) models.User {
	return models.User{
		ID:       user.userID,
		Nickname: user.nickname,
		Email:    user.email,
		Password: user.password,
		Profile: models.Profile{
			Avatar: avatar,
		},
	}
}
