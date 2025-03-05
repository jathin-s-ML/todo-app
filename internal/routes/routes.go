package routes

import (
	"github.com/gorilla/mux"
	"github.com/jathin-s-ML/todo-app/internal/handlers"
)

// SetupRoutes initializes API routes
func SetupRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/todos", handlers.GetTodos).Methods("GET")
	r.HandleFunc("/todos", handlers.AddTodo).Methods("POST")
	r.HandleFunc("/todos/{id}", handlers.UpdateTodo).Methods("PUT")
	r.HandleFunc("/todos/{id}", handlers.DeleteTodo).Methods("DELETE")
	return r
}
