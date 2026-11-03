package main

import (
	"log"
	"net/http"

	"taskmanagementapi/middleware"
	"taskmanagementapi/routes"
)

func main() {
	mux := http.NewServeMux()
	routes.RegisterRoutes(mux)

	log.Println("Task Management API running at http://localhost:8080")
	if err := http.ListenAndServe(":8080", middleware.Logging(mux)); err != nil {
		log.Fatal(err)
	}
}
