/*
todo package contains the todo model and handler.
*/
package todo

import (
	"github.com/gin-gonic/gin"
	"github.com/lin-snow/dooo/internal/model"
	"gorm.io/gorm"
)

// Create a new todo
func AddTodo(ctx *gin.Context, db *gorm.DB) {
	// Get the todo data
	var tododata model.Todo
	err := ctx.ShouldBindBodyWithJSON(&tododata)
	if err != nil {
		// Failed with JSON
		ctx.JSON(200, gin.H{
			"Message": "Failed with JSON!!!",
			"code":    400,
			"data":    tododata,
		})
	}

	// Add Todo to DB
	db.Create(&tododata)

	// Return the todo data
	ctx.JSON(200, gin.H{
		"Message": "Todo Added!!!",
		"code":    200,
		"data":    tododata,
	})
}
