package handlers

import (
	"net/http"
)

func RegisterRoutes(taskHandler *TaskHandler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path == "/tasks" {
			switch r.Method {
			case http.MethodGet:
				taskHandler.GetTask(w, r)
			case http.MethodPost:
				taskHandler.CreateTask(w, r)

			default:
				writeError(w, "Not Implemented", http.StatusBadRequest)
			}

			return
		}

	})

	mux.HandleFunc("/tasks/{id}", func(w http.ResponseWriter, r *http.Request) {
		taskHandler.TaskByID(w, r)
	})
	return Recovery(Logger(mux))
}
