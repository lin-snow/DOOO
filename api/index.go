/*
api package is the package that contains the API handlers for the Outerspace API.
*/
package api

import (
	"github.com/gin-gonic/gin"
	"github.com/lin-snow/dooo/internal/todo"
	"github.com/lin-snow/dooo/internal/user"
	"github.com/lin-snow/dooo/pkg/auth"
	"gorm.io/gorm"
)

func Router(r *gin.Engine, db *gorm.DB) *gin.Engine {
	// API Router
	api := r.Group("/api", auth.JWTAuth()) // JWT Authentication
	{
		// User Sign Out

		// User Update

		// User Delete

		// Todo Create
		api.POST("/add", func(ctx *gin.Context) {
			// AddTodo
			todo.AddTodo(ctx, db)
		})

		// Todo Query
		// Query All
		// Query By Pagination

		// Todo Update

		// Todo Delete

	}

	// Not Use Middleware
	// User Sign Up
	r.POST("/signup", func(ctx *gin.Context) {
		user.Register(ctx, db)
	})

	// User Login
	r.POST("/login", func(ctx *gin.Context) {
		user.Login(ctx, db)
	})

	return r
}
