/*
api package is the package that contains the API handlers for the Outerspace API.
*/
package api

import (
	"github.com/gin-gonic/gin"
	"github.com/lin-snow/dooo/internal/category"
	"github.com/lin-snow/dooo/internal/todo"
	"github.com/lin-snow/dooo/internal/user"
	"github.com/lin-snow/dooo/pkg/auth"
	"gorm.io/gorm"
)

func Router(r *gin.Engine, db *gorm.DB) *gin.Engine {
	// API Router
	api := r.Group("/api", auth.JWTAuth()) // JWT Authentication
	{
		//=========================(USER)===============================
		// User Update
		// (Need to Re-Login)
		// Update Username
		api.PUT("/updateusername", func(ctx *gin.Context) {
			// Update Username
			user.UpdateUsername(ctx, db)
		})
		// Update Password
		api.PUT("/updatepassword", func(ctx *gin.Context) {
			// Update Password
			user.UpdatePassword(ctx, db)
		})

		// (Do not need to Re-Login)
		// Update Email
		api.PUT("/updateemail", func(ctx *gin.Context) {
			// Update Email
			user.UpdateEmail(ctx, db)
		})
		// Update Nickname
		api.PUT("/updatenickname", func(ctx *gin.Context) {
			// Update Nickname
			user.UpdateNickname(ctx, db)
		})

		// User Delete
		api.DELETE("/deleteaccount", func(ctx *gin.Context) {
			// Delete Account
			user.DeleteAccount(ctx, db)
		})

		// User Get
		api.GET("/getuserinfo", func(ctx *gin.Context) {
			// Get User Info
			user.GetUserInfo(ctx, db)
		})

		//==========================(TODO)================================
		// Todo Create
		api.POST("/add", func(ctx *gin.Context) {
			// AddTodo
			todo.AddTodo(ctx, db)
		})

		// Todo Query
		api.GET("/querytodo", func(ctx *gin.Context) {
			// Query Todo
			todo.QueryTodo(ctx, db)
		})

		// Todo Update
		api.PUT("/updatetodo", func(ctx *gin.Context) {
			// Update Todo
			todo.UpdateTodo(ctx, db)
		})

		// Todo Delete
		api.DELETE("/deletetodo", func(ctx *gin.Context) {
			// Delete Todo
			todo.DeleteTodo(ctx, db)
		})

		//==========================(CATEGORY)===========================
		// Category Create
		api.POST("/addcategory", func(ctx *gin.Context) {
			// Add Category
			category.AddCategory(ctx, db)
		})

		// Category Query
		api.GET("/querycategory", func(ctx *gin.Context) {
			// Query Category
			category.QueryCategories(ctx, db)
		})

		// Category Update
		api.PUT("/updatecategory", func(ctx *gin.Context) {
			// Update Category
			category.UpdateCategory(ctx, db)
		})

		// Category Delete
		api.DELETE("/deletecategory", func(ctx *gin.Context) {
			// Delete Category
			category.DeleteCategory(ctx, db)
		})

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
