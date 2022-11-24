package wrapper

import (
	"context"

	stdErrors "github.com/pkg/errors"
	"google.golang.org/grpc/status"

	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
)

// DefaultHandlerGRPCError is error handler that detects the type of error and gives an error response.
func DefaultHandlerGRPCError(ctx context.Context, err error) error {
	errors.CreateLog(ctx, err)

	errCause := stdErrors.Cause(err).Error()

	return status.Error(errors.GetCodeGRPC(errCause), errCause)
}
