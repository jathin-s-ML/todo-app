package todo

type TodoManager interface {
	Add(title string)
	DeleteTask(id int) error
	MarkAsCompleted(id int) error
	List()
}
