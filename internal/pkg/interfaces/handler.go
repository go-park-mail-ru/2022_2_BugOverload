package interfaces

import "net/http"

// Handler is an interface for universal handlers.
type Handler interface {
	Action(http.ResponseWriter, *http.Request)
}
