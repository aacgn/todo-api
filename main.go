package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Task struct
type Task struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

// Init task var as a slice Task struct
var tasks []Task

// Get all tasks
func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

// Get single tasks
func getTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params
	// Loop through tasks and find with id
	for _, item := range tasks {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Task{})
}

// Create a new task
func createTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var task Task
	_ = json.NewDecoder(r.Body).Decode(&task)
	task.ID = strconv.Itoa(rand.Intn(1000000)) // Mock ID
	tasks = append(tasks, task)
	json.NewEncoder(w).Encode(task)
}

// Update a task
func updateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range tasks {
		if item.ID == params["id"] {
			tasks = append(tasks[:index], tasks[index+1:]...)
			var task Task
			_ = json.NewDecoder(r.Body).Decode(&task)
			tasks = append(tasks, task)
			json.NewEncoder(w).Encode(task)
			return
		}
	}
}

// Delete a task
func deleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range tasks {
		if item.ID == params["id"] {
			tasks = append(tasks[:index], tasks[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(tasks)
}

func main() {
	// Init Router
	router := mux.NewRouter()

	// Mock data - @todo - implement database
	tasks = append(tasks, Task{ID: "1", Title: "Task One", Description: "Description for Task One", Completed: true})
	tasks = append(tasks, Task{ID: "2", Title: "Task Two", Description: "Description for Task Two", Completed: false})

	// Route handlers
	router.HandleFunc("/api/v1/tasks", getTasks).Methods("GET")
	router.HandleFunc("/api/v1/tasks/{id}", getTask).Methods("GET")
	router.HandleFunc("/api/v1/tasks", createTask).Methods("POST")
	router.HandleFunc("/api/v1/task/{id}", updateTask).Methods("PUT")
	router.HandleFunc("/api/v1/task/{id}", deleteTask).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
