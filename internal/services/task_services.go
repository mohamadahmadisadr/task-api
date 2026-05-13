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

func (s *TaskService) GetTaskById(id int) (*models.Task, bool) {
	for _, t := range s.tasks {
		if t.ID == id {
			return &t, true
		}
	}
	return nil, false
}

func (s *TaskService) DeleteTaskByID(id int) bool {
	for i, t := range s.tasks {
		if t.ID == id {
			s.tasks = append(s.tasks[:i], s.tasks[i+1:]...)
			return true
		}
	}
	return false
}

func (s *TaskService) UpdateTaskByID(id int, updated models.Task) (*models.Task, bool) {
	for i, t := range s.tasks {
		if t.ID == id {
			s.tasks[i].Name = updated.Name
			s.tasks[i].Done = updated.Done
			return &s.tasks[i], true
		}
	}
	return nil, false
}
