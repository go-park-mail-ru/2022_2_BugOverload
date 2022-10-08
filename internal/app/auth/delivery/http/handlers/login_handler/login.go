package login_handler

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

// Action is handling request
func (ha *Handler) Action(w http.ResponseWriter, r *http.Request) {
	var loginRequest models.UserLoginRequest

	err := loginRequest.Bind(w, r)
	if err != nil {
		httpwrapper.DefaultHandlerError(w, err)
		return
	}

	user := loginRequest.GetUser()

	userFromDB, err := ha.userStorage.GetUser(user.Email)
	if err != nil {
		httpwrapper.DefaultHandlerError(w, errors.NewErrAuth(err))

		return
	}

	if userFromDB.Password != user.Password {
		httpwrapper.DefaultHandlerError(w, errors.NewErrAuth(errors.ErrLoginCombinationNotFound))

		return
	}

	newCookie := ha.cookieStorage.Create(user.Email)

	w.Header().Set("Set-Cookie", newCookie)

	httpwrapper.Response(w, http.StatusOK, loginRequest.ToPublic(&userFromDB))
}
