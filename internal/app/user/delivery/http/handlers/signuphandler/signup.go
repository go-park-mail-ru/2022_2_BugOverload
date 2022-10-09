package signuphandler

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/app/auth/delivery/http/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/auth/repository/memory"
	errors2 "go-park-mail-ru/2022_2_BugOverload/internal/app/utils/errors"
	httpwrapper2 "go-park-mail-ru/2022_2_BugOverload/internal/app/utils/httpwrapper"
	"net/http"
)

// handler is structure for API auth, login and signup processing
type handler struct {
	userStorage   *memory.UserStorage
	cookieStorage *memory.CookieStorage
}

// NewHandler is constructor for handler
func NewHandler(us *memory.UserStorage, cs *memory.CookieStorage) *handler {
	return &handler{
		us,
		cs,
	}
}

// Action is handling request
func (h *handler) Action(w http.ResponseWriter, r *http.Request) {
	var signupRequest models.UserSignupRequest

	err := signupRequest.Bind(w, r)
	if err != nil {
		httpwrapper2.DefaultHandlerError(w, err)
		return
	}

	user := signupRequest.GetUser()

	suchUserExist := h.userStorage.CheckExist(user.Email)
	if suchUserExist {
		httpwrapper2.DefaultHandlerError(w, errors2.NewErrAuth(errors2.ErrSignupUserExist))
		return
	}

	h.userStorage.Create(*user)

	newCookie := h.cookieStorage.Create(user.Email)

	w.Header().Set("Set-Cookie", newCookie)

	httpwrapper2.Response(w, http.StatusCreated, signupRequest.ToPublic(user))
}
