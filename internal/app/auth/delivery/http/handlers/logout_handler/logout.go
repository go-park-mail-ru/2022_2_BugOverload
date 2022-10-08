package logout_handler

import (
	"net/http"

	"go-park-mail-ru/2022_2_BugOverload/OLD/application/database"
	"go-park-mail-ru/2022_2_BugOverload/OLD/application/errors"
	"go-park-mail-ru/2022_2_BugOverload/OLD/application/httpwrapper"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/auth/delivery/http/models"
)

// Handler is structure for API auth, login and signup processing
type Handler struct {
	userStorage   *database.UserStorage
	cookieStorage *database.CookieStorage
}

// NewHandler is constructor for Handler
func NewHandler(us *database.UserStorage, cs *database.CookieStorage) *Handler {
	return &Handler{
		us,
		cs,
	}
}

// Action is handling request for check current client cookie and return user data
func (ha *Handler) Action(w http.ResponseWriter, r *http.Request) {
	var logoutRequest models.UserLogoutRequest

	err := logoutRequest.Bind(w, r)
	if err != nil {
		httpwrapper.DefaultHandlerError(w, err)
		return
	}

	cookieStr := r.Header.Get("Cookie")

	badCookie, err := ha.cookieStorage.DeleteCookie(cookieStr)
	if err != nil {
		httpwrapper.DefaultHandlerError(w, errors.NewErrAuth(err))
		return
	}

	w.Header().Set("Set-Cookie", badCookie)

	httpwrapper.NoContent(w)
}
