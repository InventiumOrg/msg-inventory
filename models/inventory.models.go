package models

import (
	"context"
	"database/sql"
	"log"
	"time"
)

const dbTimeout = time.Second * 5

var db *sql.DB

func New(dbPool *sql.DB) Models {
	db = dbPool
	return Models{
		Inventory: Inventory{},
	}
}

type Models struct {
	Inventory Inventory
}

type Inventory struct {
	Name      string    `json:"name"`
	Quantity  int       `json:"quantity"`
	Price     float64   `json:"price"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (i *Inventory) GetAll() ([]*Inventory, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)

	defer cancel()

	query := "SELECT * FROM inventory"

	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var inventories []*Inventory

	for rows.Next() {
		var inv Inventory

		err := rows.Scan(&inv.Name, &inv.Quantity, &inv.Price, &inv.UpdatedAt)

		if err != nil {
			log.Println("Error scanning row: ", err)
			return nil, err
		}

		inventories = append(inventories, &inv)
	}

	return inventories, nil
}

func (i *Inventory) GetByName(name string) (*Inventory, error) {
	return nil, nil
}

func (i *Inventory) Create() error {
	return nil
}

func (i *Inventory) Update() error {
	return nil
}
