package controller

import (
	"encoding/json"
	"net/http"

	models "GolangAPIProject/model"

	"github.com/gorilla/mux"
)

// CreateTodo creates a new todo
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	_ = json.NewDecoder(r.Body).Decode(&todo)

	result, err := models.InsertTodo(todo) // No need to pass *mongo.Client
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}

// GetTodos returns all todos
func GetTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := models.GetAllTodos() // No need to pass *mongo.Client
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(todos)
}

// GetTodo returns a single todo by ID
func GetTodo(w http.ResponseWriter, r *http.Request) {
	//id := r.URL.Query().Get("id")
	params := mux.Vars(r)
	id := params["id"]
	todo, err := models.GetTodoByID(id) // No need to pass *mongo.Client
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(todo)
}

// UpdateTodo updates an existing todo
func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	//id := r.URL.Query().Get("id")
	// Extract the todo ID from the URL path
	params := mux.Vars(r)
	id := params["id"]

	var todo models.Todo
	_ = json.NewDecoder(r.Body).Decode(&todo)

	err := models.UpdateTodoByID(id, todo) // No need to pass *mongo.Client
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DeleteTodo deletes a todo
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	//id := r.URL.Query().Get("id")
	params := mux.Vars(r)
	id := params["id"]

	err := models.DeleteTodoByID(id) // No need to pass *mongo.Client
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
