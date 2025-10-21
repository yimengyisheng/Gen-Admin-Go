package model

import (
	"time"
	"gorm.io/gorm"
)

// BaseModel defines the common fields for all models, with snake_case json tags.
// It replaces gorm.Model to ensure consistent API responses.
type BaseModel struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
