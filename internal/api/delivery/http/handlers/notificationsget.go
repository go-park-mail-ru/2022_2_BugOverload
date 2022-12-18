package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	stdErrors "github.com/pkg/errors"

	"go-park-mail-ru/2022_2_BugOverload/internal/notification/service"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/handler"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/middleware"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/wrapper"
)

// getUserNotificationsHandler is the structure that handles the request.
type getUserNotificationsHandler struct {
	notificationsService service.NotificationsService
}

// NewGetNotificationsHandler is constructor for getUserNotificationsHandler in this pkg.
func NewGetNotificationsHandler(service service.NotificationsService) handler.Handler {
	return &getUserNotificationsHandler{
		service,
	}
}

func (h *getUserNotificationsHandler) Configure(r *mux.Router, mw *middleware.HTTPMiddleware) {
	r.HandleFunc("/api/v1/notifications", mw.NeedAuthMiddleware(mw.SetCsrfMiddleware(h.Action))).
		Methods(http.MethodGet)
}

var upgrade = websocket.Upgrader{
	ReadBufferSize:  constparams.WSBufSizeRead,
	WriteBufferSize: constparams.WSBufSizeWrite,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func sendNewMsgNotifications(client *websocket.Conn, messages []interface{}) {
	ticker := time.NewTicker(constparams.PushTimeout)
	for _, val := range messages {
		message, err := json.Marshal(val)
		if err != nil {
			break
		}

		w, err := client.NextWriter(websocket.TextMessage)
		if err != nil {
			ticker.Stop()
			break
		}

		_, err = w.Write(message)
		if err != nil {
			ticker.Stop()
			break
		}

		w.Close()

		<-ticker.C
	}
}

// Action is a method for initial validation of the request and data and
// delivery of the data to the service at the business logic level.
// @Summary Getting notifications
// @Description Return user notifications
// @tags notifications, in_process
// @Success 200 "successfully getting"
// @Failure 400 "return error"
// @Failure 404 "not found"
// @Failure 405 "method not allowed"
// @Failure 500 "something unusual has happened"
// @Router /api/v1/notifications [GET]
func (h *getUserNotificationsHandler) Action(w http.ResponseWriter, r *http.Request) {
	connection, err := upgrade.Upgrade(w, r, nil)
	if err != nil {
		wrapper.DefaultHandlerHTTPError(
			r.Context(),
			w,
			stdErrors.WithMessagef(errors.ErrUpdateWebSocketProtocol, "Special Error [%s]", err))
	}

	go sendNewMsgNotifications(connection, h.notificationsService.GetMessages())
}
