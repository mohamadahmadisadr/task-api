package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"task-api/internal/config"
	"task-api/internal/database"
	"task-api/internal/handlers"
	"task-api/internal/services"
	"time"
)

func startServer() {

	db := database.Connect()
	defer db.Close()

	taskService := services.NewTaskService()
	taskHandler := handlers.NewTaskHandler(taskService)
	router := handlers.RegisterRoutes(taskHandler)
	cfg := config.Load()

	srv := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		log.Printf("server running on port :%s", cfg.Port)
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Println("shutting down server...")

	err := srv.Shutdown(ctx)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("server stopped")
}
