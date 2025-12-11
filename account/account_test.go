package account_test

import (
	"database/sql"
	"fmt"
	"pet-project-ToDoLIst/account"
	"testing"
)

var db *sql.DB

func TestGetID(t *testing.T) {

	name := "алексей"
	id, err := account.GetIdUser(db, name)
	if err != nil {
		t.Error("GetIdUser is failed ")
	}
	fmt.Println(id)
}
