/*
api package is the package that contains the API handlers for the Outerspace API.
*/
package api

import (
	"github.com/gin-gonic/gin"
	"github.com/lin-snow/dooo/internal/user"
	"gorm.io/gorm"
	// "github.com/lin-snow/dooo/internal/todo"
)

func Router(r *gin.Engine, db *gorm.DB) *gin.Engine {
	// User Sign Up
	r.POST("/signup", func(ctx *gin.Context) {
		user.CreateUser(ctx, db)
	})

	// User Sign In

	// User Sign Out

	// User Update

	// User Delete

	// Todo Create

	// Todo Query
	// Query All
	// Query By Pagination

	// Todo Update

	// Todo Delete

	return r
}
