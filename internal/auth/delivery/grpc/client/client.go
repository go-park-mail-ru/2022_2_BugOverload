package client

import (
	"context"

	stdErrors "github.com/pkg/errors"
	"google.golang.org/grpc"

	"go-park-mail-ru/2022_2_BugOverload/internal/auth/delivery/grpc/models"
	proto "go-park-mail-ru/2022_2_BugOverload/internal/auth/delivery/grpc/protobuf"
	modelsGlobal "go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/wrapper"
)

//go:generate mockgen -source client.go -destination mocks/mockauthclient.go -package mockAuthClient

type AuthService interface {
	Auth(ctx context.Context, user *modelsGlobal.User) (modelsGlobal.User, error)
	Login(ctx context.Context, user *modelsGlobal.User) (modelsGlobal.User, error)
	Signup(ctx context.Context, user *modelsGlobal.User) (modelsGlobal.User, error)
	GetAccess(ctx context.Context, user *modelsGlobal.User, userPassword string) error
	UpdatePassword(ctx context.Context, user *modelsGlobal.User, userPassword, userNewPassword string) error

	GetUserBySession(ctx context.Context, session *modelsGlobal.Session) (modelsGlobal.User, error)
	CreateSession(ctx context.Context, user *modelsGlobal.User) (modelsGlobal.Session, error)
	DeleteSession(ctx context.Context, session *modelsGlobal.Session) (modelsGlobal.Session, error)
}

type AuthServiceGRPCClient struct {
	authClient proto.AuthServiceClient
}

func NewAuthServiceGRPSClient(con *grpc.ClientConn) AuthService {
	return &AuthServiceGRPCClient{
		authClient: proto.NewAuthServiceClient(con),
	}
}

func (a AuthServiceGRPCClient) Auth(ctx context.Context, user *modelsGlobal.User) (modelsGlobal.User, error) {
	authProtoResponse, err := a.authClient.Auth(pkg.GetDefInfoMicroService(ctx), models.NewUserProto(user))
	if err != nil {
		return modelsGlobal.User{}, wrapper.GRPCErrorConvert(stdErrors.Wrap(err, "Auth"))
	}

	return *models.NewUser(authProtoResponse), nil
}

func (a AuthServiceGRPCClient) Login(ctx context.Context, user *modelsGlobal.User) (modelsGlobal.User, error) {
	loginProtoResponse, err := a.authClient.Login(pkg.GetDefInfoMicroService(ctx), models.NewUserProto(user))
	if err != nil {
		return modelsGlobal.User{}, wrapper.GRPCErrorConvert(stdErrors.Wrap(err, "Login"))
	}

	return *models.NewUser(loginProtoResponse), nil
}

func (a AuthServiceGRPCClient) Signup(ctx context.Context, user *modelsGlobal.User) (modelsGlobal.User, error) {
	sighupProtoResponse, err := a.authClient.Signup(pkg.GetDefInfoMicroService(ctx), models.NewUserProto(user))
	if err != nil {
		return modelsGlobal.User{}, wrapper.GRPCErrorConvert(stdErrors.Wrap(err, "Signup"))
	}

	return *models.NewUser(sighupProtoResponse), nil
}

func (a AuthServiceGRPCClient) GetAccess(ctx context.Context, user *modelsGlobal.User, userPassword string) error {
	_, err := a.authClient.GetAccess(pkg.GetDefInfoMicroService(ctx), models.NewGetAccessParamsProto(user, userPassword))
	if err != nil {
		return wrapper.GRPCErrorConvert(stdErrors.Wrap(err, "GetAccess"))
	}

	return nil
}

func (a AuthServiceGRPCClient) UpdatePassword(ctx context.Context, user *modelsGlobal.User, userPassword, newUserPassword string) error {
	_, err := a.authClient.UpdatePassword(pkg.GetDefInfoMicroService(ctx), models.NewUpdatePasswordParamsProto(user, userPassword, newUserPassword))
	if err != nil {
		return wrapper.GRPCErrorConvert(stdErrors.Wrap(err, "UpdatePassword"))
	}

	return nil
}

func (a AuthServiceGRPCClient) GetUserBySession(ctx context.Context, session *modelsGlobal.Session) (modelsGlobal.User, error) {
	userProtoResponse, err := a.authClient.GetUserBySession(pkg.GetDefInfoMicroService(ctx), models.NewSessionProto(session))
	if err != nil {
		return modelsGlobal.User{}, wrapper.GRPCErrorConvert(stdErrors.Wrap(err, "GetUserBySession"))
	}

	return *models.NewUser(userProtoResponse), nil
}

func (a AuthServiceGRPCClient) CreateSession(ctx context.Context, user *modelsGlobal.User) (modelsGlobal.Session, error) {
	sessionProtoResponse, err := a.authClient.CreateSession(pkg.GetDefInfoMicroService(ctx), models.NewUserProto(user))
	if err != nil {
		return modelsGlobal.Session{}, wrapper.GRPCErrorConvert(stdErrors.Wrap(err, "GetUserBySession"))
	}

	return *models.NewSession(sessionProtoResponse), nil
}

func (a AuthServiceGRPCClient) DeleteSession(ctx context.Context, session *modelsGlobal.Session) (modelsGlobal.Session, error) {
	sessionProtoResponse, err := a.authClient.DeleteSession(pkg.GetDefInfoMicroService(ctx), models.NewSessionProto(session))
	if err != nil {
		return modelsGlobal.Session{}, wrapper.GRPCErrorConvert(stdErrors.Wrap(err, "DeleteSession"))
	}

	return *models.NewSession(sessionProtoResponse), nil
}
