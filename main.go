package main

import (
	"github.com/Proyect/store/handlers"
	"github.com/labstack/echo"
)

func main() {
	server := echo.New()

	server.GET("/", handlers.Create)

	server.Logger.Fatal(server.Start(":1323"))
}
