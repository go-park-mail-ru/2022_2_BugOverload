package wrapper

import (
	stdErrors "github.com/pkg/errors"
	"google.golang.org/grpc/status"
)

// GRPCErrorConvert status -> error
func GRPCErrorConvert(err error) error {
	s, _ := status.FromError(stdErrors.Cause(err))

	return stdErrors.New(s.Message())
}
