package entity_test

import (
	"testing"

	"github.com/ramenGo/domain/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestCreateProtein(t *testing.T) {
	protein, err := entity.NewProtein("someImageInactiveURL", "someImageActiveURL", "Chiken",
		"Caipira Chicken Beef", 7.124)

	assert.NotNil(t, protein)
	assert.Nil(t, err)
}

func TestValidadeProteinPriceZero(t *testing.T) {
	protein, err := entity.NewProtein("someImageInactiveURL", "someImageActiveURL", "Chiken",
		"Caipira Chicken Beef", 0)

	assert.Nil(t, protein)
	assert.Error(t, err)
	assert.EqualError(t, err, "Price cant be zero")
}

func TestValidadeProteinEmptyDescription(t *testing.T) {
	protein, err := entity.NewProtein("someImageInactiveURL", "someImageActiveURL", "Chiken",
		"", 0)

	assert.Nil(t, protein)
	assert.Error(t, err)
	assert.EqualError(t, err, "Description cant be empty string")
}
