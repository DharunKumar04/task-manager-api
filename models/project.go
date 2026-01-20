package models

import (
	"time"
)

type Project struct {
	ID          uint `gorm:"primaryKey"`
	User        User
	UserID      uint `gorm:"foreignKey"`
	Tasks       []Task
	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
	CreatedAt   time.Time
}
