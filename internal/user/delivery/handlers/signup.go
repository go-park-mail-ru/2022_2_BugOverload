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

// signupHandler is the structure that handles the request for auth.
type signupHandler struct {
	userService serviceUser.UserService
	authService serviceAuth.AuthService
}

// NewSingUpHandler is constructor for signupHandler in this pkg - auth.
func NewSingUpHandler(us serviceUser.UserService, as serviceAuth.AuthService) interfaces.Handler {
	return &signupHandler{
		us,
		as,
	}
}

// Action is a method for initial validation of the request and data and
// delivery of the data to the service at the business logic level.
func (h *signupHandler) Action(w http.ResponseWriter, r *http.Request) {
	signupRequest := models.NewUserSignupRequest()

	err := signupRequest.Bind(r)
	if err != nil {
		httpwrapper2.DefaultHandlerError(w, err)
		return
	}

	user, err := h.userService.Signup(r.Context(), signupRequest.GetUser())
	if err != nil {
		httpwrapper2.DefaultHandlerError(w, errors.NewErrAuth(stdErrors.Cause(err)))
		return
	}

	newSession, err := h.authService.CreateSession(r.Context(), &user)
	if err != nil {
		httpwrapper2.DefaultHandlerError(w, errors.NewErrAuth(stdErrors.Cause(err)))
		return
	}

	w.Header().Set("Set-Cookie", newSession)

	httpwrapper2.Response(w, http.StatusCreated, signupRequest.ToPublic(&user))
}
