package popularfilmshandler

import (
	"net/http"

	stdErrors "github.com/pkg/errors"

	"go-park-mail-ru/2022_2_BugOverload/internal/app/collection/delivery/http/models"
	collectionInterface "go-park-mail-ru/2022_2_BugOverload/internal/app/collection/interfaces"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/interfaces"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/utils/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/utils/httpwrapper"
)

type handler struct {
	collectionService collectionInterface.CollectionService
}

func NewHandler(uc collectionInterface.CollectionService) interfaces.Handler {
	return &handler{
		uc,
	}
}

func (h *handler) Action(w http.ResponseWriter, r *http.Request) {
	collection, err := h.collectionService.GetPopular(r.Context())
	if err != nil {
		httpwrapper.DefaultHandlerError(w, errors.NewErrFilms(stdErrors.Cause(err)))
		return
	}

	collectionPopular := models.NewFilmsPopularRequest(collection)

	response := collectionPopular.CreateResponse()

	httpwrapper.Response(w, http.StatusOK, response)
}
