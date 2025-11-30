package database

import (
	"database/sql"
	"log"
	"pet-project-ToDoLIst/models"

	_ "github.com/lib/pq" // Драйвер PostgreSQL
)

var db *sql.DB

func InitDb() (*sql.DB, error) {

	PsqlInfo := "host=192.168.0.12 port=5432 user=postgres password=postgres dbname=postgres sslmode=disable"
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
	rows, err := db.Query(`SELECT id, task, taskstatus, comment FROM tasks`)

	if err != nil {
		log.Println("Can't SELECT data by your tables")
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	var tasks []models.Task
	for rows.Next() {
		var t models.Task
		err := rows.Scan(&t.Id, &t.Task, &t.TaskStatus, &t.Comment)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil

}
func AddTask(db *sql.DB, task models.Task) error {
	SqlStatement := (`INSERT INTO tasks (id, task, taskstatus, comment) VALUES ($1,$2 ,$3,$4)`)
	_, err := db.Exec(SqlStatement, task.Id, task.Task, task.TaskStatus, task.Comment)
	if err != nil {
		return err
	}
	return nil
}
func DeleteTask(db *sql.DB, Id int) error {
	SqlStatement := (`DELETE FROM tasks WHERE id = $1`)
	_, err := db.Exec(SqlStatement, Id)
	if err != nil {
		return err
	}
	return nil
}
func ChangeStatus(db *sql.DB, Id int) error {
	SqlStatement := (`UPDATE tasks SET taskstatus = true WHERE id = $1 `)
	_, err := db.Exec(SqlStatement, Id)
	if err != nil {
		return err
	}
	return nil
}
func FindTaskByName(db *sql.DB, task string) (*models.Task, error) {
	SqlStatement := (`SELECT id, task, taskstatus,comment FROM tasks WHERE task=$1`)
	var tasks1 models.Task
	err := db.QueryRow(SqlStatement, task).Scan(&tasks1.Id, &tasks1.Task, &tasks1.TaskStatus, &tasks1.Comment)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &tasks1, nil
}
