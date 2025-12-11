package models

type Task struct {
	Id         int    `db:"id"`
	IdUser     int    `db:"user_id"`
	Task       string `db:"task"`
	TaskStatus bool   `db:"taskstatus"`
	Comment    string `db:"comment"`
}
type TaskResult struct {
	TaskD int
	Error error
}
type User struct {
	Id_user  int    `db:"iduser"`
	Username string `db:"username"`
	Email    string `db:"email"`
	Password string `db:"passworduser"`
}
