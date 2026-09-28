package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/prabhat119/backend-boilerplate/internal/model"
)

type TodoRepository interface {
	GetAll() ([]model.Todo, error)
}

type todoRepository struct {
	db *sqlx.DB
}

func NewTodoRepository(db *sqlx.DB) TodoRepository {
	return &todoRepository{db: db}
}

func (r *todoRepository) GetAll() ([]model.Todo, error) {
	todos := []model.Todo{}

	// mock
	todos = append(todos, model.Todo{ID: 1, Title: "Configure Go Backend Boilerplate", Completed: true})
	return todos, nil
}
