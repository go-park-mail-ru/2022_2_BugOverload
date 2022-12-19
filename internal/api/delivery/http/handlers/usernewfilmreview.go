package handlers

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/api/delivery/http/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	serviceUser "go-park-mail-ru/2022_2_BugOverload/internal/user/service"
	"net/http"

	"github.com/gorilla/mux"

	mainModels "go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/handler"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/middleware"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/wrapper"
)

// newFilmReviewHandler is the structure that handles the request for auth.
type newFilmReviewHandler struct {
	userService serviceUser.UserService
}

// NewFilmReviewHandler is constructor for filmRateHandler in this pkg - settings.
func NewFilmReviewHandler(us serviceUser.UserService) handler.Handler {
	return &newFilmReviewHandler{
		us,
	}
}

func (h *newFilmReviewHandler) Configure(r *mux.Router, mw *middleware.HTTPMiddleware) {
	r.HandleFunc("/api/v1/film/{id}/review/new", mw.NeedAuthMiddleware(mw.SetCsrfMiddleware(h.Action))).Methods(http.MethodPost)
}

// Action is a method for initial validation of the request and data and
// delivery of the data to the service at the business logic level.
// @Summary New film review
// @Description  New film review with body, name, type. Body, name, type required.
// @tags user
// @Produce json
// @Param   id    path  int    true "film id"
// @Param score body models.NewFilmReviewRequest true "Request body for rate film"
// @Success 201 {object} models.FilmReviewResponse "successfully added new review"
// @Failure 400 "return error"
// @Failure 401 {object} httpmodels.ErrResponseAuthNoCookie "no cookie"
// @Failure 404 {object} httpmodels.ErrResponseAuthNoSuchCookie "no such cookie"
// @Failure 405 "method not allowed"
// @Failure 500 "something unusual has happened"
// @Router /api/v1/film/{id}/review/new [POST]
func (h *newFilmReviewHandler) Action(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value(constparams.CurrentUserKey).(mainModels.User)
	if !ok {
		wrapper.DefaultHandlerHTTPError(r.Context(), w, errors.ErrGetUserRequest)
		return
	}

	request := models.NewNewFilmReviewRequest()

	err := request.Bind(r)
	if err != nil {
		wrapper.DefaultHandlerHTTPError(r.Context(), w, err)
		return
	}

	err = h.userService.NewFilmReview(r.Context(), &user, request.GetReview(), request.GetParams())
	if err != nil {
		wrapper.DefaultHandlerHTTPError(r.Context(), w, err)
		return
	}

	response := models.NewFilmReviewResponse(&user)

	wrapper.Response(r.Context(), w, http.StatusCreated, response)
}
