package middleware

import (
	"context"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type GRPCMiddleware struct {
	log *logrus.Logger
}

func NewGRPCMiddleware(log *logrus.Logger) *GRPCMiddleware {
	return &GRPCMiddleware{
		log: log,
	}
}

func (m *GRPCMiddleware) LoggerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, _ := metadata.FromIncomingContext(ctx)

	requestID := md.Get(constparams.RequestID)

	start := time.Now()

	upgradeLogger := m.log.WithFields(logrus.Fields{
		"req_id": requestID[0],
	})

	ctx = context.WithValue(ctx, constparams.LoggerKey, upgradeLogger)

	reply, err := handler(ctx, req)

	executeTime := time.Since(start).Milliseconds()
	upgradeLogger.Infof("work time [ms]: %v", executeTime)

	return reply, err
}
