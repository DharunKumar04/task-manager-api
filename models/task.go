package models

import (
	"time"
)

type Task struct {
	ID          uint   `gorm:"primaryKey"`
	ProjectID   uint   `gorm:"foreignKey"`
	Title       string `gorm:"not null"`
	Description string `gorm:"not null"`
	Project     Project
	CreatedAt   time.Time
	DueDate     time.Time
	Status      string `gorm:"default:pending"`
}
