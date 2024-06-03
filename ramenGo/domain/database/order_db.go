package databases

import (
	"database/sql"
	"errors"

	"github.com/ramenGo/domain/entity"
)

type OrderDB struct {
	DB *sql.DB
}

func NewOrderDB(db *sql.DB) *OrderDB {
	return &OrderDB{
		DB: db,
	}
}

func (o *OrderDB) Save(e entity.Entity) error {
	stmt, err := o.DB.Prepare("INSERT INTO orders (id, brothId, proteinId, image, description) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	order, ok := e.(*entity.Order)
	if !ok {
		return errors.New("Cant save a non Order object. ")
	}

	_, err = stmt.Exec(order.Id, order.BrothId, order.ProteinId, order.Image, order.Description)
	if err != nil {
		return err
	}

	return nil
}

func (o *OrderDB) FindById(id string) (entity.Entity, error) {
	var order entity.Order

	stmt, err := o.DB.Prepare("SELECT id, brothId, proteinId, image, description FROM orders WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	row := stmt.QueryRow(id)
	err = row.Scan(
		&order.Id,
		&order.BrothId,
		&order.ProteinId,
		&order.Image,
		&order.Description,
	)

	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (o *OrderDB) GetAll() ([]entity.Entity, error) {
	orders := make([]entity.Entity, 0)

	rows, err := o.DB.Query("SELECT id, brothId, proteinId, image, description FROM orders")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var order entity.Order
		if err = rows.Scan(
			&order.Id,
			&order.BrothId,
			&order.ProteinId,
			&order.Image,
			&order.Description,
		); err != nil {
			return nil, err
		} else {
			orders = append(orders, &order)
		}
	}

	return orders, nil
}
