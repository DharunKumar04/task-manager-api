package routes

import (
	"github.com/DharunKumar04/task-manager-api/handlers"
	"github.com/DharunKumar04/task-manager-api/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, h *handlers.Handler) {
	router.GET("/ping", h.Ping)
	router.POST("/users", h.CreateUser)
	router.POST("/login", h.Login)

	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/projects/:id", h.GetProject)
		protected.POST("/projects", h.CreateProject)

		protected.GET("/tasks/:id", h.GetTask)
		protected.POST("/tasks", h.CreateTask)
	}
}
