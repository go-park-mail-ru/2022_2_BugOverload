package signuphandler

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/app/utils/errors"
	"net/http"

	stdErrors "github.com/pkg/errors"

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

	user, err := h.userService.Signup(r.Context(), signupRequest.GetUser())
	if err != nil {
		httpwrapper.DefaultHandlerError(w, errors.NewErrAuth(stdErrors.Cause(err)))
		return
	}

	newSession, err := h.authService.CreateSession(r.Context(), &user)
	if err != nil {
		httpwrapper.DefaultHandlerError(w, err)
		return
	}

	w.Header().Set("Set-Cookie", newSession)

	httpwrapper.Response(w, http.StatusCreated, signupRequest.ToPublic(&user))
}
