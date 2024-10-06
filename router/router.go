package router

import (
	"GolangAPIProject/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	// Define routes for your API
	router.HandleFunc("/api/todos", controller.CreateTodo).Methods("POST")
	router.HandleFunc("/api/todos", controller.GetTodos).Methods("GET")
	router.HandleFunc("/api/todo/{id}", controller.GetTodo).Methods("GET")
	router.HandleFunc("/api/todo/{id}", controller.UpdateTodo).Methods("PUT")
	router.HandleFunc("/api/todo/{id}", controller.DeleteTodo).Methods("DELETE")

	return router
}
