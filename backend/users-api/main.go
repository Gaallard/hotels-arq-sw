package main

import (
	"backend/app"
	"log"

	"backend/db"
	"database/sql"

	_ "github.com/lib/pq"
)

func main() {

	db.StartDbEngine()
	app.StartRoute()

	connStr := "postgres://user:password@db:5432/mydb?sslmode=disable"
	database, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()
}
