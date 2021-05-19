package database

import (
	"database/sql"
	"errors"
	"fmt"
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

	p.db = db
	return p.db
}

func (p *products) Create(product *model.Product) error {
	if product == nil {
		return errors.New("El producto esta vacío")

	}
	stmt, err := p.db.Prepare("INSERT INTO Product(idProduct, name, price, saleDate, idCategory) VALUES (?, ?, ?, ?, ?, ? )")
	if err != nil {
		return errors.New("Ocurrio un error en el Prepare")
	}
	defer stmt.Close()

	resp, err := stmt.Exec(product.IdProduct, product.Name, product.Price, product.SaleDate, product.IdCategory)
	if err != nil {
		return errors.New("Ocurrio un error en el Exec")
	}

	id, err := resp.LastInsertId()
	if err != nil {
		return errors.New("Ocurrio al obtener el last Insert Id")
	}
	rowsAff, err := resp.RowsAffected()
	if err != nil {
		return errors.New("Ocurrio al obtener el Rows Affected")
	}
	fmt.Print(id, rowsAff)

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
