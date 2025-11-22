package main

import (
	"database/sql"
	"fmt"
	"log"

	functionwithdb "github.com/Erdxd/conv-IMT-bookmarks-massivestatic.git/pet-project-ToDoLIst/Functionwithdb"
	"github.com/Erdxd/conv-IMT-bookmarks-massivestatic.git/pet-project-ToDoLIst/database"
	"github.com/Erdxd/conv-IMT-bookmarks-massivestatic.git/pet-project-ToDoLIst/models"
)

func main() {
	db, err := database.InitDb()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("___TO-DO LIST___")

menu:
	for {

		fmt.Println("1-Add Task")
		fmt.Println("2-Check Task")
		fmt.Println("3-Start Task")
		fmt.Println("4-Done Task")
		fmt.Println("5-Stats")
		fmt.Println("6-Exit")
		Useranswer := ""
		fmt.Scan(&Useranswer)
		switch Useranswer {
		case "1":
			AddTask(db)

		case "2":
			functionwithdb.GetAllTasks()
		case "3":
			fmt.Println("Which task you want to start?")
			fmt.Println()
		case "4":
			fmt.Println("Which task you want to done?")
			fmt.Println()
		case "5":

		case "6":
			break menu
		}
	}

}
func AddTask(db *sql.DB) {
	fmt.Println("Enter task id")
	var idw int
	fmt.Scan(&idw)
	fmt.Println("Enter task name")
	var taskw string
	fmt.Scan(&taskw)
	fmt.Println("Enter task status")
	var taskStatusw bool
	fmt.Scan(&taskStatusw)
	task := models.Task{Id: idw, Task: taskw, TaskStatus: taskStatusw}
	err := database.AddTask(db, task)
	if err != nil {
		log.Fatal()

	}

}
