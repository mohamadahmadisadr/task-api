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
	router := handlers.RegisterRoutes(taskHandler)
	log.Println("server running on port :8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
