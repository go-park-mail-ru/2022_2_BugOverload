package signuphandler

import (
	"context"
	"net/http"

	authInterface "go-park-mail-ru/2022_2_BugOverload/internal/app/auth/interfaces"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/user/delivery/http/models"
	userInterface "go-park-mail-ru/2022_2_BugOverload/internal/app/user/interfaces"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/utils/httpwrapper"
)

// handler is structure for API auth, login and signup processing
type handler struct {
	userService userInterface.UserService
	authService authInterface.AuthService
}

// NewHandler is constructor for handler
func NewHandler(us userInterface.UserService, as authInterface.AuthService) *handler {
	return &handler{
		us,
		as,
	}
}

// Action is handling request
func (h *handler) Action(w http.ResponseWriter, r *http.Request) {
	var signupRequest models.UserSignupRequest

	err := signupRequest.Bind(w, r)
	if err != nil {
		httpwrapper.DefaultHandlerError(w, err)
		return
	}

	var ctx context.Context

	user, err := h.userService.Signup(ctx, signupRequest.GetUser())
	if err != nil {
		httpwrapper.DefaultHandlerError(w, err)
		return
	}

	newCookie, err := h.authService.CreateSession(ctx, &user)
	if err != nil {
		httpwrapper.DefaultHandlerError(w, err)
		return
	}

	w.Header().Set("Set-Cookie", newCookie)

	httpwrapper.Response(w, http.StatusCreated, signupRequest.ToPublic(&user))
}
