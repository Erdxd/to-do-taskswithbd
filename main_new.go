package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/Erdxd/conv-IMT-bookmarks-massivestatic.git/pet-project-ToDoLIst/database"
	"github.com/Erdxd/conv-IMT-bookmarks-massivestatic.git/pet-project-ToDoLIst/models"
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

	// Запуск веб-сервера
	fmt.Println("Сервер запущен на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// indexHandler отображает главную страницу со списком задач
func indexHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := database.CheckTask()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Tasks []models.Task
	}{
		Tasks: tasks,
	}

	tmpl.Execute(w, data)
}

// addTaskHandler обрабатывает добавление новой задачи
func addTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Получаем данные из формы
		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			http.Error(w, "Неверный ID", http.StatusBadRequest)
			return
		}

		taskName := r.FormValue("task")
		taskStatus := r.FormValue("status") == "on" // чекбокс возвращает "on" если отмечен

		// Создаем новую задачу
		task := models.Task{
			Id:         id,
			Task:       taskName,
			TaskStatus: taskStatus,
		}

		// Добавляем задачу в базу данных
		err = database.AddTask(db, task)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Перенаправляем на главную страницу
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

// deleteTaskHandler обрабатывает удаление задачи
func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Получаем ID задачи из формы
		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			http.Error(w, "Неверный ID", http.StatusBadRequest)
			return
		}

		// Удаляем задачу из базы данных
		err = database.DeleteTask(db, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Перенаправляем на главную страницу
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

// changeStatusHandler обрабатывает изменение статуса задачи
func changeStatusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Получаем ID задачи из формы
		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			http.Error(w, "Неверный ID", http.StatusBadRequest)
			return
		}

		// Изменяем статус задачи в базе данных
		err = database.ChangeStatus(db, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Перенаправляем на главную страницу
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
