package auth_handler

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/app/auth/repository/memory"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/utils/errors"
	httpwrapper2 "go-park-mail-ru/2022_2_BugOverload/internal/app/utils/httpwrapper"
	"net/http"

	"go-park-mail-ru/2022_2_BugOverload/internal/app/auth/delivery/http/models"
)

// Handler is structure for API auth, login and signup processing
type Handler struct {
	userStorage   *memory.UserStorage
	cookieStorage *memory.CookieStorage
}

// NewHandler is constructor for Handler
func NewHandler(us *memory.UserStorage, cs *memory.CookieStorage) *Handler {
	return &Handler{
		us,
		cs,
	}
}

// Action is handling request for check current client cookie and return user data
func (ha *Handler) Action(w http.ResponseWriter, r *http.Request) {
	var authRequest models.UserAuthRequest

	err := authRequest.Bind(w, r)
	if err != nil {
		httpwrapper2.DefaultHandlerError(w, err)
		return
	}

	cookieStr := r.Header.Get("Cookie")

	cookie, err := ha.cookieStorage.GetCookie(cookieStr)
	if err != nil {
		httpwrapper2.DefaultHandlerError(w, errors.NewErrAuth(err))
		return
	}

	userFromDB, err := ha.userStorage.GetUser(cookie.Value)
	if err != nil {
		httpwrapper2.DefaultHandlerError(w, errors.NewErrAuth(err))
		return
	}

	httpwrapper2.Response(w, http.StatusOK, authRequest.ToPublic(&userFromDB))
}
