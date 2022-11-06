package handlers

import (
	"net/http"

	stdErrors "github.com/pkg/errors"

	"go-park-mail-ru/2022_2_BugOverload/internal/image/delivery/models"
	serviceImage "go-park-mail-ru/2022_2_BugOverload/internal/image/service"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/httpwrapper"
)

// getImageHandler is the structure that handles the request for auth.
type getImageHandler struct {
	imageService serviceImage.ImageService
}

// NewGetImageHandler is constructor for getImageHandler in this pkg - auth.
func NewGetImageHandler(is serviceImage.ImageService) pkg.Handler {
	return &getImageHandler{
		is,
	}
}

// Action is a method for initial validation of the request and data and
// delivery of the data to the service at the business logic level.
// @Summary Download image
// @Description Rule for create type object NameEssence_NameAttribute. Examples: "film_poster_hor", "user_avatar", "film_image",
// @Description "default", "user_avatar", "person_avatar", "person_image". Rule for film image: key = filmID/filmImageKey. Example 1/2
// @tags completed
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
		httpwrapper.DefaultHandlerError(w, err)
		return
	}

	getImage, err := h.imageService.DownloadImage(r.Context(), request.GetImage())
	if err != nil {
		httpwrapper.DefaultHandlerError(w, errors.NewErrImages(stdErrors.Cause(err)))
		return
	}

	httpwrapper.ResponseImage(w, http.StatusOK, getImage.Bytes)
}
