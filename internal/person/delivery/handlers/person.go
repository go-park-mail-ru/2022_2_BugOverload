package handlers

import (
	"context"
	"net/http"

	stdErrors "github.com/pkg/errors"

	"go-park-mail-ru/2022_2_BugOverload/internal/person/delivery/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/person/service"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/httpwrapper"
)

// personHandler is the structure that handles the request for
// getting film by id.
type personHandler struct {
	personService service.PersonService
}

// NewPersonHandler is constructor for personHandler in this pkg - film.
func NewPersonHandler(fs service.PersonService) pkg.Handler {
	return &personHandler{
		fs,
	}
}

// Action is a method for initial validation of the request and data and
// delivery of the data to the service at the business logic level.
// @Summary Person full info
// @Description Getting person info by id
// @tags in_dev
// @Produce json
// @Param id  path int true "person id"
// @Success 200 {object} models.PersonResponse "return person"
// @Failure 400 "return error"
// @Failure 404 {object} httpmodels.ErrResponsePersonNoSuchPerson "no such person"
// @Failure 405 "method not allowed"
// @Failure 500 "something unusual has happened"
// @Router /api/v1/person/{id} [GET]
func (h *personHandler) Action(w http.ResponseWriter, r *http.Request) {
	personRequest := models.NewPersonRequest()

	err := personRequest.Bind(r)
	if err != nil {
		httpwrapper.DefaultHandlerError(w, err)
		return
	}

	requestParams := pkg.GetPersonParamsCtx{
		CountFilms: personRequest.CountFilms,
	}

	ctx := context.WithValue(r.Context(), pkg.GetReviewsParamsKey, requestParams)

	person, err := h.personService.GePersonByID(ctx, personRequest.GetPerson())
	if err != nil {
		httpwrapper.DefaultHandlerError(w, errors.NewErrImages(stdErrors.Cause(err)))
		return
	}

	personResponse := models.NewPersonResponse(&person)

	httpwrapper.Response(w, http.StatusOK, personResponse)
}
