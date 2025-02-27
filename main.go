package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type todo struct {
	id        int
	tasktitle string
	completed bool
}

type todolist struct {
	todos  map[int]*todo
	nextid int
}

type todomanager interface {
	add(title string)
	deletetask(id int) error
	markascompleted(id int) error
	list()
}

func newtodolist() *todolist {
	return &todolist{todos: make(map[int]*todo), nextid: 1}
}

func (tl *todolist) add(title string) {
	title = strings.TrimSpace(title)
	if title == "" {
		fmt.Println("Task not mentioned")
		return
	}
	tl.todos[tl.nextid] = &todo{id: tl.nextid, tasktitle: title, completed: false}
	fmt.Printf("Task added: %s (ID: %d)\n", title, tl.nextid)
	tl.nextid++
}

func (tl *todolist) deletetask(id int) error {
	if _, exists := tl.todos[id]; exists {
		delete(tl.todos, id)
		return nil
	}
	return fmt.Errorf("Task not found (ID: %d)", id)
}

func (tl *todolist) markascompleted(id int) error {
	if task, exists := tl.todos[id]; exists {
		task.completed = true
		fmt.Println("Task marked as completed")
		return nil
	}
	return fmt.Errorf("Invalid task ID: %d", id)
}

func (tl *todolist) list() {
	if len(tl.todos) == 0 {
		fmt.Println("No tasks to display")
		return
	}
	fmt.Println("List of Tasks:")
	for _, task := range tl.todos {
		status := "Incomplete"
		if task.completed {
			status = "Completed"
		}
		fmt.Printf("ID: %d | Task: %s | Status: %s\n", task.id, task.tasktitle, status)
	}
}

func main() {
	fmt.Println("TODO Application")
	reader := bufio.NewReader(os.Stdin)
	todomanager := newtodolist()

	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1. Add Task")
		fmt.Println("2. List All Tasks")
		fmt.Println("3. Mark a Task as Completed")
		fmt.Println("4. Delete a Task")
		fmt.Println("5. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Print("Enter the task you want to add: ")
			title, _ := reader.ReadString('\n')
			todomanager.add(title)

		case 2:
			todomanager.list()

		case 3:
			fmt.Print("Enter the Task ID to mark as complete: ")
			var id int
			fmt.Scan(&id)
			if err := todomanager.markascompleted(id); err != nil {
				fmt.Println(err)
			}

		case 4:
			fmt.Print("Enter the Task ID to delete: ")
			var id int
			fmt.Scan(&id)
			fmt.Scanln()
			if err := todomanager.deletetask(id); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Task deleted successfully")
			}

		case 5:
			fmt.Println("Exiting TODO application")
			return

		default:
			fmt.Println("Invalid choice. Please enter a valid option")
		}
	}
}
