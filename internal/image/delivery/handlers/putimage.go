package handlers

import (
	mainModels "go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg"
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

// putImageHandler is the structure that handles the request for auth.
type putImageHandler struct {
	imageService serviceImage.ImageService
}

// NewPutImageHandler is constructor for putImageHandler in this pkg - auth.
func NewPutImageHandler(is serviceImage.ImageService) handler.Handler {
	return &putImageHandler{
		is,
	}
}

func (h *putImageHandler) Configure(r *mux.Router, mw *middleware.Middleware) {
	r.HandleFunc("/api/v1/image", mw.SetCsrfMiddleware(mw.CheckAuthMiddleware(h.Action))).Methods(http.MethodPut).Queries("object", "{object}", "key", "{key}")
}

// Action is a method for initial validation of the request and data and
// delivery of the data to the service at the business logic level.
// @Summary Change image
// @Description Rule for create type object NameEssence_NameAttribute. Examples: "film_poster_hor", "user_avatar"
// @tags completed_not_tested_waiting_integration_auth
// @produce json
// @Param   object    query  string  true  "type object"
// @Param   key       query  string  true  "key for found id for sessions"
// @Success 204 "successfully upload"
// @Failure 400 {object} httpmodels.ErrResponseImageDefault "return error"
// @Failure 401 "no cookie"
// @Failure 403 "no access"
// @Failure 405 "method not allowed"
// @Failure 500 "something unusual has happened"
// @Router /api/v1/image [PUT]
func (h *putImageHandler) Action(w http.ResponseWriter, r *http.Request) {
	request := models.NewPutImageRequest()

	err := request.Bind(r)
	if err != nil {
		httpwrapper.DefaultHandlerError(w, errors.NewErrValidation(stdErrors.Cause(err)))
		return
	}

	user, ok := r.Context().Value(pkg.CurrentUserKey).(*mainModels.User)
	if !ok {
		httpwrapper.DefaultHandlerError(w, errors.NewErrAuth(errors.ErrGetUserRequest))
		return
	}

	image := request.GetImage()

	if !user.IsAdmin && image.Object != pkg.ImageObjectUserAvatar {
		httpwrapper.DefaultHandlerError(w, errors.NewErrAccess(errors.ErrNoAccess))
		return
	}

	err = h.imageService.UploadImage(r.Context(), image)
	if err != nil {
		httpwrapper.DefaultHandlerError(w, errors.NewErrImages(stdErrors.Cause(err)))
		return
	}

	httpwrapper.NoBody(w, http.StatusNoContent)
}
