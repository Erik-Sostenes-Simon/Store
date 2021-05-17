package main

import (
	"log"

	"github.com/Proyect/store/database"
	"github.com/Proyect/store/routes"
	"github.com/labstack/echo"
)

func main() {
	server := echo.New()

	p := database.NewProduct()
	routes.Routes(server, p)

	db := p.ConnectionDB()

	if err := db.Ping(); err != nil {
		panic(err.Error())
	}

	log.Print("Servidor iniciado")
	server.Logger.Fatal(server.Start(":1323"))
}
