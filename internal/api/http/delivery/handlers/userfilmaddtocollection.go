package handlers

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/api/http/delivery/models"
	mainModels "go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	serviceUser "go-park-mail-ru/2022_2_BugOverload/internal/user/service"
	"net/http"

	"github.com/gorilla/mux"

	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/handler"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/middleware"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/wrapper"
)

// addFilmToUserCollectionHandler is the structure that handles the request for
// adding film to collection.
type addFilmToUserCollectionHandler struct {
	userService serviceUser.UserService
}

// NewAddFilmToUserCollectionHandler is constructor for addFilmToUserCollectionHandler in this pkg
func NewAddFilmToUserCollectionHandler(us serviceUser.UserService) handler.Handler {
	return &addFilmToUserCollectionHandler{
		us,
	}
}

func (h *addFilmToUserCollectionHandler) Configure(r *mux.Router, mw *middleware.HTTPMiddleware) {
	r.HandleFunc("/api/v1/film/{id:[0-9]+}/save", mw.CheckAuthMiddleware(mw.SetCsrfMiddleware(h.Action))).
		Methods(http.MethodPost)
}

// Action is a method for initial validation of the request and data and
// delivery of the data to the service at the business logic level.
// @Summary Add film to user collection
// @Description Add film to user collection. User should be author of collection.
// @tags user
// @Produce json
// @Param   id    path  int    true "film id"
// @Param idCollection body models.AddFilmToUserCollectionRequest true "Request body for add film to user collection"
// @Success 204 "success"
// @Failure 401 {object} httpmodels.ErrResponseAuthNoCookie "no cookie"
// @Failure 404 {object} httpmodels.ErrResponseAuthNoSuchCookie "no such cookie"
// @Failure 405 "method not allowed"
// @Failure 500 "something unusual has happened"
// @Router /api/v1/film/{id}/save [POST]
func (h *addFilmToUserCollectionHandler) Action(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value(constparams.CurrentUserKey).(mainModels.User)
	if !ok {
		wrapper.DefaultHandlerHTTPError(r.Context(), w, errors.ErrGetUserRequest)
		return
	}

	request := models.NewAddFilmRequest()

	err := request.Bind(r)
	if err != nil {
		wrapper.DefaultHandlerHTTPError(r.Context(), w, err)
		return
	}

	err = h.userService.AddFilmToUserCollection(r.Context(), &user, request.GetParams())
	if err != nil {
		wrapper.DefaultHandlerHTTPError(r.Context(), w, err)
		return
	}

	wrapper.NoBody(w, http.StatusNoContent)
}
