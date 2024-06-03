package usecases

import (
	databases "github.com/ramenGo/domain/database"
	"github.com/ramenGo/domain/entity"
)

type CreateOrderUseCase struct {
	database databases.Database
}

type CreateOrderRequestDTO struct {
	BrothId   string
	ProteinId string
}

type CreateOrderResponseDTO struct {
	Id          string
	Description string
	Image       string
}

func NewCreateOrderUseCase(db databases.Database) (*CreateOrderUseCase, error) {

	return &CreateOrderUseCase{
		database: db,
	}, nil
}

func (uc *CreateOrderUseCase) Execute(o *CreateOrderRequestDTO) (*CreateOrderResponseDTO, error) {
	var order *entity.Order

	order, err := entity.NewOrder(o.BrothId, o.ProteinId, "someImage", "someDescription")
	if err != nil {
		return nil, err
	}

	err = uc.database.Save(order)
	if err != nil {
		return nil, err
	}

	return &CreateOrderResponseDTO{
		Id:          order.Id,
		Description: order.Description,
		Image:       order.Image,
	}, nil

}
