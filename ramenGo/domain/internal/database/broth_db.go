package database

import (
	"database/sql"
	"errors"

	"github.com/ramenGo/domain/internal/entity"
)

type BrothDB struct {
	DB *sql.DB
}

func NewBrothDB(db *sql.DB) *BrothDB {
	return &BrothDB{
		DB: db,
	}
}

func (b *BrothDB) Save(e entity.Entity) error {
	stmt, err := b.DB.Prepare("INSERT INTO broths (id, imageInactive, imageActive, name, description, price) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	broth, ok := e.(*entity.Broth)
	if !ok {
		return errors.New("Cant save a non Broth object. ")
	}

	_, err = stmt.Exec(broth.Id, broth.ImageInactive, broth.ImageActive, broth.Name,
		broth.Description, broth.Price)
	if err != nil {
		return err
	}

	return nil
}

func (b *BrothDB) FindById(id string) (*entity.Entity, error) {
	panic("Not Implemented Yeat")
}
