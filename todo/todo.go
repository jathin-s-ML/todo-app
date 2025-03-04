
package todo

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func RunTodoApp() {
	fmt.Println("TODO Application")
	reader := bufio.NewReader(os.Stdin)
	todoManager := NewTodoList()

	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1. Add Task")
		fmt.Println("2. List All Tasks")
		fmt.Println("3. Mark a Task as Completed")
		fmt.Println("4. Delete a Task")
		fmt.Println("5. Fetch Tasks Concurrently")
		fmt.Println("6. Exit")

		fmt.Print("Enter your choice: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		switch choice {
		case 1:
			fmt.Print("Enter the task you want to add: ")
			title, _ := reader.ReadString('\n')
			todoManager.Add(title)

		case 2:
			todoManager.List()

		case 3:
			fmt.Print("Enter the Task ID to mark as complete: ")
			idStr, _ := reader.ReadString('\n')
			id, err := strconv.Atoi(strings.TrimSpace(idStr))
			if err != nil {
				fmt.Println("Invalid ID. Please enter a number.")
				continue
			}
			if err := todoManager.MarkAsCompleted(id); err != nil {
				fmt.Println(err)
			}

		case 4:
			fmt.Print("Enter the Task ID to delete: ")
			idStr, _ := reader.ReadString('\n')
			id, err := strconv.Atoi(strings.TrimSpace(idStr))
			if err != nil {
				fmt.Println("Invalid ID. Please enter a number.")
				continue
			}
			if err := todoManager.DeleteTask(id); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Task deleted successfully")
			}

		case 5:
			todoManager.FetchTasksConcurrently()

		case 6:
			fmt.Println("Exiting TODO application")
			return

		default:
			fmt.Println("Invalid choice. Please enter a valid option")
		}
	}
}

