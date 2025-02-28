package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Todo struct {
	ID        int
	Title     string
	Completed bool
}

var todos = make(map[int]Todo)
var nextID = 1

func AddTodo(title string) {
	title = strings.TrimSpace(title)
	if title == "" {
		fmt.Println("Error: Task title cannot be empty!")
		return
	}
	todos[nextID] = Todo{ID: nextID, Title: title, Completed: false}
	fmt.Printf("Task added: %s (ID: %d)\n", title, nextID)
	nextID++
}

func ListTodos() {
	if len(todos) == 0 {
		fmt.Println("There are no tasks to display.")
		return
	}
	fmt.Println("\n--- TODO LIST ---")
	for _, todo := range todos {
		status := "Pending"
		if todo.Completed {
			status = "Completed"
		}
		fmt.Printf("ID: %d | Task: %s | Status: %s\n", todo.ID, todo.Title, status)
	}
}

func MarkAsCompleted(taskID int) {
	if task, exists := todos[taskID]; exists {
		task.Completed = true
		todos[taskID] = task
		fmt.Printf("Task (ID: %d) marked as completed.\n", taskID)
	} else {
		fmt.Println("Error: Task ID not found!")
	}
}

func DeleteTodo(taskID int) {
	if _, exists := todos[taskID]; exists {
		delete(todos, taskID)
		fmt.Printf("Task (ID: %d) deleted successfully.\n", taskID)
	} else {
		fmt.Println("Error: Task ID not found!")
	}
}

func main() {
	fmt.Println("Welcome to the TODO Application!")
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1. Add Task")
		fmt.Println("2. List All Tasks")
		fmt.Println("3. Mark a Task as Completed")
		fmt.Println("4. Delete a Task")
		fmt.Println("5. Exit")

		fmt.Print("Enter your choice: ")
		var choice int
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Invalid input! Please enter a number.")
			continue
		}

		switch choice {
		case 1:
			fmt.Print("Enter the task title: ")
			taskTitle, _ := reader.ReadString('\n')
			AddTodo(taskTitle)

		case 2:
			ListTodos()

		case 3:
			fmt.Print("Enter the ID of the task to mark as completed: ")
			var id int
			fmt.Scan(&id)
			MarkAsCompleted(id)

		case 4:
			fmt.Print("Enter the ID of the task to delete: ")
			var id int
			fmt.Scan(&id)
			DeleteTodo(id)

		case 5:
			fmt.Println("Exiting the TODO application.")
			os.Exit(0)

		default:
			fmt.Println("Invalid choice! Please select a valid option.")
		}
	}
}
