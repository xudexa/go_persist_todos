package persistence

import (
	"github.com/jmoiron/sqlx"
	domain "github.com/xudexa/go_domain_todos/domain"
)

type TodoRepository struct {
	DB *sqlx.DB
}

func NewTodoRepository(db *sqlx.DB) *TodoRepository {
	return &TodoRepository{
		DB: db,
	}
}

func (r *TodoRepository) AddTodo(todo *domain.Todo) error {
	query := `INSERT INTO todos (id, title, description, importance, urgency, planning_date, expected_work, actual_work, state, status, assigned_person, created_at, last_modified_at)
			  VALUES (:id, :title, :description, :importance, :urgency, :planning_date, :expected_work, :actual_work, :state, :status, :assigned_person, :created_at, :last_modified_at)`

	_, err := r.DB.NamedExec(query, todo)
	return err
}

func (r *TodoRepository) UpdateTodo(todo *domain.Todo) error {
	query := `UPDATE todos SET title=:title, description=:description, importance=:importance, urgency=:urgency, planning_date=:planning_date, expected_work=:expected_work, actual_work=:actual_work, state=:state, status=:status, assigned_person=:assigned_person, last_modified_at=:last_modified_at
			  WHERE id=:id`

	_, err := r.DB.NamedExec(query, todo)
	return err
}

func (r *TodoRepository) DeleteTodo(id string) error {
	query := `DELETE FROM todos WHERE id = :id`

	_, err := r.DB.NamedExec(query, map[string]interface{}{
		"id": id,
	})
	return err
}

func (r *TodoRepository) GetAllTodos() (domain.Todos, error) {
	todos := domain.Todos{}
	query := `SELECT * FROM todos WHERE status = 'actived' AND state != 'done'`

	if err := r.DB.Select(&todos, query); err != nil {
		return nil, err
	}
	return todos, nil
}

func (r *TodoRepository) GetAllArchivedTodos() (domain.Todos, error) {
	todos := domain.Todos{}
	query := `SELECT * FROM todos WHERE status = 'archived' ORDER BY last_modified_at DESC`

	if err := r.DB.Select(&todos, query); err != nil {
		return nil, err
	}
	return todos, nil
}

func (r *TodoRepository) GetTodo(id string) (*domain.Todo, error) {
	todo := &domain.Todo{}
	query := `SELECT * FROM todos WHERE id = :id`

	if err := r.DB.Get(todo, query, map[string]interface{}{
		"id": id,
	}); err != nil {
		return nil, err
	}
	return todo, nil
}
