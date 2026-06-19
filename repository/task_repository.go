package repository

import (
	"database/sql"
	"github.com/lucasmilhao/go-crud/models"
)

type TaskRepository interface {
	FindAll() ([]models.Task, error)
	FindById(id int) (*models.Task, error)
	Save(task models.Task) error
	Update(id int, task models.Task) error
	Delete(id int) (bool, error)
}

type taskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) FindAll() ([]models.Task, error) {
	rows, err := r.db.Query("SELECT * FROM tasks")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	tasks := make([]models.Task, 0)

	for rows.Next() {
		var task models.Task

		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status); err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (r *taskRepository) FindById(id int) (*models.Task, error) {
	var task models.Task

	err := r.db.QueryRow("SELECT * FROM tasks WHERE id = ?", id).Scan(&task.ID, &task.Title, &task.Description, &task.Status)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (r *taskRepository) Save(task models.Task) error {
	_, err := r.db.Exec("INSERT INTO tasks (title, description, status) VALUES (?, ?, ?)",
		task.Title, task.Description, task.Status,
	)

	return err
}

func (r *taskRepository) Update(id int, task models.Task) error {
	_, err := r.db.Exec("UPDATE tasks SET title = ?, description = ?, status = ? WHERE id = ?", task.Title, task.Description, task.Status, id)

	return err
}

func (r *taskRepository) Delete(id int) (bool, error) {
	result, err := r.db.Exec("DELETE FROM tasks WHERE id = ?", id)

	if err != nil {
		return false, err
	}

	rows, err := result.RowsAffected()

	return rows > 0, err
}
