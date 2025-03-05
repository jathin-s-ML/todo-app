package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jathin-s-ML/todo-app/internal/models"
	"github.com/jathin-s-ML/todo-app/internal/store"
)

// Get all todos
func GetTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(store.Todos)
}

// Add a new todo
func AddTodo(w http.ResponseWriter, r *http.Request) {
	var newTodo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&newTodo); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	newTodo.ID = store.IDCounter
	store.IDCounter++
	store.Todos = append(store.Todos, newTodo)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTodo)
}

// Update a todo by ID
func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var updatedTodo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&updatedTodo); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	for i, todo := range store.Todos {
		if todo.ID == id {
			store.Todos[i].Task = updatedTodo.Task
			store.Todos[i].Completed = updatedTodo.Completed
			json.NewEncoder(w).Encode(store.Todos[i])
			return
		}
	}
	http.Error(w, "Todo not found", http.StatusNotFound)
}

// Delete a todo by ID
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for i, todo := range store.Todos {
		if todo.ID == id {
			store.Todos = append(store.Todos[:i], store.Todos[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Todo not found", http.StatusNotFound)
}
