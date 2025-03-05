package main

import (
	"fmt"
	"net/http"
	"todo-app/internal/routes"
	"todo-app/internal/store"
)

func main() {
	todoStore := store.NewTodoStore()
	r := routes.InitializeRouter(todoStore)

	fmt.Println("Server running on port 8080...")
	http.ListenAndServe(":8080", r)
}
