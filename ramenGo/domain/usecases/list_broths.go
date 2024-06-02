package usecases

import (
	databases "github.com/ramenGo/domain/database"
	"github.com/ramenGo/domain/entity"
)

type ListBrothsUseCase struct {
	database databases.Database
}

type ListBrothsDTO struct {
	Broths []entity.Entity
}

func NewListBrothsUseCase(db databases.Database) (*ListBrothsUseCase, error) {

	return &ListBrothsUseCase{
		database: db,
	}, nil
}

func (uc *ListBrothsUseCase) Execute() (interface{}, error) {
	result, err := uc.database.GetAll()
	if err != nil {
		return nil, err
	}

	return &ListBrothsDTO{
		Broths: result,
	}, nil
}
