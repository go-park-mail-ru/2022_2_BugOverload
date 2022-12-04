package server

import (
	"context"
	"net"

	"google.golang.org/grpc"

	"go-park-mail-ru/2022_2_BugOverload/internal/auth/delivery/grpc/models"
	proto "go-park-mail-ru/2022_2_BugOverload/internal/auth/delivery/grpc/protobuf"
	"go-park-mail-ru/2022_2_BugOverload/internal/auth/service"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/wrapper"
)

type AuthServiceGRPCServer struct {
	grpcServer *grpc.Server

	authManager    service.AuthService
	sessionManager service.SessionService
}

func NewAuthServiceGRPCServer(grpcServer *grpc.Server, am service.AuthService, sm service.SessionService) *AuthServiceGRPCServer {
	return &AuthServiceGRPCServer{
		grpcServer:     grpcServer,
		authManager:    am,
		sessionManager: sm,
	}
}

func (s *AuthServiceGRPCServer) StartGRPCServer(listenURL string) error {
	lis, err := net.Listen("tcp", listenURL)
	if err != nil {
		return err
	}

	proto.RegisterAuthServiceServer(s.grpcServer, s)

	return s.grpcServer.Serve(lis)
}

func (s *AuthServiceGRPCServer) Auth(ctx context.Context, user *proto.User) (*proto.User, error) {
	authProtoResponse, err := s.authManager.Auth(ctx, models.NewUser(user))
	if err != nil {
		return &proto.User{}, wrapper.DefaultHandlerGRPCError(ctx, err)
	}

	return models.NewUserProto(&authProtoResponse), nil
}

func (s *AuthServiceGRPCServer) Login(ctx context.Context, user *proto.User) (*proto.User, error) {
	loginProtoResponse, err := s.authManager.Login(ctx, models.NewUser(user))
	if err != nil {
		return &proto.User{}, wrapper.DefaultHandlerGRPCError(ctx, err)
	}

	return models.NewUserProto(&loginProtoResponse), nil
}

func (s *AuthServiceGRPCServer) Signup(ctx context.Context, user *proto.User) (*proto.User, error) {
	signupProtoResponse, err := s.authManager.Signup(ctx, models.NewUser(user))
	if err != nil {
		return &proto.User{}, wrapper.DefaultHandlerGRPCError(ctx, err)
	}

	return models.NewUserProto(&signupProtoResponse), nil
}

func (s *AuthServiceGRPCServer) GetAccess(ctx context.Context, params *proto.GetAccessParams) (*proto.Nothing, error) {
	user, password := models.NewGetAccessParams(params)

	err := s.authManager.GetAccess(ctx, user, password)
	if err != nil {
		return &proto.Nothing{}, wrapper.DefaultHandlerGRPCError(ctx, err)
	}

	return &proto.Nothing{}, nil
}

func (s *AuthServiceGRPCServer) UpdatePassword(ctx context.Context, params *proto.UpdatePasswordParams) (*proto.Nothing, error) {
	user, password, newPassword := models.NewUpdatePasswordParams(params)

	err := s.authManager.UpdatePassword(ctx, user, password, newPassword)
	if err != nil {
		return &proto.Nothing{}, wrapper.DefaultHandlerGRPCError(ctx, err)
	}

	return &proto.Nothing{}, nil
}

func (s *AuthServiceGRPCServer) GetUserBySession(ctx context.Context, session *proto.Session) (*proto.User, error) {
	user, err := s.sessionManager.GetUserBySession(ctx, models.NewSession(session))
	if err != nil {
		return &proto.User{}, wrapper.DefaultHandlerGRPCError(ctx, err)
	}

	return models.NewUserProto(&user), nil
}

func (s *AuthServiceGRPCServer) CreateSession(ctx context.Context, user *proto.User) (*proto.Session, error) {
	session, err := s.sessionManager.CreateSession(ctx, models.NewUser(user))
	if err != nil {
		return &proto.Session{}, wrapper.DefaultHandlerGRPCError(ctx, err)
	}

	return models.NewSessionProto(&session), nil
}

func (s *AuthServiceGRPCServer) DeleteSession(ctx context.Context, session *proto.Session) (*proto.Session, error) {
	sessionDel, err := s.sessionManager.DeleteSession(ctx, models.NewSession(session))
	if err != nil {
		return &proto.Session{}, wrapper.DefaultHandlerGRPCError(ctx, err)
	}

	return models.NewSessionProto(&sessionDel), nil
}
