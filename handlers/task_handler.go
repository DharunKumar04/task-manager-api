package handlers

import (
	"net/http"
	"time"

	"github.com/DharunKumar04/task-manager-api/models"
	"github.com/gin-gonic/gin"
)

type CreateTaskRequest struct {
	Name        string `json:"name" binding:"required,min=5"`
	Description string `json:"description" binding:"required"`
	ProjectID   uint   `json:"project" binding:"required"`
	DueDate     string `json:"duedate" binding:"required,datetime=2006-01-02"`
	Status      string `json:"status" binding:"required"`
}

func (h *Handler) CreateTask(ctx *gin.Context) {
	var req CreateTaskRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	dueDate, err := time.Parse("2006-01-02", req.DueDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid date format. Use YYYY-MM-DD"})
		return
	}

	newTask := models.Task{
		Title:       req.Name,
		Description: req.Description,
		ProjectID:   req.ProjectID,
		DueDate:     dueDate,
		Status:      req.Status,
	}

	if result := h.DB.Create(&newTask); result.Error != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": "Failed to create new Project"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Project Created Successfully",
		"name":    newTask.Title,
	})

}
