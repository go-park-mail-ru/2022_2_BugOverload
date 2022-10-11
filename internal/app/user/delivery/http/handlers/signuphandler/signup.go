package signuphandler

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/app/interfaces"
	"net/http"

	stdErrors "github.com/pkg/errors"

	authInterface "go-park-mail-ru/2022_2_BugOverload/internal/app/auth/interfaces"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/user/delivery/http/models"
	userInterface "go-park-mail-ru/2022_2_BugOverload/internal/app/user/interfaces"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/utils/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/utils/httpwrapper"
)

type handler struct {
	userService userInterface.UserService
	authService authInterface.AuthService
}

func NewHandler(us userInterface.UserService, as authInterface.AuthService) interfaces.Handler {
	return &handler{
		us,
		as,
	}
}

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
		httpwrapper.DefaultHandlerError(w, errors.NewErrAuth(stdErrors.Cause(err)))
		return
	}

	w.Header().Set("Set-Cookie", newSession)

	httpwrapper.Response(w, http.StatusCreated, signupRequest.ToPublic(&user))
}
