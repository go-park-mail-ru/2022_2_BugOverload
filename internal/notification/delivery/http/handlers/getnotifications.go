package handlers

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"

	"go-park-mail-ru/2022_2_BugOverload/internal/notification/service"
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
	r.HandleFunc("/api/v1/notifications", h.Action).
		Methods(http.MethodGet)
}

var upgrade = websocket.Upgrader{
	ReadBufferSize:  constparams.WSBufSizeRead,
	WriteBufferSize: constparams.WSBufSizeWrite,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func sendNewMsgNotifications(client *websocket.Conn, message []byte) {
	ticker := time.NewTicker(1 * time.Minute)
	for {
		err := client.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			break
		}

		<-ticker.C
	}
}

// func newMessage() []byte {
//	data, _ := json.Marshal(map[string]string{
//		"email":   "123",
//		"name":    "123" + " " + "123",
//		"subject": "123" + " " + "123",
//	})
//	return data
// }

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
		wrapper.DefaultHandlerHTTPError(r.Context(), w, errors.ErrUpdateWebSocketProtocol)
	}

	go sendNewMsgNotifications(connection, []byte("asd"))
}
