package pkg

import (
	"context"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/metadata"
)

func GetDefInfoMicroService(ctx context.Context) context.Context {
	requestID, ok := ctx.Value(RequestIDKey).(string)
	if !ok {
		logrus.Fatal("GetDefInfoMicroService: err convert context -> string")
	}

	ctx = context.Background()
	md := metadata.Pairs(
		RequestID, requestID,
	)

	return metadata.NewOutgoingContext(ctx, md)
}
