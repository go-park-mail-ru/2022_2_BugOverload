package handlers

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"net/http"
	"time"

	stdErrors "github.com/pkg/errors"

	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/httpwrapper"
	serviceAuth "go-park-mail-ru/2022_2_BugOverload/internal/session/service"
	"go-park-mail-ru/2022_2_BugOverload/internal/user/delivery/models"
	serviceUser "go-park-mail-ru/2022_2_BugOverload/internal/user/service"
)

// signupHandler is the structure that handles the request for auth.
type signupHandler struct {
	userService serviceUser.UserService
	authService serviceAuth.SessionService
}

// NewSingUpHandler is constructor for signupHandler in this pkg - auth.
func NewSingUpHandler(us serviceUser.UserService, as serviceAuth.SessionService) pkg.Handler {
	return &signupHandler{
		us,
		as,
	}
}

// Action is a method for initial validation of the request and data and
// delivery of the data to the service at the business logic level.
// @Summary New user registration
// @Description Sending login and password for registration
// @tags user
// @Accept json
// @Produce json
// @Param user body models.UserSignupRequest true "Request body for signup"
// @Success 201 {object} models.UserSignupResponse "successfully signup"
// @Failure 400 {object} httpmodels.ErrResponseAuthDefault "return error"
// @Failure 404 {object} httpmodels.ErrResponseAuthNoSuchUser "such user not found"
// @Failure 405 "method not allowed"
// @Failure 500 "something unusual has happened"
// @Router /v1/auth/signup [POST]
func (h *signupHandler) Action(w http.ResponseWriter, r *http.Request) {
	signupRequest := models.NewUserSignupRequest()

	err := signupRequest.Bind(r)
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

	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    newSession,
		Expires:  time.Now().Add(pkg.TimeoutLiveCookie),
		HttpOnly: true,
	}

	http.SetCookie(w, cookie)

	signupResponse := models.NewUserSignUpResponse()

	httpwrapper.Response(w, http.StatusCreated, signupResponse.ToPublic(&user))
}
