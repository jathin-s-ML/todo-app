package todo

import (
	"fmt"
	"sort"
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

	// Collect tasks into a slice
	var tasks []*Todo
	for _, task := range tl.todos {
		tasks = append(tasks, task)
	}

	// Sort tasks by ID
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].ID < tasks[j].ID
	})

	// Print sorted tasks
	fmt.Println("List of Tasks (Sorted by ID):")
	for _, task := range tasks {
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

	// Collect tasks from channel
	var tasks []*Todo
	for task := range ch {
		tasks = append(tasks, task)
	}

	// Sort tasks by ID
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].ID < tasks[j].ID
	})

	// Print sorted tasks
	fmt.Println("Fetching tasks concurrently (Sorted by ID):")
	for _, task := range tasks {
		status := "Incomplete"
		if task.Completed {
			status = "Completed"
		}
		fmt.Printf("ID: %d | Task: %s | Status: %s\n", task.ID, task.TaskTitle, status)
	}
}
