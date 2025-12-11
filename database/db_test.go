package database_test

import (
	"pet-project-ToDoLIst/database"
	"pet-project-ToDoLIst/models"
	"testing"
)

func TestTimeForTask(t *testing.T) {
	data := make(chan models.TaskResult)

	n := 4
	database.TimeForTask(n, 1, data)
	result := <-data
	if result.Error != nil {
		return
	}

}
