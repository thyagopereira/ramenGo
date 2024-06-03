package mocks

import (
	"github.com/ramenGo/domain/entity"
	"github.com/stretchr/testify/mock"
)

type OrderDBMock struct {
	mock.Mock
}

func (m *OrderDBMock) FindById(id string) (entity.Entity, error) {
	args := m.Called(id)
	return args.Get(0).(entity.Entity), args.Error(1)
}

func (m *OrderDBMock) Save(e entity.Entity) error {
	args := m.Called(e)
	return args.Error(0)
}

func (m *OrderDBMock) GetAll() ([]entity.Entity, error) {
	args := m.Called()
	return args.Get(0).([]entity.Entity), args.Error(1)
}
