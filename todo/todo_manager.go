
package todo

import (
	"fmt"
	"strings"
	"sync"
)

type TodoManager interface {
	Add(title string)
	DeleteTask(id int) error
	MarkAsCompleted(id int) error
	List()
	FetchTasksConcurrently()
}

type TodoList struct {
	todos  map[int]*Todo
	nextID int
	sync.Mutex
}

func NewTodoList() *TodoList {
	return &TodoList{todos: make(map[int]*Todo), nextID: 1}
}

func (tl *TodoList) Add(title string) {
	title = strings.TrimSpace(title)
	if title == "" {
		fmt.Println("Task not mentioned")
		return
	}
	tl.Lock()
	tl.todos[tl.nextID] = &Todo{ID: tl.nextID, TaskTitle: title, Completed: false}
	fmt.Printf("Task added: %s (ID: %d)\n", title, tl.nextID)
	tl.nextID++
	tl.Unlock()
}

func (tl *TodoList) DeleteTask(id int) error {
	tl.Lock()
	defer tl.Unlock()
	if _, exists := tl.todos[id]; exists {
		delete(tl.todos, id)
		return nil
	}
	return fmt.Errorf("task not found (ID: %d)", id)
}

func (tl *TodoList) MarkAsCompleted(id int) error {
	tl.Lock()
	defer tl.Unlock()
	if task, exists := tl.todos[id]; exists {
		task.Completed = true
		fmt.Println("Task marked as completed")
		return nil
	}
	return fmt.Errorf("invalid task ID: %d", id)
}

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

func (tl *TodoList) FetchTasksConcurrently() {
	var wg sync.WaitGroup
	ch := make(chan *Todo, len(tl.todos))

	tl.Lock()
	for _, task := range tl.todos {
		wg.Add(1)
		go func(t *Todo) {
			defer wg.Done()
			ch <- t
		}(task)
	}
	tl.Unlock()

	go func() {
		wg.Wait()
		close(ch)
	}()

	fmt.Println("Fetching tasks concurrently:")
	for task := range ch {
		status := "Incomplete"
		if task.Completed {
			status = "Completed"
		}
		fmt.Printf("ID: %d | Task: %s | Status: %s\n", task.ID, task.TaskTitle, status)
	}
}
