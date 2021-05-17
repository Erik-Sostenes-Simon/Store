package routes

import (
	"github.com/Proyect/store/handlers"
	"github.com/labstack/echo"
)

func Routes(e *echo.Echo, storage handlers.Storage) {
	h := handlers.NewProduct(storage)
	route := e.Group("/v1/products/")

	route.POST("create", h.Create)
	route.DELETE("delete:id", h.Delete)
	route.PUT("update:id", h.Update)
	route.GET("get-all", h.GetAll)
	route.GET("get:id", h.GetById)
}
