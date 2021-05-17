package database

import (
	"database/sql"
	"log"

	"github.com/Proyect/store/handlers"
	"github.com/Proyect/store/model"
	_ "github.com/go-sql-driver/mysql"
)

type products struct {
	db      *sql.DB
	product model.Product
}

func NewProduct() handlers.Storage {
	return &products{}
}

func (p *products) ConnectionDB() *sql.DB {
	db, err := sql.Open("mysql", "root:Loindeseable09@tcp(127.0.0.1:3306)/store")

	if err != nil {
		log.Panic("Ocurrio un problema al conectarse a la base de datos")
	}
	log.Print("Conexión establecida")

	return db
}

func (p *products) Create(product *model.Product) error {
	return nil
}
func (p *products) Delete(id int) error {
	return nil
}
func (p *products) Update(id int, product *model.Product) error {
	return nil
}
func (p *products) GetById(id int) (model.Product, error) {
	return model.Product{}, nil
}
func (p *products) GetAll() (model.Products, error) {
	var sliceProducts model.Products
	return sliceProducts, nil
}
