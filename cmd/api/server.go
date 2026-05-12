package main

import (
	"log"
	"net/http"
	"task-api/internal/handlers"
)

func startServer() {
	http.HandleFunc("/tasks", handlers.Logger(handlers.TaskHandler))
	log.Println("server running on port :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
