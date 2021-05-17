package handlers

import (
	"database/sql"

	"github.com/Proyect/store/model"
)

type Storage interface {
	Create(product *model.Product) error
	Delete(id int) error
	Update(id int, product *model.Product) error
	GetById(id int) (model.Product, error)
	GetAll() (model.Products, error)
	ConnectionDB() *sql.DB
}
