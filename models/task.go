package models

type Task struct {
	Id         int    `db:"id"`
	Task       string `db:"task"`
	TaskStatus bool   `db:"taskstatus"`
}
