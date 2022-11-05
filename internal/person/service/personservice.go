package service

import (
	"context"

	stdErrors "github.com/pkg/errors"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/person/repository"
)

// PersonService provides universal service for work with persons.
type PersonService interface {
	GePersonByID(ctx context.Context, person *models.Person) (models.Person, error)
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

// GePersonByID is the service that accesses the interface PersonService.
func (p *personService) GePersonByID(ctx context.Context, person *models.Person) (models.Person, error) {
	personRepo, err := p.personRepo.GetPersonByID(ctx, person)
	if err != nil {
		return models.Person{}, stdErrors.Wrap(err, "GetPersonByID")
	}

	return personRepo, nil
}
