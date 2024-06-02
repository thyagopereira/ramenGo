package usecases

import (
	"github.com/ramenGo/domain/internal/database"
	"github.com/ramenGo/domain/internal/entity"
)

type ListBrothsUseCase struct {
	database database.Database
}

type ListBrothsDTO struct {
	Broths []entity.Entity
}

func NewListBrothsUseCase(db database.Database) (*ListBrothsUseCase, error) {

	return &ListBrothsUseCase{
		database: db,
	}, nil
}

func (uc *ListBrothsUseCase) Execute() (*ListBrothsDTO, error) {
	result, err := uc.database.GetAll()
	if err != nil {
		return nil, err
	}

	return &ListBrothsDTO{
		Broths: result,
	}, nil
}
