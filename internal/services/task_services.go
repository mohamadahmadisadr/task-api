package services

import (
	"task-api/internal/models"
)

type TaskService struct {
	tasks []models.Task
}

func NewTaskService() *TaskService {
	return &TaskService{
		tasks: []models.Task{},
	}
}

func (s *TaskService) GetTask() []models.Task {
	return s.tasks
}

func (s *TaskService) CreateTask(task models.Task) models.Task {

	task.ID = len(s.tasks) + 1
	s.tasks = append(s.tasks, task)
	return task
}
