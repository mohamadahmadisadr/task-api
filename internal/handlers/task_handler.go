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
		writeError(w, "Not supported", http.StatusBadRequest)
	}
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, "Only Get Method is allowed", http.StatusMethodNotAllowed)
		return
	}
	tasks := services.GetTask()

	json.NewEncoder(w).Encode(tasks)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, "Only Post Method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		writeError(w, "invalid request", http.StatusBadRequest)
		return
	}
	if errMsg := validateTask(task); errMsg != "" {
		writeError(w, errMsg, http.StatusBadRequest)
		return
	}
	created := services.CreateTask(task)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)

}
