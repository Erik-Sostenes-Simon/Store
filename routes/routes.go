package routes

import (
	"github.com/Proyect/store/handlers"
	"github.com/labstack/echo"
)

func Routes(e *echo.Echo) {
	route := e.Group("/v1/products")

	route.GET("/create", handlers.Create)
}
