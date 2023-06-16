package routes

import (
	"todo/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	routes := mux.NewRouter()
	routes.HandleFunc("/task", middleware.GetallTask).Methods("GET", "OPTIONS")
	routes.HandleFunc("/tasks", middleware.Createtask).Methods("POST", "OPTIONS")
	routes.HandleFunc("/task/{id}", middleware.Taskcomplete).Methods("PUT", "OPTIONS")
	routes.HandleFunc("/task/undotask/{id}", middleware.Undotask).Methods("PUT", "OPTIONS")
	routes.HandleFunc("/task/deletetask/{id}", middleware.Deletetask).Methods("DELETE", "OPTIONS")
	routes.HandleFunc("/task/deletealltask/", middleware.Deletealltask).Methods("DELETE ", "OPTIONS")
	return routes
}
