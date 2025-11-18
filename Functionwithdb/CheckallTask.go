package functionwithdb

import (
	"log"

	"github.com/Erdxd/conv-IMT-bookmarks-massivestatic.git/pet-project-ToDoLIst/database"
)

func GetAllTasks() {
	tasks, err := database.CheckTask()
	if err != nil {
		log.Fatal(err)

	}
	log.Printf("All tasks: %+v\n", tasks)

}
