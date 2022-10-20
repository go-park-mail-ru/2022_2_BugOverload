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
// @Summary Get image
// @Description Getting image by path
// @tags image
// @Accept json
// @Produce json
// @Param image body models.GetImageRequest true "Request body for get image"
// @Success 200 "successfully getting"
// @Failure 400 "return error"
// @Failure 404 "such image not found"
// @Failure 405 "method not allowed"
// @Failure 500 "something unusual has happened"
// @Router /v1/image [GET]
func (h *getImageHandler) Action(w http.ResponseWriter, r *http.Request) {
	getImageRequest := models.NewGetImageRequest()

	err := getImageRequest.Bind(r)
	if err != nil {
		httpwrapper.DefaultHandlerError(w, err)
		return
	}

	image := getImageRequest.GetImage()

	binImage, err := h.imageService.GetImage(r.Context(), image)
	if err != nil {
		httpwrapper.DefaultHandlerError(w, errors.NewErrImages(stdErrors.Cause(err)))
		return
	}

	httpwrapper.ResponseImage(w, http.StatusOK, binImage)
}
