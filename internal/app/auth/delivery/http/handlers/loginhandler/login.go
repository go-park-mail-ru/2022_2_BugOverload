package loginhandler

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/app/auth/repository/memory"
	errors2 "go-park-mail-ru/2022_2_BugOverload/internal/app/utils/errors"
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

// Action is handling request
func (ha *Handler) Action(w http.ResponseWriter, r *http.Request) {
	var loginRequest models.UserLoginRequest

	err := loginRequest.Bind(w, r)
	if err != nil {
		httpwrapper2.DefaultHandlerError(w, err)
		return
	}

	user := loginRequest.GetUser()

	userFromDB, err := ha.userStorage.GetUser(user.Email)
	if err != nil {
		httpwrapper2.DefaultHandlerError(w, errors2.NewErrAuth(err))

		return
	}

	if userFromDB.Password != user.Password {
		httpwrapper2.DefaultHandlerError(w, errors2.NewErrAuth(errors2.ErrLoginCombinationNotFound))

		return
	}

	newCookie := ha.cookieStorage.Create(user.Email)

	w.Header().Set("Set-Cookie", newCookie)

	httpwrapper2.Response(w, http.StatusOK, loginRequest.ToPublic(&userFromDB))
}
