package service

import (
	"github.com/prabhat119/backend-boilerplate/internal/model"
	"github.com/prabhat119/backend-boilerplate/internal/repository"
)

type TodoService interface {
	GetTodos() ([]model.Todo, error)
}

type todoService struct {
	repo repository.TodoRepository
}

func NewTodoService(repo repository.TodoRepository) TodoService {
	return &todoService{repo: repo}
}

func (s *todoService) GetTodos() ([]model.Todo, error) {
	return s.repo.GetAll()
}
