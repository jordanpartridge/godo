package repository

import (
	"database/sql"
	"time"

	"github.com/jordanpartridge/godo/internal/model"
)

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) Create(title, description string) (*model.Task, error) {
	query := `INSERT INTO tasks (title, description) VALUES (?, ?)`

	result, err := r.db.Exec(query, title, description)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	task := &model.Task{
		ID:          int(id),
		Title:       title,
		Description: description,
		Completed:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
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
