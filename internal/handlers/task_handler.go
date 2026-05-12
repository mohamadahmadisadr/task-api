package handlers

import (
	"encoding/json"
	"net/http"
	"task-api/internal/models"
	"task-api/internal/services"
)

func TaskHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		GetTask(w, r)
	case http.MethodPost:
		CreateTask(w, r)
	default:
		http.Error(w, "Not supported", http.StatusBadRequest)
	}
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GEt Method is allowed", http.StatusMethodNotAllowed)
		return
	}
	tasks := services.GetTask()

	json.NewEncoder(w).Encode(tasks)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only Post Method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	if task.Name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}
	created := services.CreateTask(task)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
}
