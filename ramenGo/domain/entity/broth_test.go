package entity_test

import (
	"testing"

	"github.com/ramenGo/domain/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestCreateBroth(t *testing.T) {
	broth, err := entity.NewBroth("someImageInactiveURL", "someImageActiveURL", "Chiken",
		"Caipira Chicken Broth", 22.567)

	assert.NotNil(t, broth)
	assert.Nil(t, err)
}

func TestValidadeBrothPriceZero(t *testing.T) {
	broth, err := entity.NewBroth("someImageInactiveURL", "someImageActiveURL", "Chiken",
		"Caipira Chicken Broth", 0)

	assert.Nil(t, broth)
	assert.Error(t, err)
	assert.EqualError(t, err, "Price cant be zero")
}

func TestValidadeBrothEmptyDescription(t *testing.T) {
	broth, err := entity.NewBroth("someImageInactiveURL", "someImageActiveURL", "Chiken",
		"", 0)

	assert.Nil(t, broth)
	assert.Error(t, err)
	assert.EqualError(t, err, "Description cant be empty string")
}
