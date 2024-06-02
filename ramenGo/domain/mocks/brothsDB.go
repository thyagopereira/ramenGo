package mocks

import (
	"github.com/ramenGo/domain/internal/entity"
	"github.com/stretchr/testify/mock"
)

type BrothDBMock struct {
	mock.Mock
}

func (m *BrothDBMock) FindById(id string) (entity.Entity, error) {
	args := m.Called(id)
	return args.Get(0).(entity.Entity), args.Error(1)
}

func (m *BrothDBMock) Save(e entity.Entity) error {
	args := m.Called(e)
	return args.Error(0)
}

func (m *BrothDBMock) GetAll() ([]entity.Entity, error) {
	args := m.Called()
	return args.Get(0).([]entity.Entity), args.Error(1)
}
