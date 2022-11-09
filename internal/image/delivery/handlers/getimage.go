package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	stdErrors "github.com/pkg/errors"

	"go-park-mail-ru/2022_2_BugOverload/internal/image/delivery/models"
	serviceImage "go-park-mail-ru/2022_2_BugOverload/internal/image/service"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/handler"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/httpwrapper"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/middleware"
)

// getImageHandler is the structure that handles the request for auth.
type getImageHandler struct {
	imageService serviceImage.ImageService
}

// NewGetImageHandler is constructor for getImageHandler in this pkg - auth.
func NewGetImageHandler(is serviceImage.ImageService) handler.Handler {
	return &getImageHandler{
		is,
	}
}

func (h *getImageHandler) Configure(r *mux.Router, mw *middleware.Middleware) {
	r.HandleFunc("/api/v1/image", h.Action).
		Methods(http.MethodGet).
		Queries("object", "{object}", "key", "{key}")
}

// Action is a method for initial validation of the request and data and
// delivery of the data to the service at the business logic level.
// @Summary Download image
// @Description Rule for create type object NameEssence_NameAttribute. Examples: "film_poster_hor", "user_avatar", "film_image",
// @Description "default", "user_avatar", "person_avatar", "person_image". Rule for film image: key = filmID/filmImageKey. Example 1/2.
// @Description For login key - login, for signup key - signup, both with type default
// @tags image, completed
// @produce json
// @produce image/webp
// @Param   object    query  string  true  "type object"
// @Param   key       query  string  true  "key for found"
// @Success 200 "successfully getting"
// @Failure 400 {object} httpmodels.ErrResponseImageDefault "return error"
// @Failure 404 {object} httpmodels.ErrResponseImageNoSuchImage "such image not found"
// @Failure 405 "method not allowed"
// @Failure 500 "something unusual has happened"
// @Router /api/v1/image [GET]
func (h *getImageHandler) Action(w http.ResponseWriter, r *http.Request) {
	request := models.NewGetImageRequest()

	err := request.Bind(r)
	if err != nil {
		httpwrapper.DefaultHandlerError(w, errors.NewErrValidation(stdErrors.Cause(err)))
		return
	}

	getImage, err := h.imageService.GetImage(r.Context(), request.GetImage())
	if err != nil {
		httpwrapper.DefaultHandlerError(w, errors.NewErrImages(stdErrors.Cause(err)))
		errors.CreateLog(r.Context(), err)
		return
	}

	httpwrapper.ResponseImage(w, http.StatusOK, getImage.Bytes)
}
