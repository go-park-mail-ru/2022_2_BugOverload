package handlers

import (
	httpwrapper2 "go-park-mail-ru/2022_2_BugOverload/internal/pkg/httpwrapper"
	"net/http"

	stdErrors "github.com/pkg/errors"

	serviceAuth "go-park-mail-ru/2022_2_BugOverload/internal/auth/service"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/interfaces"
	"go-park-mail-ru/2022_2_BugOverload/internal/user/delivery/models"
	serviceUser "go-park-mail-ru/2022_2_BugOverload/internal/user/service"
)

// loginHandler is the structure that handles the request for auth.
type loginHandler struct {
	userService serviceUser.UserService
	authService serviceAuth.AuthService
}

// NewLoginHandler is constructor for loginHandler in this pkg - auth.
func NewLoginHandler(us serviceUser.UserService, as serviceAuth.AuthService) interfaces.Handler {
	return &loginHandler{
		us,
		as,
	}
}

// Action is a method for initial validation of the request and data and
// delivery of the data to the service at the business logic level.
func (h *loginHandler) Action(w http.ResponseWriter, r *http.Request) {
	loginRequest := models.NewUserLoginRequest()

	err := loginRequest.Bind(r)
	if err != nil {
		httpwrapper2.DefaultHandlerError(w, err)
		return
	}

	user := loginRequest.GetUser()

	userLogged, err := h.userService.Login(r.Context(), user)
	if err != nil {
		httpwrapper2.DefaultHandlerError(w, errors.NewErrAuth(stdErrors.Cause(err)))
		return
	}

	newSession, err := h.authService.CreateSession(r.Context(), &userLogged)
	if err != nil {
		httpwrapper2.DefaultHandlerError(w, errors.NewErrAuth(stdErrors.Cause(err)))
		return
	}

	w.Header().Set("Set-Cookie", newSession)

	httpwrapper2.Response(w, http.StatusOK, loginRequest.ToPublic(&userLogged))
}
