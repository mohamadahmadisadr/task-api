package handlers

import (
	"net/http"
)

func RegisterRoutes(taskHandler *TaskHandler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			taskHandler.GetTask(w, r)
			return
		}

		if r.Method == http.MethodPost {
			taskHandler.CreateTask(w, r)
			return
		}

		writeError(w, "Method now allowed", http.StatusBadRequest)
	})
	return Logger(mux)
}
