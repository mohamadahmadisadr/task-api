package handlers

import (
	"encoding/json"
	"net/http"
	"task-api/internal/models"
	"task-api/internal/services"
)

type TaskHandler struct {
	service *services.TaskService
}

func NewTaskHandler(service *services.TaskService) *TaskHandler {
	return &TaskHandler{
		service: service,
	}
}

func (h *TaskHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, "Only Get Method is allowed", http.StatusMethodNotAllowed)
		return
	}
	tasks := h.service.GetTask()

	json.NewEncoder(w).Encode(tasks)
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, "Only Post Method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var req CreateTaskRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	task := models.Task{
		Name: req.Name,
		Done: req.Done,
	}
	if err != nil {
		writeError(w, "invalid request", http.StatusBadRequest)
		return
	}
	if errMsg := validateTask(task); errMsg != "" {
		writeError(w, errMsg, http.StatusBadRequest)
		return
	}
	created := h.service.CreateTask(task)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)

}
