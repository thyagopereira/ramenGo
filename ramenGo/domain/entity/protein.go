package entity

import (
	"errors"

	"github.com/google/uuid"
)

type Protein struct {
	Id            string
	ImageInactive string
	ImageActive   string
	Name          string
	Description   string
	Price         float64
}

func NewProtein(imageInactive, imageActive, name, description string, price float64) (*Protein, error) {
	protein := &Protein{
		Id:            uuid.NewString(),
		ImageInactive: imageInactive,
		ImageActive:   imageActive,
		Name:          name,
		Description:   description,
		Price:         price,
	}

	valid, err := protein.validate()
	if valid {
		return protein, nil
	} else {
		return nil, err
	}

}

func (p *Protein) validate() (bool, error) {
	if p.Id == "" {
		return false, errors.New("Id cant be empty string")
	} else if p.ImageInactive == "" {
		return false, errors.New("ImageInactive cant be empty string")
	} else if p.ImageActive == "" {
		return false, errors.New("ImageActive cant be empty string")
	} else if p.Name == "" {
		return false, errors.New("Name cant be empty string")
	} else if p.Description == "" {
		return false, errors.New("Description cant be empty string")
	} else if p.Price == 0 {
		return false, errors.New("Price cant be zero")
	}

	return true, nil
}
