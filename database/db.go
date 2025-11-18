package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDb() (*sql.DB, error) {

	PsqlInfo := "host=127.0.0.1 port=5432 user=postgres password=postgres dbname=postgres sslmode=disable"
	var err error
	db, err = sql.Open("postgres", PsqlInfo)
	if err != nil {
		fmt.Println("Failed to open with your data")
		return nil, err

	}
	err = db.Ping()
	if err != nil {
		fmt.Println("Failed to connect to the database")
		return nil, err
	}
	log.Println("Succesfully connected to the database")
	return db, nil

}
