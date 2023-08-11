package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

type TodoResponse struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Completed   bool      `json:"completed" default:"false"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at,omitempty"`
}
