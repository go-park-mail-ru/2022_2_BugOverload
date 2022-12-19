package service_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	stdErrors "github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	mockPersonRepository "go-park-mail-ru/2022_2_BugOverload/internal/warehouse/repository/person/mocks"
	"go-park-mail-ru/2022_2_BugOverload/internal/warehouse/service"
)

func TestPersonService_GetPersonByID_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockPersonRepository.NewMockRepository(ctrl)

	// Data
	inputPerson := &models.Person{
		ID:   1,
		Name: "test",
	}

	inputParams := &constparams.GetPersonParams{
		CountImages: 1,
		CountFilms:  1,
	}

	output := models.Person{
		ID:     1,
		Avatar: "avatar",
	}

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().GetPersonByID(ctx, inputPerson, inputParams).Return(output, nil)

	// Action
	personService := service.NewPersonService(repository)

	actual, err := personService.GetPersonByID(ctx, inputPerson, inputParams)

	// Check success
	require.Nil(t, err, "Handling must be without errors")

	// Check result handling
	require.Equal(t, output, actual)
}

func TestPersonService_GetPersonByID_NOT_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockPersonRepository.NewMockRepository(ctrl)

	// Data
	inputPerson := &models.Person{
		ID:   1,
		Name: "test",
	}

	inputParams := &constparams.GetPersonParams{
		CountImages: 1,
		CountFilms:  1,
	}

	expectedErr := errors.ErrPersonNotFound

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().GetPersonByID(ctx, inputPerson, inputParams).Return(models.Person{}, expectedErr)

	// Action
	personService := service.NewPersonService(repository)

	_, actualErr := personService.GetPersonByID(ctx, inputPerson, inputParams)

	// Check success
	require.NotNil(t, actualErr, "Handling must be error")

	// Check result handling
	require.Equal(t, stdErrors.Cause(expectedErr), stdErrors.Cause(actualErr))
}
