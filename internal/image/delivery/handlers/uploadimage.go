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

// uploadImageHandler is the structure that handles the request for auth.
type uploadImageHandler struct {
	imageService serviceImage.ImageService
}

// NewUploadImageHandler is constructor for uploadImageHandler in this pkg - auth.
func NewUploadImageHandler(is serviceImage.ImageService) pkg.Handler {
	return &uploadImageHandler{
		is,
	}
}

// Action is a method for initial validation of the request and data and
// delivery of the data to the service at the business logic level.
// @Summary Upload image
// @Description Rule for create type object NameEssence_NameAttribute. Examples: "film_poster_hor", "user_avatar"
// @tags image
// @produce json
// @Param   object    query  string  true  "type object"
// @Param   key       query  string  true  "key for found"
// @Success 200 "successfully upload"
// @Failure 400 {object} httpmodels.ErrResponseImageDefault "return error"
// @Failure 405 "method not allowed"
// @Failure 500 "something unusual has happened"
// @Router /api/v1/image [POST]
func (h *uploadImageHandler) Action(w http.ResponseWriter, r *http.Request) {
	getImageRequest := models.NewUploadImageRequest()

	err := getImageRequest.Bind(r)
	if err != nil {
		httpwrapper.DefaultHandlerError(w, err)
		return
	}

	image := getImageRequest.GetImage()

	err = h.imageService.UploadImage(r.Context(), image)
	if err != nil {
		httpwrapper.DefaultHandlerError(w, errors.NewErrImages(stdErrors.Cause(err)))
		return
	}

	httpwrapper.NoContent(w)
}
