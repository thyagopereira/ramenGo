package usecases_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/ramenGo/domain/internal/entity"
	"github.com/ramenGo/domain/mocks"
	"github.com/ramenGo/domain/usecases"
	"github.com/stretchr/testify/assert"
)

func TestListBrothsUcExec(t *testing.T) {

	broths := []entity.Entity{
		&entity.Broth{
			Id:            uuid.NewString(),
			ImageInactive: "imageInactive",
			ImageActive:   "ImageActive",
			Name:          "Delicious Broth",
			Description:   "descriptive about the broth",
			Price:         11.9,
		},
		&entity.Broth{
			Id:            uuid.NewString(),
			ImageInactive: "imageInactive2",
			ImageActive:   "ImageActive2",
			Name:          "Delicious Broth 2",
			Description:   "descriptive about the broth 2",
			Price:         11.30,
		},
	}

	databaseMock := &mocks.BrothDBMock{}
	databaseMock.On("GetAll").Return(broths, nil)

	uc, err := usecases.NewListBrothsUseCase(databaseMock)
	assert.Nil(t, err)
	assert.NotNil(t, uc)

	expected := &usecases.ListBrothsDTO{
		Broths: broths,
	}

	output, err := uc.Execute()
	assert.Nil(t, err)
	assert.EqualValues(t, expected, output)

	databaseMock.AssertExpectations(t)
	databaseMock.AssertCalled(t, "GetAll")

}
