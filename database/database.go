package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectionDB() *sql.DB {
	db, err := sql.Open("mysql", "root:Loindeseable09@tcp(127.0.0.1:3306)/store")

	if err != nil {
		log.Panic("Ocurrio un problema al conectarse a la base de datos")
	}
	log.Print("Conexión establecida")

	return db
}
