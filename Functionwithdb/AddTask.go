package functionwithdb

import (
	"database/sql"
	"log"

	"github.com/Erdxd/conv-IMT-bookmarks-massivestatic.git/pet-project-ToDoLIst/database"
	"github.com/Erdxd/conv-IMT-bookmarks-massivestatic.git/pet-project-ToDoLIst/models"
)

func AddTask(db *sql.DB) {

	task := models.Task{Id: 2, Task: "Go to the gym", TaskStatus: false}
	err := database.AddTask(db, task)
	if err != nil {
		log.Fatal()

	}

}
