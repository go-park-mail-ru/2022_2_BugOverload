package signup_handler

import (
	"go-park-mail-ru/2022_2_BugOverload/OLD/application/database"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/auth/delivery/http/models"
	"net/http"

	"go-park-mail-ru/2022_2_BugOverload/OLD/application/errors"
	"go-park-mail-ru/2022_2_BugOverload/OLD/application/httpwrapper"
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
	var signupRequest models.UserSignupRequest

	err := signupRequest.Bind(w, r)
	if err != nil {
		httpwrapper.DefaultHandlerError(w, err)
		return
	}

	user := signupRequest.GetUser()

	suchUserExist := ha.userStorage.CheckExist(user.Email)
	if suchUserExist {
		httpwrapper.DefaultHandlerError(w, errors.NewErrAuth(errors.ErrSignupUserExist))
		return
	}

	ha.userStorage.Create(*user)

	newCookie := ha.cookieStorage.Create(user.Email)

	w.Header().Set("Set-Cookie", newCookie)

	httpwrapper.Response(w, http.StatusCreated, signupRequest.ToPublic(user))
}
