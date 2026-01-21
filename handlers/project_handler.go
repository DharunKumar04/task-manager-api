package handlers

import (
	"net/http"
	"time"

	"github.com/DharunKumar04/task-manager-api/models"
	"github.com/gin-gonic/gin"
)

type CreateProjectRequest struct {
	Name        string `json:"name" binding:"required,min=5"`
	Description string `json:"description" binding:"required"`
	UserID      uint   `json:"userid" binding:"required"`
}

type GetProjectResponse struct {
	ID          uint            `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	User        GetUserResponse `json:"user"`
	Tasks       []models.Task   `json:"tasks"`
	CreatedAt   time.Time       `json:"createdAt"`
}

func (h *Handler) GetProject(ctx *gin.Context) {
	projectID := ctx.Param("id")

	var project models.Project

	result := h.DB.Preload("User").Preload("Tasks").First(&project, projectID)

	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}
	response := GetProjectResponse{
		ID:          project.ID,
		Name:        project.Name,
		Description: project.Description,
		User: GetUserResponse{
			ID:    project.User.ID,
			Email: project.User.Email,
		},
		Tasks:     project.Tasks,
		CreatedAt: project.CreatedAt,
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Project fetched Successfully",
		"data":    response,
	})
}

func (h *Handler) CreateProject(ctx *gin.Context) {
	var req CreateProjectRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newProject := models.Project{
		Name:        req.Name,
		Description: req.Description,
		UserID:      req.UserID,
	}

	if result := h.DB.Create(&newProject); result.Error != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": "Failed to create new Project"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Project Created Successfully",
		"name":    newProject.Name,
	})
}
