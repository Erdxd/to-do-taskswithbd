package database

import (
	"database/sql"

	"log"

	"github.com/Erdxd/conv-IMT-bookmarks-massivestatic.git/pet-project-ToDoLIst/models"
	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDb() (*sql.DB, error) {

	PsqlInfo := "host=127.0.0.1 port=5432 user=postgres password=postgres dbname=postgres sslmode=disable"
	var err error
	db, err = sql.Open("postgres", PsqlInfo)
	if err != nil {
		log.Println("Failed to open with your data")
		return nil, err

	}
	err = db.Ping()
	if err != nil {
		log.Println("Failed to connect to the database")
		return nil, err
	}

	log.Println("Succesfully connected to the database")
	return db, nil

}
func CheckTask() ([]models.Task, error) {
	rows, err := db.Query(`SELECT id, task, taskstatus FROM public.tasks`)

	if err != nil {
		log.Println("Can't SELECT data by your tables")
		return nil, err
	}
	defer rows.Close()
	var tasks []models.Task
	for rows.Next() {
		var t models.Task
		err := rows.Scan(&t.Id, &t.Task, &t.TaskStatus)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil

}
func AddTask(db *sql.DB, task models.Task) error {
	SqlStatement := (`INSERT INTO tasks (id, task, taskstatus) VALUES ($1,$2 ,$3)`)
	_, err := db.Exec(SqlStatement, task.Id, task.Task, task.TaskStatus)
	if err != nil {
		return err
	}
	return nil
}
