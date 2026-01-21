package routes

import (
	"github.com/DharunKumar04/task-manager-api/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, h *handlers.Handler) {
	router.GET("/ping", h.Ping)
	router.POST("/users", h.CreateUser)
	router.GET("/projects/:id", h.GetProject)
	router.POST("/projects", h.CreateProject)
	router.GET("/tasks/:id", h.GetTask)
	router.POST("/tasks", h.CreateTask)
}
