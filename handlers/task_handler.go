package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"taskmanagementapi/models"
)

var (
	mu     sync.Mutex
	tasks  []models.Task
	nextID = 1
)

// writeJSON sends a JSON response with the given status code.
func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// methodNotAllowed returns HTTP 405 and lists the allowed methods.
func methodNotAllowed(w http.ResponseWriter, allowed string) {
	w.Header().Set("Allow", allowed)
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

// TasksHandler handles /tasks (POST and GET).
func TasksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		createTask(w, r)
	case http.MethodGet:
		listTasks(w)
	default:
		methodNotAllowed(w, "GET, POST")
	}
}

// TaskByIDHandler handles /tasks/{id} (GET, PUT and DELETE).
func TaskByIDHandler(w http.ResponseWriter, r *http.Request) {
	idText := strings.TrimPrefix(r.URL.Path, "/tasks/")
	id, err := strconv.Atoi(idText)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		getTask(w, id)
	case http.MethodPut:
		updateTask(w, r, id)
	case http.MethodDelete:
		deleteTask(w, id)
	default:
		methodNotAllowed(w, "GET, PUT, DELETE")
	}
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var t models.Task
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	t.Title = strings.TrimSpace(t.Title)
	if t.Title == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(t.Status) == "" {
		t.Status = "Pending"
	}

	mu.Lock()
	t.ID = nextID
	nextID++
	tasks = append(tasks, t)
	mu.Unlock()

	writeJSON(w, http.StatusCreated, t)
}

func listTasks(w http.ResponseWriter) {
	mu.Lock()
	list := make([]models.Task, len(tasks))
	copy(list, tasks)
	mu.Unlock()

	writeJSON(w, http.StatusOK, list)
}

func getTask(w http.ResponseWriter, id int) {
	mu.Lock()
	defer mu.Unlock()

	for _, t := range tasks {
		if t.ID == id {
			writeJSON(w, http.StatusOK, t)
			return
		}
	}
	http.Error(w, "Task not found", http.StatusNotFound)
}

func updateTask(w http.ResponseWriter, r *http.Request, id int) {
	var in models.Task
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	in.Title = strings.TrimSpace(in.Title)
	if in.Title == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Title = in.Title
			tasks[i].Description = in.Description
			if strings.TrimSpace(in.Status) != "" {
				tasks[i].Status = in.Status
			}
			writeJSON(w, http.StatusOK, tasks[i])
			return
		}
	}
	http.Error(w, "Task not found", http.StatusNotFound)
}

func deleteTask(w http.ResponseWriter, id int) {
	mu.Lock()
	defer mu.Unlock()

	for i := range tasks {
		if tasks[i].ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Task not found", http.StatusNotFound)
}
