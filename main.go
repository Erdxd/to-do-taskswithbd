package main

import (
	"fmt"
	"log"

	"github.com/Erdxd/conv-IMT-bookmarks-massivestatic.git/pet-project-ToDoLIst/database"
)

func main() {
	_, err := database.InitDb()
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

		case "2":

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
