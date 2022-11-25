package pkg

import (
	"context"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/metadata"
)

func GetDefInfoMicroService(ctx context.Context) context.Context {
	requestID, ok := ctx.Value(constparams.RequestIDKey).(string)
	if !ok {
		logrus.Fatal("GetDefInfoMicroService: err convert context -> string")
	}

	ctx = context.Background()
	md := metadata.Pairs(
		constparams.RequestID, requestID,
	)

	return metadata.NewOutgoingContext(ctx, md)
}
