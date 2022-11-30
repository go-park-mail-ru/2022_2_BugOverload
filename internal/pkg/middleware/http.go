package middleware

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/rs/cors"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"github.com/urfave/negroni"

	"go-park-mail-ru/2022_2_BugOverload/internal/auth/delivery/grpc/client"
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/monitoring"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/security"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/wrapper"
)

type HTTPMiddleware struct {
	log     *logrus.Logger
	auth    client.AuthService
	cors    cors.Cors
	metrics monitoring.Monitoring
}

func NewHTTPMiddleware(log *logrus.Logger, auth client.AuthService, config *pkg.Cors, metrics monitoring.Monitoring) *HTTPMiddleware {
	cors := cors.New(cors.Options{
		AllowedMethods:   config.Methods,
		AllowedOrigins:   config.Origins,
		AllowCredentials: config.Credentials,
		AllowedHeaders:   config.Headers,
		Debug:            config.Debug,
	})

	return &HTTPMiddleware{
		auth:    auth,
		log:     log,
		cors:    *cors,
		metrics: metrics,
	}
}

func (m *HTTPMiddleware) SetCORSMiddleware(h http.Handler) http.Handler {
	return m.cors.Handler(h)
}

func (m *HTTPMiddleware) SetDefaultLoggerMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), constparams.LoggerKey, m.log)

		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (m *HTTPMiddleware) UpdateDefaultLoggerMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger, ok := r.Context().Value(constparams.LoggerKey).(*logrus.Logger)
		if !ok {
			logrus.Fatal("GetLoggerContext: err convert context -> *logrus.Logger")
		}

		requestID := uuid.NewV4().String()

		start := time.Now()

		upgradeLogger := logger.WithFields(logrus.Fields{
			"url":                 r.URL.Path,
			"method":              r.Method,
			"remote_addr":         r.RemoteAddr,
			constparams.RequestID: requestID,
		})

		ctx := context.WithValue(r.Context(), constparams.LoggerKey, upgradeLogger)

		ctx = context.WithValue(ctx, constparams.RequestIDKey, requestID)

		h.ServeHTTP(w, r.WithContext(ctx))

		executeTime := time.Since(start).Milliseconds()
		upgradeLogger.Infof("work time [ms]: %v", executeTime)
	})
}

func (m *HTTPMiddleware) SetSizeRequest(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		strLength := r.Header.Get("Content-Length")
		if strLength == "" {
			h.ServeHTTP(w, r)
			return
		}

		length, err := strconv.Atoi(strLength)
		if err != nil {
			wrapper.DefaultHandlerHTTPError(r.Context(), w, errors.ErrConvertLength)
			return
		}

		if length > constparams.BufSizeRequest {
			wrapper.DefaultHandlerHTTPError(r.Context(), w, errors.ErrBigRequest)
			return
		}

		h.ServeHTTP(w, r)
	})
}

func (m *HTTPMiddleware) NeedAuthMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(constparams.SessionCookieName)
		if err != nil {
			wrapper.DefaultHandlerHTTPError(r.Context(), w, errors.ErrNoCookie)
			return
		}

		currentSession := &models.Session{ID: cookie.Value}

		user, err := m.auth.GetUserBySession(r.Context(), currentSession)
		if err != nil {
			wrapper.DefaultHandlerHTTPError(r.Context(), w, err)
			return
		}

		ctx := context.WithValue(r.Context(), constparams.CurrentUserKey, user)
		r = r.WithContext(ctx)

		h.ServeHTTP(w, r)
	})
}

func (m *HTTPMiddleware) TryAuthMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(constparams.SessionCookieName)
		if err != nil {
			h.ServeHTTP(w, r)
			return
		}

		currentSession := &models.Session{ID: cookie.Value}

		user, err := m.auth.GetUserBySession(r.Context(), currentSession)
		if err != nil {
			h.ServeHTTP(w, r)
			return
		}

		ctx := context.WithValue(r.Context(), constparams.CurrentUserKey, user)
		r = r.WithContext(ctx)

		h.ServeHTTP(w, r)
	})
}

func (m *HTTPMiddleware) SetCsrfMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-Csrf-Token")

		cookie, err := r.Cookie(constparams.SessionCookieName)
		if err != nil {
			wrapper.DefaultHandlerHTTPError(r.Context(), w, errors.ErrNoCookie)
			return
		}

		currentSession := &models.Session{ID: cookie.Value}

		user, err := m.auth.GetUserBySession(r.Context(), currentSession)
		if err != nil {
			wrapper.DefaultHandlerHTTPError(r.Context(), w, err)
			return
		}

		currentSession.User = &user

		_, err = security.CheckCsrfToken(currentSession, token)
		if err != nil {
			wrapper.DefaultHandlerHTTPError(r.Context(), w, errors.ErrCsrfTokenInvalid)
			return
		}

		h.ServeHTTP(w, r)
	})
}

func (m *HTTPMiddleware) SetDefaultMetrics(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wrappedWriter := negroni.NewResponseWriter(w)

		start := time.Now()
		h.ServeHTTP(wrappedWriter, r)

		statusCode := wrappedWriter.Status()

		m.metrics.GetExecution().
			WithLabelValues(strconv.Itoa(statusCode), r.URL.String(), r.Method).
			Observe(time.Since(start).Seconds())

		m.metrics.GetRequestCounter().Inc()

		if statusCode < constparams.MaxSuccessResponse {
			m.metrics.GetSuccessHits().WithLabelValues(
				strconv.Itoa(statusCode),
				r.URL.String(),
				r.Method,
			).Inc()
		} else {
			m.metrics.GetErrorsHits().WithLabelValues(
				strconv.Itoa(statusCode),
				r.URL.String(),
				r.Method,
			).Inc()
		}
	})
}
