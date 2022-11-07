package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	stdErrors "github.com/pkg/errors"

	"go-park-mail-ru/2022_2_BugOverload/internal/person/delivery/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/person/service"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/handler"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/httpwrapper"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/middleware"
)

// personHandler is the structure that handles the request for
// getting film by id.
type personHandler struct {
	personService service.PersonService
}

// NewPersonHandler is constructor for personHandler in this pkg - film.
func NewPersonHandler(fs service.PersonService) handler.Handler {
	return &personHandler{
		fs,
	}
}

func (h *personHandler) Configure(r *mux.Router, mw *middleware.Middleware) {
	r.HandleFunc("/api/v1/person/{id:[0-9]+}", h.Action).
		Methods(http.MethodGet).
		Queries("count_films", "{count_films}", "count_images", "{count_images}")
}

// Action is a method for initial validation of the request and data and
// delivery of the data to the service at the business logic level.
// @Summary Person full info
// @Description Getting person info by id
// @tags completed
// @Produce json
// @Param id  path int true "person id"
// @Param count_films  query int true "count best films"
// @Param count_images  query int true "count images person"
// @Success 200 {object} models.PersonResponse "return person info with best films"
// @Failure 400 "return error"
// @Failure 404 {object} httpmodels.ErrResponsePersonNoSuchPerson "no such person"
// @Failure 405 "method not allowed"
// @Failure 500 "something unusual has happened"
// @Router /api/v1/person/{id} [GET]
func (h *personHandler) Action(w http.ResponseWriter, r *http.Request) {
	request := models.NewPersonRequest()

	err := request.Bind(r)
	if err != nil {
		httpwrapper.DefaultHandlerError(w, errors.NewErrValidation(stdErrors.Cause(err)))
		return
	}

	person, err := h.personService.GePersonByID(r.Context(), request.GetPerson(), request.GetParams())
	if err != nil {
		httpwrapper.DefaultHandlerError(w, errors.NewErrPerson(stdErrors.Cause(err)))
		return
	}

	response := models.NewPersonResponse(&person)

	httpwrapper.Response(w, http.StatusOK, response)
}
