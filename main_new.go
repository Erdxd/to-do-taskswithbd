package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"pet-project-ToDoLIst/account"
	"pet-project-ToDoLIst/database"
	"pet-project-ToDoLIst/models"
)

var db *sql.DB
var tmpl = template.Must(template.ParseFiles("templates/index.html"))
var tmplRegister = template.Must(template.ParseFiles("templates/register.html"))
var tmplLogin = template.Must(template.ParseFiles("templates/login.html"))

func main() {
	var err error
	db, err = database.InitDb()
	if err != nil {
		log.Fatal(err)
	}

	// Обработчики маршрутов
	http.HandleFunc("/", RegisterPageHandler)
	http.HandleFunc("/main", indexHandler)
	http.HandleFunc("/add", addTaskHandler)
	http.HandleFunc("/delete", deleteTaskHandler)
	http.HandleFunc("/changestatus", changeStatusHandler)
	http.HandleFunc("/findbyname", FindTaskByNameHandler)
	http.HandleFunc("/starttask", TimeToTaskHandler)
	http.HandleFunc("/register", RegisterPageHandler)
	http.HandleFunc("/login", LoginPageHandler)

	log.Println("Server starting on :8080...")
	err = http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("iduser")
	if err != nil {
		http.Error(w, "Не смогли извлечь куки", http.StatusBadRequest)
	}
	iduser, err := strconv.Atoi(cookie.Value)
	if err != nil {

	}

	if iduser == 0 {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}

	tasks, err := database.CheckTask(iduser)
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
		cookie, err := r.Cookie("iduser")
		if err != nil {
			http.Error(w, "Не смогли извлечь куки", http.StatusBadRequest)
		}
		iduser, err := strconv.Atoi(cookie.Value)
		if err != nil {

		}

		if iduser == 0 {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}
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

		err = database.AddTask(db, task, iduser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/main", http.StatusSeeOther)
	}
}

func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		cookie, err := r.Cookie("iduser")
		if err != nil {
			http.Error(w, "Не смогли извлечь куки", http.StatusBadRequest)
		}
		iduser, err := strconv.Atoi(cookie.Value)
		if err != nil {

		}

		if iduser == 0 {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}

		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			http.Error(w, "Неверный ID", http.StatusBadRequest)
			return
		}

		err = database.DeleteTask(db, id, iduser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/main", http.StatusSeeOther)
	}
}

func changeStatusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		cookie, err := r.Cookie("iduser")
		if err != nil {
			http.Error(w, "Не смогли извлечь куки", http.StatusBadRequest)
		}
		iduser, err := strconv.Atoi(cookie.Value)
		if err != nil {

		}

		if iduser == 0 {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}

		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			http.Error(w, "Неверный ID", http.StatusBadRequest)
			return
		}

		err = database.ChangeStatus(db, id, iduser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/main", http.StatusSeeOther)
	}
}
func FindTaskByNameHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		cookie, err := r.Cookie("iduser")
		if err != nil {
			http.Error(w, "Не смогли извлечь куки", http.StatusBadRequest)
			return
		}
		iduser, err := strconv.Atoi(cookie.Value)
		if err != nil {
			http.Error(w, "Не смогли извлечь куки", http.StatusBadRequest)
			return

		}

		if iduser == 0 {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}
		TaskFind := r.FormValue("findtask")

		data1, err := database.FindTaskByName(db, TaskFind, iduser)
		if err != nil {
			http.Error(w, "Задача не найдена", http.StatusBadRequest)
			return
		}
		tasksAll, err := database.CheckTask(iduser)
		if tasksAll == nil {
			http.Error(w, "Задача не найдена", http.StatusBadRequest)
			return
		}

		if err != nil {
			return
		}
		if data1 == nil {
			http.Error(w, "Задача не найдена", http.StatusBadRequest)
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

	http.Redirect(w, r, "/main", http.StatusSeeOther)
}
func TimeToTaskHandler(w http.ResponseWriter, r *http.Request) {
	done := make(chan models.TaskResult)
	cookie, err := r.Cookie("iduser")
	if err != nil {
		http.Error(w, "Не смогли извлечь куки", http.StatusBadRequest)
		return
	}
	iduser, err := strconv.Atoi(cookie.Value)
	if err != nil {

	}

	if iduser == 0 {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}

	Time, err := strconv.Atoi(r.FormValue("time"))
	if err != nil {
		http.Error(w, "Нет значения времени", http.StatusBadRequest)
		return
	}
	taskid, err := strconv.Atoi(r.FormValue("taskid"))

	if err != nil {
		http.Error(w, "Нет значения времени", http.StatusBadRequest)
		return
	}
	go database.TimeForTask(Time, taskid, done)
	result := <-done
	if result.Error != nil {
		http.Error(w, "Произошла ошибка", http.StatusInternalServerError)
		return
	}
	err = database.ChangeStatus(db, result.TaskD, iduser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}

	http.Redirect(w, r, "/main", http.StatusSeeOther)

}
func CreateAcoountHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		NameAccount := r.FormValue("username")
		Email := r.FormValue("Email")
		Passworduser := r.FormValue("passworduser")
		User := models.User{
			Username: NameAccount,
			Email:    Email,
			Password: Passworduser,
		}
		err := account.CreateAccount(db, User)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}
func RegisterPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		NameAccount := r.FormValue("username")
		Email := r.FormValue("Email")
		Passworduser := r.FormValue("passworduser")
		User := models.User{
			Username: NameAccount,
			Email:    Email,
			Password: Passworduser,
		}
		err := account.CreateAccount(db, User)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
	tmplRegister.Execute(w, nil)

}
func LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		NameAccount := r.FormValue("username")
		Passwordfromuser := r.FormValue("passworduser")
		passwordfromdb, err := account.LoginAccount(db, NameAccount)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if passwordfromdb != Passwordfromuser {
			http.Error(w, "password is incorrect", http.StatusInternalServerError)
			return
		} else if passwordfromdb == Passwordfromuser {
			iduser, err := account.GetIdUser(db, NameAccount)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			http.SetCookie(w, &http.Cookie{
				Name:  "iduser",
				Value: strconv.Itoa(iduser),
				Path:  "/",
			})
			http.Redirect(w, r, "/main", http.StatusSeeOther)

		}
	}
	tmplLogin.Execute(w, nil)

}
