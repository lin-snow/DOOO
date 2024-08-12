/*
model package is responsible for defining the data models.
*/
package model

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
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

// JWT Model
// MySigningKey
var MySigningKey = []byte("DOOOOOOOOOOOOOOO!")

// Claim model
type MyClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// UserToUpdate Model
type UserToUpdate struct {
	Username string `json:"username,omitempty"`
	Nickname string `json:"nickname,omitempty"`
	Password string `json:"password,omitempty"`
	Email    string `json:"email,omitempty"`
}

// Define the delete detail
type DeleteDetail struct {
	Usertodel string `json:"usertodel"`
	Deldata   bool   `json:"deldata"`
}
