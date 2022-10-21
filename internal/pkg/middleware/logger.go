package middleware

import (
	"context"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"

	pkgInner "go-park-mail-ru/2022_2_BugOverload/internal/pkg"
)

type LoggerMiddleware struct {
	log *logrus.Logger
}

func NewLoggerMiddleware(log *logrus.Logger) LoggerMiddleware {
	return LoggerMiddleware{
		log: log,
	}
}

func (umd *LoggerMiddleware) SetDefaultLoggerMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), pkgInner.LoggerKey, umd.log)

		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (umd *LoggerMiddleware) UpdateDefaultLoggerMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger, ok := r.Context().Value(pkgInner.LoggerKey).(*logrus.Logger)
		if !ok {
			logrus.Fatal("GetLoggerContext: err convert context -> *logrus.Logger")
		}

		start := time.Now()
		upgradeLogger := logger.WithFields(logrus.Fields{
			"urls":        r.URL.Path,
			"method":      r.Method,
			"remote_addr": r.RemoteAddr,
			"req_id":      uuid.NewV4(),
		})

		ctx := context.WithValue(r.Context(), pkgInner.LoggerKey, upgradeLogger)

		h.ServeHTTP(w, r.WithContext(ctx))

		executeTime := time.Since(start).Milliseconds()
		upgradeLogger.Infof("work time [ms]: %v", executeTime)
	})
}
