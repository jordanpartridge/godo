package repository

import (
	"database/sql"

	"github.com/jordanpartridge/godo/internal/model"
)

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) Create(title, description string) (*model.Task, error) {
	query := `INSERT INTO tasks (title, description) VALUES (?, ?) RETURNING id, created_at, updated_at`

	task := &model.Task{
		Title:       title,
		Description: description,
		Completed:   false,
	}

	err := r.db.QueryRow(query, title, description).Scan(&task.ID, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (r *TaskRepository) All() ([]model.Task, error) {
	query := `SELECT id, title, description, completed, created_at, updated_at FROM tasks ORDER BY created_at ASC`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []model.Task
	for rows.Next() {
		var task model.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Completed, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (r *TaskRepository) MarkComplete(id int) error {
	query := `UPDATE tasks SET completed = TRUE, updated_at = CURRENT_TIMESTAMP WHERE id = ?`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *TaskRepository) Delete(id int) error {
	query := `DELETE FROM tasks WHERE id = ?`
	_, err := r.db.Exec(query, id)
	return err
}
