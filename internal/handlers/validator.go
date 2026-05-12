package handlers

import "task-api/internal/models"

func validateTask(task models.Task) string {
	if task.Name == "" {
		return "name is required"
	}

	return ""
}
