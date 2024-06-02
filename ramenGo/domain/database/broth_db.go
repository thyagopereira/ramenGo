package databases

import (
	"database/sql"
	"errors"

	"github.com/ramenGo/domain/entity"
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

func (b *BrothDB) FindById(id string) (entity.Entity, error) {
	var broth entity.Broth

	stmt, err := b.DB.Prepare("SELECT id, imageInactive, imageActive, name, description, price FROM broths WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	row := stmt.QueryRow(id)
	err = row.Scan(
		&broth.Id,
		&broth.ImageInactive,
		&broth.ImageActive,
		&broth.Name,
		&broth.Description,
		&broth.Price,
	)

	if err != nil {
		return nil, err
	}

	return &broth, nil
}

func (b *BrothDB) GetAll() ([]entity.Entity, error) {
	broths := make([]entity.Entity, 0)

	rows, err := b.DB.Query("SELECT id, imageInactive, imageActive, name, description, price FROM broths")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var broth entity.Broth
		if err = rows.Scan(
			&broth.Id,
			&broth.ImageInactive,
			&broth.ImageActive,
			&broth.Name,
			&broth.Description,
			&broth.Price,
		); err != nil {
			return nil, err
		} else {
			broths = append(broths, &broth)
		}
	}

	return broths, nil
}
