package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prabhat119/backend-boilerplate/internal/controller"
)

func Routes(router *gin.Engine, injection *Injection) {
	router.Use(func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusNoContent)
			return
		}
		ctx.Next()
	})

	api := router.Group("/api/v1")
	{
		// Public Routes
		// api.POST("/login", container.AuthController.Login)

		// Secure Routes
		protected := api.Group("/")
		protected.Use(controller.AuthMiddleware(injection.JWTService))
		{
			protected.GET("/todos", injection.TodoController.FetchTodos)
		}
	}
}
