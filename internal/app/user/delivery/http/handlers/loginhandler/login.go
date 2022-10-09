package loginhandler

import (
	memory2 "go-park-mail-ru/2022_2_BugOverload/internal/app/auth/repository/memory"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/user/repository/memory"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/utils/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/utils/httpwrapper"
	"net/http"

	"go-park-mail-ru/2022_2_BugOverload/internal/app/user/delivery/http/models"
)

// handler is structure for API auth, login and signup processing
type handler struct {
	userStorage   *memory.userRepo
	cookieStorage *memory2.cookieRepo
}

// NewHandler is constructor for handler
func NewHandler(us *memory.userRepo, cs *memory2.cookieRepo) *handler {
	return &handler{
		us,
		cs,
	}
}

// Action is handling request
func (h *handler) Action(w http.ResponseWriter, r *http.Request) {
	var loginRequest models.UserLoginRequest

	err := loginRequest.Bind(w, r)
	if err != nil {
		httpwrapper.DefaultHandlerError(w, err)
		return
	}

	user := loginRequest.GetUser()

	userFromDB, err := h.userStorage.Login(user.Email)
	if err != nil {
		httpwrapper.DefaultHandlerError(w, errors.NewErrAuth(err))
		return
	}

	if userFromDB.Password != user.Password {
		httpwrapper.DefaultHandlerError(w, errors.NewErrAuth(errors.ErrLoginCombinationNotFound))
		return
	}

	newCookie := h.cookieStorage.CreateSession(user.Email)

	w.Header().Set("Set-Cookie", newCookie)

	httpwrapper.Response(w, http.StatusOK, loginRequest.ToPublic(&userFromDB))
}
