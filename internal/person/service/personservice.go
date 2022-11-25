package service

import (
	"context"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"

	stdErrors "github.com/pkg/errors"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/person/repository"
)

//go:generate mockgen -source personservice.go -destination mocks/mockpersonservice.go -package mockPersonService

// PersonService provides universal service for work with persons.
type PersonService interface {
	GetPersonByID(ctx context.Context, person *models.Person, params *constparams.GetPersonParams) (models.Person, error)
}

// personService is implementation for users service corresponding to the PersonService interface.
type personService struct {
	personRepo repository.PersonRepository
}

// NewPersonService is constructor for personService.
func NewPersonService(pr repository.PersonRepository) PersonService {
	return &personService{
		personRepo: pr,
	}
}

// GetPersonByID is the service that accesses the interface PersonService.
func (p *personService) GetPersonByID(ctx context.Context, person *models.Person, params *constparams.GetPersonParams) (models.Person, error) {
	personRepo, err := p.personRepo.GetPersonByID(ctx, person, params)
	if err != nil {
		return models.Person{}, stdErrors.Wrap(err, "GetPersonByID")
	}

	return personRepo, nil
}
