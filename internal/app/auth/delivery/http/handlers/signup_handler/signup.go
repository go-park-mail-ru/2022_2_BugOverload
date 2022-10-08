package signup_handler

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/app/auth/delivery/http/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/auth/repository/memory"
	errors2 "go-park-mail-ru/2022_2_BugOverload/internal/app/utils/errors"
	httpwrapper2 "go-park-mail-ru/2022_2_BugOverload/internal/app/utils/httpwrapper"
	"net/http"
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
	var signupRequest models.UserSignupRequest

	err := signupRequest.Bind(w, r)
	if err != nil {
		httpwrapper2.DefaultHandlerError(w, err)
		return
	}

	user := signupRequest.GetUser()

	suchUserExist := ha.userStorage.CheckExist(user.Email)
	if suchUserExist {
		httpwrapper2.DefaultHandlerError(w, errors2.NewErrAuth(errors2.ErrSignupUserExist))
		return
	}

	ha.userStorage.Create(*user)

	newCookie := ha.cookieStorage.Create(user.Email)

	w.Header().Set("Set-Cookie", newCookie)

	httpwrapper2.Response(w, http.StatusCreated, signupRequest.ToPublic(user))
}
