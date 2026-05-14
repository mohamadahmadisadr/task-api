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

	tasks := []models.Task{}

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

func (r *TaskRepository) CreateTask(task models.Task) (*models.Task, error) {
	query := `
		INSERT INTO tasks (name, done)
		VALUES ($1, $2)
		RETURNING id
	`
	err := r.db.QueryRow(
		query,
		task.Name,
		task.Done,
	).Scan(&task.ID)

	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (r *TaskRepository) GetTaskByID(id int) (*models.Task, bool) {
	row := r.db.QueryRow(`
		SELECT id, name, done
		FROM tasks
		WHERE id = $1
	`, id)

	var task models.Task
	err := row.Scan(
		&task.ID,
		&task.Name,
		&task.Done,
	)

	if err != nil {
		return nil, false
	}

	return &task, true
}

func (r *TaskRepository) DeleteTaskByID(id int) bool {
	result, err := r.db.Exec(`
		DELETE FROM tasks
		WHERE id = $1
	`, id)

	if err != nil {
		return false
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false
	}

	return rowsAffected > 0
}

func (r *TaskRepository) UpdateTaskByID(id int, updated models.Task) (*models.Task, bool) {
	row := r.db.QueryRow(`
		SELECT id, name, done
		FROM tasks
		WHERE id = $1
	`, id)

	var task models.Task
	err := row.Scan(
		&task.ID,
		&task.Name,
		&task.Done,
	)

	if err != nil {
		return nil, false
	}

	_, err = r.db.Exec(`
		UPDATE tasks
		SET name = $1, done = $2
		WHERE id = $3
	`, updated.Name, updated.Done, id)

	if err != nil {
		return nil, false
	}

	return &task, true
}
