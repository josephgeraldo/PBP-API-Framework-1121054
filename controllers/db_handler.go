package controllers

import (
	"database/sql"
	"log"
)

//Connect to Database
func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/tugasecho")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
