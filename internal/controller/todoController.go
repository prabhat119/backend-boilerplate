package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prabhat119/backend-boilerplate/internal/service"
)

type TodoController struct {
	service service.TodoService
}

func NewTodoController(s service.TodoService) *TodoController {
	return &TodoController{service: s}
}

func (ctrl *TodoController) FetchTodos(c *gin.Context) {
	todos, err := ctrl.service.GetTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todos)
}
