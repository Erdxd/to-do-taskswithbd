package main

import (
	"database/sql"
	"fmt"
	"log"

	functionwithdb "github.com/Erdxd/conv-IMT-bookmarks-massivestatic.git/pet-project-ToDoLIst/Functionwithdb"
	"github.com/Erdxd/conv-IMT-bookmarks-massivestatic.git/pet-project-ToDoLIst/database"
	"github.com/Erdxd/conv-IMT-bookmarks-massivestatic.git/pet-project-ToDoLIst/models"
)

func main1() {
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
		fmt.Println("5-Delete Task")
		fmt.Println("6-Stats")
		fmt.Println("7-Exit")
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
			ChangeStatus(db)
		case "5":
			DeleteTask(db)
		case "6":

		case "7":
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
func DeleteTask(db *sql.DB) {
	var IdForDelete int
	fmt.Println("Enter Id")
	fmt.Scan(&IdForDelete)

	err := database.DeleteTask(db, IdForDelete)
	if err != nil {
		log.Fatal(err)
	}
}
func ChangeStatus(db *sql.DB) {
	var IdForChange int
	fmt.Println("Enter Id")

	fmt.Scan(&IdForChange)

	err := database.ChangeStatus(db, IdForChange)
	if err != nil {
		log.Fatal(err)
	}
}
