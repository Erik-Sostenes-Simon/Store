package main

import (
	"github.com/Proyect/store/routes"
	"github.com/labstack/echo"
)

func main() {
	server := echo.New()

	routes.Routes(server)

	server.Logger.Fatal(server.Start(":1323"))
}
