package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"task-api/internal/dto"
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
	tasks, err := h.service.GetTask()
	if err != nil {
		writeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

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
	created, err := h.service.CreateTask(task)
	if err != nil {
		writeError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	resp := dto.TaskResponse{
		ID:   created.ID,
		Name: created.Name,
		Done: task.Done,
	}
	json.NewEncoder(w).Encode(resp)

}

func (h *TaskHandler) TaskByID(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	parts := strings.Split(path, "/")
	idStr := parts[2]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeError(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		task, found := h.service.GetTaskById(id)
		if !found {
			writeError(w, "Task Not Found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(task)
	case http.MethodDelete:
		ok := h.service.DeleteTaskByID(id)
		if !ok {
			writeError(w, "Task Not Found", http.StatusNotFound)
		}
		w.WriteHeader(http.StatusNoContent)
	case http.MethodPut:
		var req CreateTaskRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			writeError(w, "invalid request body", http.StatusBadRequest)
			return
		}
		if errMsg := validateTask(models.Task{Name: req.Name, Done: req.Done}); errMsg != "" {
			writeError(w, errMsg, http.StatusBadRequest)
			return
		}
		updated, ok := h.service.UpdateTaskByID(id, models.Task{
			Name: req.Name,
			Done: req.Done,
		})
		if !ok {
			writeError(w, "task not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(updated)

	default:
		writeError(w, "Method Not allowed", http.StatusBadRequest)

	}

}
