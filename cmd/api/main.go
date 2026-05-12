package main

import (
	"encoding/json"
	"net/http"
	"task-api/internal/handlers"
)

type HealthResponse struct {
	Status string `json:"status"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	if r.Method != http.MethodGet {
		http.Error(w, "Only Get Method is allowed", http.StatusMethodNotAllowed)
		return
	}

	response := HealthResponse{
		Status: "oK",
	}

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Internal Service Error", http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/tasks", handlers.TaskHandler)
	http.ListenAndServe(":8080", nil)
}
