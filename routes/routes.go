package routes

import (
	"github.com/Proyect/store/handlers"
	"github.com/labstack/echo"
)

func Routes(e *echo.Echo) {
	route := e.Group("/v1/products/")

	route.POST("create", handlers.Create)
	route.DELETE("delete:id", handlers.Delete)
	route.PUT("update: id", handlers.Update)
	route.GET("get-all", handlers.GetAll)
	route.GET("get:id", handlers.GetById)
}
