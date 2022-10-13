package authhandler

import (
	"context"
	"net/http"

	stdErrors "github.com/pkg/errors"

	authInterface "go-park-mail-ru/2022_2_BugOverload/internal/app/auth/interfaces"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/interfaces"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/user/delivery/http/models"
	userInterface "go-park-mail-ru/2022_2_BugOverload/internal/app/user/interfaces"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/utils/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/utils/httpwrapper"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/utils/params"
)

// handler is the structure that handles the request for auth.
type handler struct {
	userService userInterface.UserService
	authService authInterface.AuthService
}

// NewHandler is constructor for handler in this pkg - auth.
func NewHandler(us userInterface.UserService, as authInterface.AuthService) interfaces.Handler {
	return &handler{
		us,
		as,
	}
}

// Action is a method for initial validation of the request and data and
// delivery of the data to the service at the business logic level.
func (h *handler) Action(w http.ResponseWriter, r *http.Request) {
	var authRequest models.UserAuthRequest

	err := authRequest.Bind(w, r)
	if err != nil {
		httpwrapper.DefaultHandlerError(w, err)
		return
	}

	cookieStr := r.Header.Get("Cookie")

	ctx := context.WithValue(r.Context(), params.CookieKey, cookieStr)

	user, err := h.authService.GetUserBySession(ctx)
	if err != nil {
		httpwrapper.DefaultHandlerError(w, errors.NewErrAuth(stdErrors.Cause(err)))
		return
	}

	httpwrapper.Response(w, http.StatusOK, authRequest.ToPublic(&user))
}
