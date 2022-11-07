package middleware

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/rs/cors"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/httpwrapper"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/security"
	sessionService "go-park-mail-ru/2022_2_BugOverload/internal/session/service"
)

type Middleware struct {
	log     *logrus.Logger
	session sessionService.SessionService
	cors    cors.Cors
}

func NewMiddleware(log *logrus.Logger, session sessionService.SessionService, config *pkg.Cors) *Middleware {
	cors := cors.New(cors.Options{
		AllowedMethods:   config.Methods,
		AllowedOrigins:   config.Origins,
		AllowCredentials: config.Credentials,
		AllowedHeaders:   config.Headers,
		Debug:            config.Debug,
	})

	return &Middleware{
		session: session,
		log:     log,
		cors:    *cors,
	}
}

func (m *Middleware) SetCORSMiddleware(h http.Handler) http.Handler {
	return m.cors.Handler(h)
}

func (m *Middleware) SetDefaultLoggerMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), pkg.LoggerKey, m.log)

		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (m *Middleware) UpdateDefaultLoggerMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger, ok := r.Context().Value(pkg.LoggerKey).(*logrus.Logger)
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

		ctx := context.WithValue(r.Context(), pkg.LoggerKey, upgradeLogger)

		h.ServeHTTP(w, r.WithContext(ctx))

		executeTime := time.Since(start).Milliseconds()
		upgradeLogger.Infof("work time [ms]: %v", executeTime)
	})
}

func (m *Middleware) SetSizeRequest(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		strLength := r.Header.Get("Content-Length")
		if strLength == "" {
			h.ServeHTTP(w, r)
			return
		}

		length, err := strconv.Atoi(strLength)
		if err != nil {
			httpwrapper.DefaultHandlerError(w, errors.NewErrValidation(errors.ErrConvertLength))
			return
		}

		if length > pkg.BufSizeRequest {
			httpwrapper.DefaultHandlerError(w, errors.NewErrValidation(errors.ErrBigRequest))
			return
		}

		h.ServeHTTP(w, r)
	})
}

func (m *Middleware) CheckAuthMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_id")
		if err != nil {
			httpwrapper.DefaultHandlerError(w, errors.NewErrValidation(errors.ErrNoCookie))
			return
		}

		currentSession := models.Session{ID: cookie.Value}

		user, err := m.session.GetUserBySession(r.Context(), currentSession)
		if err != nil {
			httpwrapper.DefaultHandlerError(w, errors.NewErrAuth(err))
			return
		}

		ctx := context.WithValue(r.Context(), pkg.CurrentUserKey, user)
		r = r.WithContext(ctx)

		h.ServeHTTP(w, r)
	})
}

func (m *Middleware) SetCsrfMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-Csrf-Token")

		if token == "" {
			httpwrapper.DefaultHandlerError(w, errors.NewErrAuth(errors.ErrCsrfTokenNotFound))
			return
		}

		cookie, err := r.Cookie("session_id")
		if err != nil {
			httpwrapper.DefaultHandlerError(w, errors.NewErrValidation(errors.ErrNoCookie))
			return
		}

		currentSession := models.Session{ID: cookie.Value}

		user, err := m.session.GetUserBySession(r.Context(), currentSession)
		if err != nil {
			httpwrapper.DefaultHandlerError(w, errors.NewErrAuth(err))
			return
		}

		currentSession.User = &user

		correctToken, err := security.CheckCsrfToken(&currentSession, token)
		if err != nil {
			httpwrapper.DefaultHandlerError(w, errors.NewErrAuth(err))
			return
		}
		if !correctToken {
			httpwrapper.DefaultHandlerError(w, errors.NewErrAuth(errors.ErrCsrfTokenInvalid))
			return
		}

		h.ServeHTTP(w, r)
	})
}
