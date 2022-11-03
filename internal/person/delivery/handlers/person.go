package handlers

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"net/http"

	servicePersons "go-park-mail-ru/2022_2_BugOverload/internal/film/service"
)

// personHandler is the structure that handles the request for
// getting film by id.
type personHandler struct {
	personService servicePersons.FilmsService
}

// NewPersonHandler is constructor for personHandler in this pkg - film.
func NewPersonHandler(fs servicePersons.FilmsService) pkg.Handler {
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
	// in dev
	//  vars := mux.Vars(r)
}
