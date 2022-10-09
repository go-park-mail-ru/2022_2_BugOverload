package logouthandler

import (
	memory2 "go-park-mail-ru/2022_2_BugOverload/internal/app/auth/repository/memory"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/user/repository/memory"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/utils/errors"
	httpwrapper2 "go-park-mail-ru/2022_2_BugOverload/internal/app/utils/httpwrapper"
	"net/http"

	"go-park-mail-ru/2022_2_BugOverload/internal/app/user/delivery/http/models"
)

// handler is structure for API auth, login and signup processing
type handler struct {
	userStorage   *memory.userRepo
	cookieStorage *memory2.memoryCookieRepo
}

// NewHandler is constructor for handler
func NewHandler(us *memory.userRepo, cs *memory2.memoryCookieRepo) *handler {
	return &handler{
		us,
		cs,
	}
}

// Action is handling request for check current client cookie and return user data
func (h *handler) Action(w http.ResponseWriter, r *http.Request) {
	var logoutRequest models.UserLogoutRequest

	err := logoutRequest.Bind(w, r)
	if err != nil {
		httpwrapper2.DefaultHandlerError(w, err)
		return
	}

	cookieStr := r.Header.Get("Cookie")

	badCookie, err := h.cookieStorage.DeleteCookie(cookieStr)
	if err != nil {
		httpwrapper2.DefaultHandlerError(w, errors.NewErrAuth(err))
		return
	}

	w.Header().Set("Set-Cookie", badCookie)

	httpwrapper2.NoContent(w)
}
