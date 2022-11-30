package middleware

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/monitoring"
)

type GRPCMiddleware struct {
	log     *logrus.Logger
	metrics monitoring.Monitoring
}

func NewGRPCMiddleware(log *logrus.Logger, metrics monitoring.Monitoring) *GRPCMiddleware {
	return &GRPCMiddleware{
		log:     log,
		metrics: metrics,
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

func (m *GRPCMiddleware) MetricsInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()

	reply, err := handler(ctx, req)

	errStatus, _ := status.FromError(err)

	code := errStatus.Code()

	if code != codes.OK {
		m.metrics.GetErrorsHits().WithLabelValues(code.String(), info.FullMethod, "").Inc()
	} else {
		m.metrics.GetSuccessHits().WithLabelValues(codes.OK.String(), info.FullMethod, "").Inc()
	}

	m.metrics.GetExecution().
		WithLabelValues(code.String(), info.FullMethod, info.FullMethod).
		Observe(time.Since(start).Seconds())

	m.metrics.GetRequestCounter().Inc()

	return reply, err
}
