package account_test

import (
	"database/sql"
	"fmt"
	"log"
	"pet-project-ToDoLIst/account"
	"testing"

	_ "github.com/lib/pq"
)

func TestGetID(t *testing.T) {
	PsqlInfo := "host=192.168.0.12 port=5432 user=postgres password=postgres dbname=postgres sslmode=disable"

	db, err := sql.Open("postgres", PsqlInfo)
	if err != nil {
		log.Println("Failed to open with your data")
		return

	}
	err = db.Ping()
	if err != nil {
		log.Println("Failed to connect to the database")

	}

	log.Println("Succesfully connected to the database")

	name := "алексей"
	id, _ := account.GetIdUser(db, name)
	fmt.Println(id)

}
