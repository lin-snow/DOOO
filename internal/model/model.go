/*
model package is responsible for defining the data models.
*/
package model

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

// Code
const (
	// Code For Success
	SUCCESS = 666
	// Code For Error (Server)
	ERR_JSON = 1000
	// Code For Error (USER)
	ERR_USER_NOT_FOUND = 2001
	ERR_USER_NOT_SAME  = 2002

	ERR_PASSWORD_WRONG = 2011
	ERR_PASSWORD_BLANK = 2012
	ERR_PASSWORD_SAME  = 2013

	ERR_USERNAME_EXIST = 2021
	ERR_USERNAME_BLANK = 2022
	ERR_USERNAME_SAME  = 2023

	ERR_EMAIL_EXIST = 2031
	ERR_EMAIL_BLANK = 2032
	ERR_EMAIL_SAME  = 2033

	ERR_NICKNAME_BLANK = 2041
	// Code For Error (TODO)
	ERR_TODO_NOT_FOUND = 3001
	ERR_TODO_WRONG     = 3002

	// Code For Error (AUTH)
	ERR_AUTH_NOT_FOUND = 4001
	ERR_AUTH_INVALID   = 4002
	ERR_TOKEN_INVALID  = 4003
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
