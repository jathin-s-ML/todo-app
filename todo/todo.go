package todo

import (
	"fmt"
	"strings"
	"sync"
)

type TodoList struct {
	todos  map[int]*Todo
	nextID int
	sync.Mutex
}

// NewTodoList creates a new TodoList instance
func NewTodoList() *TodoList {
	return &TodoList{todos: make(map[int]*Todo), nextID: 1}
}

// Add a new task to the todo list
func (tl *TodoList) Add(title string) {
	title = strings.TrimSpace(title)
	if title == "" {
		fmt.Println("Task not mentioned")
		return
	}
	tl.Lock()
	defer tl.Unlock()

	tl.todos[tl.nextID] = &Todo{ID: tl.nextID, TaskTitle: title, Completed: false}
	fmt.Printf("Task added: %s (ID: %d)\n", title, tl.nextID)
	tl.nextID++
}

// DeleteTask removes a task from the list
func (tl *TodoList) DeleteTask(id int) error {
	tl.Lock()
	defer tl.Unlock()

	if _, exists := tl.todos[id]; exists {
		delete(tl.todos, id)
		return nil
	}
	return fmt.Errorf("Task not found (ID: %d)", id)
}

// MarkAsCompleted marks a task as completed
func (tl *TodoList) MarkAsCompleted(id int) error {
	tl.Lock()
	defer tl.Unlock()

	if task, exists := tl.todos[id]; exists {
		task.Completed = true
		fmt.Println("Task marked as completed")
		return nil
	}
	return fmt.Errorf("Invalid task ID: %d", id)
}

// List all tasks
func (tl *TodoList) List() {
	tl.Lock()
	defer tl.Unlock()

	if len(tl.todos) == 0 {
		fmt.Println("No tasks to display")
		return
	}
	fmt.Println("List of Tasks:")
	for _, task := range tl.todos {
		status := "Incomplete"
		if task.Completed {
			status = "Completed"
		}
		fmt.Printf("ID: %d | Task: %s | Status: %s\n", task.ID, task.TaskTitle, status)
	}
}
