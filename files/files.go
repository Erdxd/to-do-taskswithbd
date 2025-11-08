package files

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Task struct {
	Task []Task
}

type Account struct {
	TimeOnTask time.Time `json:"TimeSpend" `
	Done       bool      `json:"Done" `
	ID         int       `json:"ID" `
	Title      string    `json:"Title" `
}

func SaveTask(tasks []Task) {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		fmt.Println("Ошибка при кодировании в JSON:", err)
		return
	}

	err = os.WriteFile("Task.json", data, 0644)
	if err != nil {
		fmt.Println("Ошибка при записи файла:", err)
		return
	}
}
func LoadTask() []Task {
	var tasks []Task
	data, err := os.ReadFile("Task.json")
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}
		}
		fmt.Println("Ошибка при чтении файла:", err)
		return []Task{}
	}
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		fmt.Println("Ошибка при парсинге в JSON:", err)
		return []Task{}
	}

	fmt.Println(string(data))
	return tasks
}
