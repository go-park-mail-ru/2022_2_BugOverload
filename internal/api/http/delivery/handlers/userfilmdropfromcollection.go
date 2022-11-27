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

// dropFilmFromUserCollectionHandler is the structure that handles the request for
// drop film from collection.
type dropFilmFromUserCollectionHandler struct {
	userService serviceUser.UserService
}

// NewDropFilmFromUserCollectionHandler is constructor for dropFilmFromUserCollectionHandler in this pkg
func NewDropFilmFromUserCollectionHandler(us serviceUser.UserService) handler.Handler {
	return &dropFilmFromUserCollectionHandler{
		us,
	}
}

func (h *dropFilmFromUserCollectionHandler) Configure(r *mux.Router, mw *middleware.HTTPMiddleware) {
	r.HandleFunc("/api/v1/film/{id:[0-9]+}/remove", mw.CheckAuthMiddleware(mw.SetCsrfMiddleware(h.Action))).
		Methods(http.MethodDelete)
}

// Action is a method for initial validation of the request and data and
// delivery of the data to the service at the business logic level.
// @Summary Remove film from user collection
// @Description Remove film from user collection. User should be author of collection.
// @tags user
// @Produce json
// @Param   id    path  int    true "film id"
// @Param idCollection body models.DropFilmFromUserCollectionRequest true "Request body for drop film from user collection"
// @Success 204 "success"
// @Failure 401 {object} httpmodels.ErrResponseAuthNoCookie "no cookie"
// @Failure 404 {object} httpmodels.ErrResponseAuthNoSuchCookie "no such cookie"
// @Failure 405 "method not allowed"
// @Failure 500 "something unusual has happened"
// @Router /api/v1/film/{id}/remove [DELETE]
func (h *dropFilmFromUserCollectionHandler) Action(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value(constparams.CurrentUserKey).(mainModels.User)
	if !ok {
		wrapper.DefaultHandlerHTTPError(r.Context(), w, errors.ErrGetUserRequest)
		return
	}

	request := models.NewDropFilmRequest()

	err := request.Bind(r)
	if err != nil {
		wrapper.DefaultHandlerHTTPError(r.Context(), w, err)
		return
	}

	err = h.userService.DropFilmFromUserCollection(r.Context(), &user, request.GetParams())
	if err != nil {
		wrapper.DefaultHandlerHTTPError(r.Context(), w, err)
		return
	}

	wrapper.NoBody(w, http.StatusNoContent)
}
