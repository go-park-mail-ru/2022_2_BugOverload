package handlers

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/api/delivery/http/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"net/http"

	"github.com/gorilla/mux"

	serviceImage "go-park-mail-ru/2022_2_BugOverload/internal/image/service"
	mainModels "go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/handler"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/middleware"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/wrapper"
)

// postImageHandler is the structure that handles the request for auth.
type postImageHandler struct {
	imageService serviceImage.ImageService
}

// NewPostImageHandler is constructor for putImageHandler in this pkg - auth.
func NewPostImageHandler(is serviceImage.ImageService) handler.Handler {
	return &postImageHandler{
		is,
	}
}

func (h *postImageHandler) Configure(r *mux.Router, mw *middleware.HTTPMiddleware) {
	r.HandleFunc("/api/v1/image", mw.NeedAuthMiddleware(mw.SetCsrfMiddleware(h.Action))).
		Methods(http.MethodPost).
		Queries("object", "{object}", "key", "{key}")
}

// Action is a method for initial validation of the request and data and
// delivery of the data to the service at the business logic level.
// @Summary Upload image
// @Description Rule for create type object NameEssence_NameAttribute. Examples: "film_poster_hor", "user_avatar". Object, key - required.
// @tags image, not_actual_completed_not_tested_waiting_integration_auth
// @produce json
// @Param   object    query  string  true  "type object"
// @Param   key       query  string  true  "key for found"
// @Success 201 "successfully upload"
// @Failure 400 {object} httpmodels.ErrResponseImageDefault "return error"
// @Failure 401 "no cookie"
// @Failure 403 "no access"
// @Failure 405 "method not allowed"
// @Failure 500 "something unusual has happened"
// @Router /api/v1/image [POST]
func (h *postImageHandler) Action(w http.ResponseWriter, r *http.Request) {
	request := models.NewPostImageRequest()

	err := request.Bind(r)
	if err != nil {
		wrapper.DefaultHandlerHTTPError(r.Context(), w, err)
		return
	}

	user, ok := r.Context().Value(constparams.CurrentUserKey).(mainModels.User)
	if !ok {
		wrapper.DefaultHandlerHTTPError(r.Context(), w, errors.ErrGetUserRequest)
		return
	}

	if !user.IsAdmin {
		wrapper.DefaultHandlerHTTPError(r.Context(), w, errors.ErrGetUserRequest)
		return
	}

	err = h.imageService.UpdateImage(r.Context(), request.GetImage())
	if err != nil {
		wrapper.DefaultHandlerHTTPError(r.Context(), w, err)
		return
	}

	wrapper.NoBody(w, http.StatusCreated)
}
