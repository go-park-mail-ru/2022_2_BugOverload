package errors

import (
	"context"

	stdErrors "github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"go-park-mail-ru/2022_2_BugOverload/internal/pkg"
)

func CreateLog(ctx context.Context, errFull error) {
	errShort := stdErrors.Cause(errFull)

	logger, ok := ctx.Value(pkg.LoggerKey).(*logrus.Entry)
	if !ok {
		logrus.Infof("CreateLog: errFull convert context -> *logrus.Logger on errFull [%s]", errShort)
	}

	logger.Error(errFull)

	logger.Info(errShort)
}
