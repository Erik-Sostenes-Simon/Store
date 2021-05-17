package database

import (
	"database/sql"
	"log"
	"os"
)

func ConnectionDB() *sql.DB {
	mysql := os.Getenv("mysql")
	root := os.Getenv("root")
	password := os.Getenv("Loindeseable09")
	table := os.Getenv("store")

	db, err := sql.Open(mysql, root+":"+password+"@tcp(127.0.0.1:3306)/"+table)

	if err != nil {
		log.Panic("Ocurrio un problema al conectarse a la base de datos")
	}
	log.Print("Conexión establecida")

	return db
}
