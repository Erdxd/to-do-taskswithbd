package AddTask

const (
	host     = "127.0.0.1"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

type Task struct {
	ID          int
	NAME        string
	Task_status bool
}

func Addtask() {

}
