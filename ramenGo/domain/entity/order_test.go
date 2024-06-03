package entity_test

import (
	"testing"

	"github.com/ramenGo/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestCreateOrder(t *testing.T) {
	order, err := entity.NewOrder("someBrothId", "someProteinId", "someImage", "someDescription")

	assert.NotNil(t, order)
	assert.Nil(t, err)
}

func TestValidateOrderEmptyBrothId(t *testing.T) {
	order, err := entity.NewOrder("", "someProteinId", "someImage", "someDescription")

	assert.Nil(t, order)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "BrothId cant be empty string")

}

func TestValidateOrderEmptyProteinId(t *testing.T) {
	order, err := entity.NewOrder("someBrothId", "", "someImage", "someDescription")

	assert.Nil(t, order)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "ProteinId cant be empty string")

}
