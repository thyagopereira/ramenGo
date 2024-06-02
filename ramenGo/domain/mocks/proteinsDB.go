package mocks

import (
	"github.com/ramenGo/domain/entity"
	"github.com/stretchr/testify/mock"
)

type ProteinDBMock struct {
	mock.Mock
}

func (m *ProteinDBMock) FindById(id string) (entity.Entity, error) {
	args := m.Called(id)
	return args.Get(0).(entity.Entity), args.Error(1)
}

func (m *ProteinDBMock) Save(e entity.Entity) error {
	args := m.Called(e)
	return args.Error(0)
}

func (m *ProteinDBMock) GetAll() ([]entity.Entity, error) {
	args := m.Called()
	return args.Get(0).([]entity.Entity), args.Error(1)
}
