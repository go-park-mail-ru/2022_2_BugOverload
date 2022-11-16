package errors

import (
	"context"

	stdErrors "github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"go-park-mail-ru/2022_2_BugOverload/internal/pkg"
)

func CreateLog(ctx context.Context, errFull error) {
	errCause := stdErrors.Cause(errFull)

	logger, ok := ctx.Value(pkg.LoggerKey).(*logrus.Entry)
	if !ok {
		logrus.Infof("CreateLog: errFull convert context -> *logrus.Logger on errFull [%s]", errCause)
		return
	}

	logLevel, err := GetLogLevelErr(errCause)
	if err != nil {
		logger.Error(stdErrors.Wrap(errFull, "Undefined error"))
		return
	}

	switch logLevel {
	case errLogLevel:
		logger.Error(errFull)
	case debugLogLevel:
		logger.Debug(errFull)
	default:
		logger.Info(errCause)
	}
}
