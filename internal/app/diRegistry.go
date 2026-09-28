package app

import (
	"github.com/jmoiron/sqlx"
	"github.com/prabhat119/backend-boilerplate/internal/config"
	"github.com/prabhat119/backend-boilerplate/internal/controller"
	"github.com/prabhat119/backend-boilerplate/internal/repository"
	"github.com/prabhat119/backend-boilerplate/internal/service"
)

type Injection struct {
	JWTService     service.JWTService
	TodoController *controller.TodoController
}

func NewInjection(db *sqlx.DB, cfg config.Config) *Injection {

	jwtService := service.NewJWTService(cfg.JWTSecret)

	todoRepo := repository.NewTodoRepository(db)
	todoService := service.NewTodoService(todoRepo)
	todoController := controller.NewTodoController(todoService)

	return &Injection{
		JWTService:     jwtService,
		TodoController: todoController,
	}
}
