package interfaces

import "net/http"

type Handler interface {
	Action(http.ResponseWriter, *http.Request)
}
