package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"go-park-mail-ru/2022_2_BugOverload/internal/image/delivery/models"
	serviceImage "go-park-mail-ru/2022_2_BugOverload/internal/image/service"
	mainModels "go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/handler"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/middleware"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/wrapper"
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
	r.HandleFunc("/api/v1/image", mw.CheckAuthMiddleware(mw.SetCsrfMiddleware(h.Action))).
		Methods(http.MethodPut).
		Queries("object", "{object}")
}

// Action is a method for initial validation of the request and data and
// delivery of the data to the service at the business logic level.
// @Summary Change image
// @Description Rule for create type object NameEssence_NameAttribute. Examples: "film_poster_hor", "user_avatar".
// @Description Object, key - required. For user_avatar key not required.
// @tags image
// @produce json
// @Param   object    query  string  true  "type object"
// @Param   key       query  string  false  "key image"
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
		wrapper.DefaultHandlerHTTPError(r.Context(), w, err)
		return
	}

	user, ok := r.Context().Value(pkg.CurrentUserKey).(mainModels.User)
	if !ok {
		wrapper.DefaultHandlerHTTPError(r.Context(), w, err)
		return
	}

	image := request.GetImage()

	if !user.IsAdmin {
		if image.Object != pkg.ImageObjectUserAvatar {
			wrapper.DefaultHandlerHTTPError(r.Context(), w, err)
			return
		}

		image.Key = strconv.Itoa(user.ID)
	}

	err = h.imageService.UpdateImage(r.Context(), image)
	if err != nil {
		wrapper.DefaultHandlerHTTPError(r.Context(), w, wrapper.GRPCErrorConvert(err))
		return
	}

	wrapper.NoBody(w, http.StatusNoContent)
}
