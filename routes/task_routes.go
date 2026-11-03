package routes

import (
	"net/http"

	"taskmanagementapi/handlers"
)

// RegisterRoutes registers the task URL patterns on the given mux.
func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/tasks", handlers.TasksHandler)
	mux.HandleFunc("/tasks/", handlers.TaskByIDHandler)
}
