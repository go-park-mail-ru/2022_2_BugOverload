package handlers

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/api/http/delivery/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/image/delivery/grpc/client"
	"net/http"

	"github.com/gorilla/mux"

	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/handler"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/middleware"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/wrapper"
)

// getImageHandler is the structure that handles the request for auth.
type getImageHandler struct {
	imageService client.ImageService
}

// NewGetImageHandler is constructor for getImageHandler in this pkg - auth.
func NewGetImageHandler(service client.ImageService) handler.Handler {
	return &getImageHandler{
		service,
	}
}

func (h *getImageHandler) Configure(r *mux.Router, mw *middleware.HTTPMiddleware) {
	r.HandleFunc("/api/v1/image", h.Action).
		Methods(http.MethodGet).
		Queries("object", "{object}", "key", "{key}")
}

// Action is a method for initial validation of the request and data and
// delivery of the data to the service at the business logic level.
// @Summary Download image
// @Description Rule for create type object NameEssence_NameAttribute. Examples for object: "film_poster_hor", "user_avatar", "film_image",
// @Description "default", "user_avatar", "person_avatar", "person_image", "collection_poster". Rule for film image: key = filmID/filmImageKey. Example 1/2.
// @Description For login key - login, for signup key - signup, both with type default. Object, key - required.
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
		wrapper.DefaultHandlerHTTPError(r.Context(), w, err)
		return
	}

	getImage, err := h.imageService.GetImage(r.Context(), request.GetImage())
	if err != nil {
		wrapper.DefaultHandlerHTTPError(r.Context(), w, err)
		return
	}

	wrapper.ResponseImage(r.Context(), w, http.StatusOK, getImage.Bytes)
}
