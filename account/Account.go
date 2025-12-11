package account

import (
	"database/sql"
	"pet-project-ToDoLIst/models"
)

func CreateAccount(db *sql.DB, user models.User) error {
	SqlStatement := (`INSERT INTO "user" (username,email,passworduser) VALUES ($1,$2,$3)`)
	_, err := db.Exec(SqlStatement, user.Username, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}
func LoginAccount(db *sql.DB, username string) (string, error) {
	var password string
	err := db.QueryRow(`SELECT passworduser FROM "user" WHERE username = $1`, username).Scan(&password)
	if err != nil {
		return "", err
	}

	return password, err
}
func GetIdUser(db *sql.DB, username string) (int, error) {
	var Id int
	err := db.QueryRow(`SELECT iduser FROM "user" WHERE username = $1`, username).Scan(&Id)
	if err != nil {
		return 0, err
	}
	return Id, err
}
