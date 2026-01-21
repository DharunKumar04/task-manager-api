package handlers

import (
	"net/http"
	"time"

	"github.com/DharunKumar04/task-manager-api/models"
	"github.com/gin-gonic/gin"
)

type CreateTaskRequest struct {
	Name        string `json:"name" binding:"required,min=5"`
	UserID      uint   `json:"userid" binding:"required"`
	Description string `json:"description" binding:"required"`
	ProjectID   uint   `json:"project" binding:"required"`
	DueDate     string `json:"duedate" binding:"required,datetime=2006-01-02"`
	Status      string `json:"status" binding:"required"`
}

type GetTaskResponse struct {
	ID          uint            `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	User        GetUserResponse `json:"user"`
	ProjectName string          `json:"projectname"`
	CreatedAt   time.Time       `json:"createdAt"`
}

func (h *Handler) GetTask(ctx *gin.Context) {
	taskID := ctx.Param("id")

	var task models.Task

	result := h.DB.Preload("User").Preload("Project").First(&task, taskID)

	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	response := GetTaskResponse{
		ID:          task.ID,
		Name:        task.Title,
		Description: task.Description,
		User: GetUserResponse{
			ID:    task.User.ID,
			Email: task.User.Email,
		},
		ProjectName: task.Project.Name,
		CreatedAt:   task.CreatedAt,
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Task fetched Successfully",
		"data":    response,
	})

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
		UserID:      req.UserID,
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
		"message": "Task Created Successfully",
		"name":    newTask.Title,
		"id":      newTask.ID,
	})

}
