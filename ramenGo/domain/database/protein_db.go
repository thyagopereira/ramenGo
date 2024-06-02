package databases

import (
	"database/sql"
	"errors"

	"github.com/ramenGo/domain/entity"
)

type ProteinDB struct {
	DB *sql.DB
}

func NewProteinDB(db *sql.DB) *ProteinDB {
	return &ProteinDB{
		DB: db,
	}
}

func (b *ProteinDB) Save(e entity.Entity) error {
	stmt, err := b.DB.Prepare("INSERT INTO proteins (id, imageInactive, imageActive, name, description, price) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	protein, ok := e.(*entity.Protein)
	if !ok {
		return errors.New("Cant save a non Protein object. ")
	}

	_, err = stmt.Exec(protein.Id, protein.ImageInactive, protein.ImageActive, protein.Name,
		protein.Description, protein.Price)
	if err != nil {
		return err
	}

	return nil
}

func (b *ProteinDB) FindById(id string) (entity.Entity, error) {
	var protein entity.Protein

	stmt, err := b.DB.Prepare("SELECT id, imageInactive, imageActive, name, description, price FROM proteins WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	row := stmt.QueryRow(id)
	err = row.Scan(
		&protein.Id,
		&protein.ImageInactive,
		&protein.ImageActive,
		&protein.Name,
		&protein.Description,
		&protein.Price,
	)

	if err != nil {
		return nil, err
	}

	return &protein, nil
}

func (b *ProteinDB) GetAll() ([]entity.Entity, error) {
	proteins := make([]entity.Entity, 0)

	rows, err := b.DB.Query("SELECT id, imageInactive, imageActive, name, description, price FROM proteins")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var protein entity.Protein
		if err = rows.Scan(
			&protein.Id,
			&protein.ImageInactive,
			&protein.ImageActive,
			&protein.Name,
			&protein.Description,
			&protein.Price,
		); err != nil {
			return nil, err
		} else {
			proteins = append(proteins, &protein)
		}
	}

	return proteins, nil
}
