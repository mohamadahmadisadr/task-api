package repository

import (
	"database/sql"
	"task-api/internal/models"
)

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{
		db: db,
	}
}

func (r *TaskRepository) GetTasks() ([]models.Task, error) {
	rows, err := r.db.Query(`
		SELECT id, name, done
		FROM tasks
	`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tasks []models.Task

	for rows.Next() {
		var task models.Task
		err := rows.Scan(
			&task.ID,
			&task.Name,
			&task.Done,
		)

		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}
	return tasks, nil
}
