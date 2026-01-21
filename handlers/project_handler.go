package handlers

import (
	"net/http"

	"github.com/DharunKumar04/task-manager-api/models"
	"github.com/gin-gonic/gin"
)

type CreateProjectRequest struct {
	Name        string `json:"name" binding:"required,min=5"`
	Description string `json:"description" binding:"required"`
	UserID      uint   `json:"userid" binding:"required"`
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
