package entity

import (
	"errors"

	"github.com/google/uuid"
)

type Broth struct {
	Id            string
	ImageInactive string
	ImageActive   string
	Name          string
	Description   string
	Price         float64
}

func NewBroth(imageInactive, imageActive, name, description string, price float64) (*Broth, error) {
	broth := &Broth{
		Id:            uuid.NewString(),
		ImageInactive: imageInactive,
		ImageActive:   imageActive,
		Name:          name,
		Description:   description,
		Price:         price,
	}

	valid, err := broth.validate()
	if valid {
		return broth, nil
	} else {
		return nil, err
	}
}

func (b *Broth) validate() (bool, error) {
	if b.Id == "" {
		return false, errors.New("Id cant be empty string")
	} else if b.ImageInactive == "" {
		return false, errors.New("ImageInactive cant be empty string")
	} else if b.ImageActive == "" {
		return false, errors.New("ImageActive cant be empty string")
	} else if b.Name == "" {
		return false, errors.New("Name cant be empty string")
	} else if b.Description == "" {
		return false, errors.New("Description cant be empty string")
	} else if b.Price == 0 {
		return false, errors.New("Price cant be zero")
	}

	return true, nil
}
