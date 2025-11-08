package AddTask

import "fmt"

var Task = []string{}

func Addtask() {

	fmt.Println("Write name your new task")
	var tasknew string
	fmt.Scan(&tasknew)
	if tasknew == "" {
		fmt.Println("Ops,You dont write new task, wrIte again please")
	} else if tasknew != "" {
		Task = append(Task, tasknew)
	}

}
func CheckAllTask() {
	fmt.Println(Task)
}
