package usecases_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/ramenGo/domain/internal/entity"
	"github.com/ramenGo/domain/mocks"
	"github.com/ramenGo/domain/usecases"
	"github.com/stretchr/testify/assert"
)

func TestListProteinsUcExec(t *testing.T) {

	proteins := []entity.Entity{
		&entity.Protein{
			Id:            uuid.NewString(),
			ImageInactive: "imageInactive",
			ImageActive:   "ImageActive",
			Name:          "Delicious Protein",
			Description:   "descriptive about the Protein",
			Price:         11.9,
		},
		&entity.Protein{
			Id:            uuid.NewString(),
			ImageInactive: "imageInactive2",
			ImageActive:   "ImageActive2",
			Name:          "Delicious Protein 2",
			Description:   "descriptive about the protein 2",
			Price:         11.30,
		},
	}

	databaseMock := &mocks.ProteinDBMock{}
	databaseMock.On("GetAll").Return(proteins, nil)

	uc, err := usecases.NewListProteinsUseCase(databaseMock)
	assert.Nil(t, err)
	assert.NotNil(t, uc)

	expected := &usecases.ListProteinsDTO{
		Proteins: proteins,
	}

	output, err := uc.Execute()
	assert.Nil(t, err)
	assert.EqualValues(t, expected, output)

	databaseMock.AssertExpectations(t)
	databaseMock.AssertCalled(t, "GetAll")

}
