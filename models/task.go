package models

type Task struct {
	Id         int    `db:"id"`
	Task       string `db:"task"`
	TaskStatus bool   `db:"taskstatus"`
	Comment    string `db:"comment"`
}
type TaskResult struct {
	TaskD int
	Error error
}
