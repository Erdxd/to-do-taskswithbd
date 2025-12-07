package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"pet-project-ToDoLIst/database"
	"pet-project-ToDoLIst/models"
)

var db *sql.DB
var tmpl = template.Must(template.ParseFiles("templates/index.html"))

func main() {
	var err error
	db, err = database.InitDb()
	if err != nil {
		log.Fatal(err)
	}

	// Обработчики маршрутов
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/add", addTaskHandler)
	http.HandleFunc("/delete", deleteTaskHandler)
	http.HandleFunc("/changestatus", changeStatusHandler)
	http.HandleFunc("/findbyname", FindTaskByNameHandler)
	http.HandleFunc("/starttask", TimeToTask)

	// Запуск веб-сервера
	log.Println("Server starting on :8080...")
	err = http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := database.CheckTask()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		TasksAll   []models.Task
		SearchTask []models.Task
	}{
		TasksAll:   tasks,
		SearchTask: nil,
	}

	tmpl.Execute(w, data)
}

func addTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			http.Error(w, "Неверный ID", http.StatusBadRequest)
			return
		}

		taskName := r.FormValue("task")
		taskStatus := r.FormValue("status") == "on"
		taskcomment := r.FormValue("comment")

		task := models.Task{
			Id:         id,
			Task:       taskName,
			TaskStatus: taskStatus,
			Comment:    taskcomment,
		}

		err = database.AddTask(db, task)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			http.Error(w, "Неверный ID", http.StatusBadRequest)
			return
		}

		err = database.DeleteTask(db, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func changeStatusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			http.Error(w, "Неверный ID", http.StatusBadRequest)
			return
		}

		err = database.ChangeStatus(db, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
func FindTaskByNameHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		TaskFind := r.FormValue("findtask")

		data1, err := database.FindTaskByName(db, TaskFind)
		if err != nil {
			http.Error(w, "Задача не найдена", http.StatusBadRequest)
			return
		}
		tasksAll, err := database.CheckTask()

		if err != nil {
			return
		}

		SliceTask := []models.Task{*data1}
		data := struct {
			TasksAll   []models.Task
			SearchTask []models.Task
		}{SearchTask: SliceTask,
			TasksAll: tasksAll}
		tmpl.Execute(w, data)

	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
func TimeToTask(w http.ResponseWriter, r *http.Request) {
	done := make(chan models.TaskResult)

	Time, err := strconv.Atoi(r.FormValue("time"))
	taskid, err := strconv.Atoi(r.FormValue("taskid"))

	if err != nil {
		http.Error(w, "Нет значения времени", http.StatusBadRequest)
		return
	}
	go database.TimeForTask(Time, taskid, done)
	result := <-done
	if result.Error != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = database.ChangeStatus(db, result.TaskD)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}

	http.Redirect(w, r, "/", http.StatusSeeOther)

}
