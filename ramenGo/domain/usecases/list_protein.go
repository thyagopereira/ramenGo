package usecases

import (
	"github.com/ramenGo/domain/internal/database"
	"github.com/ramenGo/domain/internal/entity"
)

type ListProteinsUseCase struct {
	database database.Database
}

type ListProteinsDTO struct {
	Proteins []entity.Entity
}

func NewListProteinsUseCase(db database.Database) (*ListProteinsUseCase, error) {

	return &ListProteinsUseCase{
		database: db,
	}, nil
}

func (uc *ListProteinsUseCase) Execute() (*ListProteinsDTO, error) {
	result, err := uc.database.GetAll()
	if err != nil {
		return nil, err
	}

	return &ListProteinsDTO{
		Proteins: result,
	}, nil
}
