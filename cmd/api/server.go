package main

import (
	"log"
	"net/http"
	"task-api/internal/handlers"
	"task-api/internal/services"
)

func startServer() {

	taskService := services.NewTaskService()
	taskHandler := handlers.NewTaskHandler(taskService)
	http.HandleFunc("/tasks", handlers.Logger(taskHandler.TaskHandler))
	log.Println("server running on port :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
