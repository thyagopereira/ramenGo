package usecases_test

import (
	"testing"

	"github.com/ramenGo/domain/mocks"
	"github.com/ramenGo/domain/usecases"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateOrderUc(t *testing.T) {

	databaseMock := &mocks.OrderDBMock{}
	uc, err := usecases.NewCreateOrderUseCase(databaseMock)
	assert.Nil(t, err)
	assert.NotNil(t, uc)

	request := &usecases.CreateOrderRequestDTO{
		BrothId:   "someBrothId",
		ProteinId: "someProteinId",
	}

	databaseMock.On("Save", mock.Anything).Return(nil)

	response, err := uc.Execute(request)
	assert.Nil(t, err)
	assert.NotNil(t, response)

	assert.NotEmpty(t, response.Id)
	assert.NotEmpty(t, response.Description)
	assert.NotEmpty(t, response.Image)

}
