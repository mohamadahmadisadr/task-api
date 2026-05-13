package main

import (
	"log"
	"net/http"
	"task-api/internal/config"
	"task-api/internal/handlers"
	"task-api/internal/services"
)

func startServer() {

	taskService := services.NewTaskService()
	taskHandler := handlers.NewTaskHandler(taskService)
	router := handlers.RegisterRoutes(taskHandler)
	cfg := config.Load()
	log.Printf("server running on port :%s", cfg.Port)
	err := http.ListenAndServe(":"+cfg.Port, router)
	if err != nil {
		log.Fatal(err)
	}
}
