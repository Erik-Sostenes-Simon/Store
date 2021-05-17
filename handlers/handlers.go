package handlers

import (
	"net/http"
	"time"

	"github.com/Proyect/store/model"
	"github.com/labstack/echo"
)

type product struct {
	storage Storage
}

func NewProduct(storage Storage) product {
	return product{storage: storage}
}
func (p *product) Create(c echo.Context) error {
	return nil
}

func (p *product) Delete(c echo.Context) error {
	return nil
}

func (p *product) Update(c echo.Context) error {
	return nil
}
func (p *product) GetAll(c echo.Context) error {

	return c.String(http.StatusOK, "Hola mundo")
}
func (p *product) GetById(c echo.Context) error {
	p1 := &model.Product{
		IdProduct:  1,
		Name:       "Sabritas",
		Price:      10.0,
		SaleDate:   time.Now(),
		IdCategory: 2,
	}
	return c.JSON(http.StatusOK, p1)
}
