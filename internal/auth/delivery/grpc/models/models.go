package models

import (
	proto "go-park-mail-ru/2022_2_BugOverload/internal/auth/delivery/grpc/protobuf"
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
)

// User
func NewUserProto(user *models.User) *proto.User {
	return &proto.User{
		ID:       uint32(user.ID),
		Nickname: user.Nickname,
		Email:    user.Email,
		Password: []byte(user.Password),
		Avatar:   user.Avatar,
	}
}

func NewUser(user *proto.User) *models.User {
	return &models.User{
		ID:       int(user.ID),
		Nickname: user.Nickname,
		Email:    user.Email,
		Password: string(user.Password),
		Avatar:   user.Avatar,
	}
}

// Session
func NewSessionProto(session *models.Session) *proto.Session {
	user := &proto.User{}

	if session.User != nil {
		user = NewUserProto(session.User)
	}

	return &proto.Session{
		ID:   session.ID,
		User: user,
	}
}

func NewSession(session *proto.Session) *models.Session {
	return &models.Session{
		ID:   session.ID,
		User: NewUser(session.User),
	}
}

// GetAccessParams
func NewGetAccessParamsProto(user *models.User, userPassword string) *proto.GetAccessParams {
	return &proto.GetAccessParams{
		UserPassword: userPassword,
		User:         NewUserProto(user),
	}
}

func NewGetAccessParams(params *proto.GetAccessParams) (*models.User, string) {
	return NewUser(params.User), params.UserPassword
}

// UpdatePasswordParams
func NewUpdatePasswordParamsProto(user *models.User, userPassword string, newUserPassword string) *proto.UpdatePasswordParams {
	return &proto.UpdatePasswordParams{
		UserPassword:    userPassword,
		NewUserPassword: newUserPassword,
		User:            NewUserProto(user),
	}
}

func NewUpdatePasswordParams(params *proto.UpdatePasswordParams) (*models.User, string, string) {
	return NewUser(params.User), params.UserPassword, params.NewUserPassword
}
