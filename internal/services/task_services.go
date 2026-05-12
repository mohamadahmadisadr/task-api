package services

import (
	"sync"
	"task-api/internal/models"
)

var mu sync.Mutex

var Tasks = []models.Task{}

func GetTask() []models.Task {
	mu.Lock()
	defer mu.Unlock()
	return Tasks
}

func CreateTask(task models.Task) models.Task {

	mu.Lock()
	defer mu.Unlock()

	task.ID = len(Tasks) + 1
	Tasks = append(Tasks, task)
	return task
}
