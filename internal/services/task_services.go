package services

import (
	"task-api/internal/models"
	"task-api/internal/repository"
)

type TaskService struct {
	repo *repository.TaskRepository
}

func NewTaskService(repo *repository.TaskRepository) *TaskService {
	return &TaskService{
		repo: repo,
	}
}

func (s *TaskService) GetTask() ([]models.Task, error) {
	return s.repo.GetTasks()
}

func (s *TaskService) CreateTask(task models.Task) (*models.Task, error) {
	return s.repo.CreateTask(task)
}

func (s *TaskService) GetTaskById(id int) (*models.Task, bool) {
	return s.repo.GetTaskByID(id)
}

func (s *TaskService) DeleteTaskByID(id int) bool {
	return s.repo.DeleteTaskByID(id)
}

func (s *TaskService) UpdateTaskByID(id int, updated models.Task) (*models.Task, bool) {
	return s.repo.UpdateTaskByID(id, updated)
}
