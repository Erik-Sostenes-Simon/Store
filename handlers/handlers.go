package handlers

import (
	"net/http"
	"time"

	"github.com/Proyect/store/model"
	"github.com/labstack/echo"
)

func Create(c echo.Context) error {
	return nil
}

func Delete(c echo.Context) error {
	return nil
}

func Update(c echo.Context) error {
	return nil
}
func GetAll(c echo.Context) error {
	return nil
}
func GetById(c echo.Context) error {
	p := &model.Product{
		IdProduct:  1,
		Name:       "Sabritas",
		Price:      10.0,
		SaleDate:   time.Now(),
		IdCategory: 2,
	}
	return c.JSON(http.StatusOK, p)
}
