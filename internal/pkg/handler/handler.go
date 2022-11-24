package handler

import (
	"net/http"

	"github.com/gorilla/mux"

	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/middleware"
)

// Handler is an interface for universal handlers.
type Handler interface {
	Action(http.ResponseWriter, *http.Request)
	Configure(*mux.Router, *middleware.HTTPMiddleware)
}
