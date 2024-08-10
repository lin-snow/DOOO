/*
model package is responsible for defining the data models.
*/
package model

import (
	"time"

	"gorm.io/gorm"
)

// Toda Model
type Todo struct {
	gorm.Model            // ID, CreatedAt, UpdatedAt, DeletedAt
	Title       string    `json:"title"`
	Description string    `json:"description"`
	IsCompleted bool      `json:"isCompleted" gorm:"default:false"`
	DueDate     time.Time `json:"dueDate"`

	// Foreign key
	UserID uint `json:"userId" gorm:"not null"`
}

// User Model
type User struct {
	gorm.Model        // ID, CreatedAt, UpdatedAt, DeletedAt
	Username   string `json:"username" gorm:"not null;unique"`
	Nickname   string `json:"nickname"`
	Password   string `json:"password" gorm:"not null"`
	Email      string `json:"email" gorm:"not null;unique"`
	Role       string `json:"role" gorm:"default:'user'"`

	// Foreign key
	Todos []Todo `json:"todos"`
}
