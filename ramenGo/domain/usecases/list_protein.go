package usecases

import (
	databases "github.com/ramenGo/domain/database"
	"github.com/ramenGo/domain/entity"
)

type ListProteinsUseCase struct {
	database databases.Database
}

type ListProteinsDTO struct {
	Proteins []entity.Entity
}

func NewListProteinsUseCase(db databases.Database) (*ListProteinsUseCase, error) {

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
