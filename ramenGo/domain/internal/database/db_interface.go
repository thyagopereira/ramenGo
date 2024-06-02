package database

import "github.com/ramenGo/domain/internal/entity"

type Database interface {
	FindById(id string) (*entity.Entity, error)
	Save(e *entity.Entity) error
	GetAll() ([]entity.Entity, error)
}
