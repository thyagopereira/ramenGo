package entity

import (
	"errors"

	"github.com/google/uuid"
)

type Order struct {
	Id          string
	BrothId     string
	ProteinId   string
	Image       string
	Description string
}

func NewOrder(brothId, proteinId, image, description string) (*Order, error) {
	order := &Order{
		Id:          uuid.NewString(),
		BrothId:     brothId,
		ProteinId:   proteinId,
		Image:       image,
		Description: description,
	}

	valid, err := order.validate()
	if valid {
		return order, nil
	} else {
		return nil, err
	}
}

func (o *Order) validate() (bool, error) {
	if o.Id == "" {
		return false, errors.New("Id cant be empty string")
	} else if o.BrothId == "" {
		return false, errors.New("BrothId cant be empty string")
	} else if o.ProteinId == "" {
		return false, errors.New("ProteinId cant be empty string")
	} else if o.Image == "" {
		return false, errors.New("Image cant be empty string")
	} else if o.Description == "" {
		return false, errors.New("Description cant be empty string")
	}

	return true, nil
}
