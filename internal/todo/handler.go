/*
todo package contains the todo model and handler.
*/
package todo

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lin-snow/dooo/internal/model"
	"github.com/lin-snow/dooo/pkg/auth"
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
			"code":    model.ERR_JSON,
			"data":    tododata,
		})

		return
	}

	// Check the todo data
	curuser := auth.GetCurrentUser(ctx, db)
	if tododata.Title == "" || tododata.UserID == 0 || tododata.UserID != curuser.ID {
		// The todo data has wrong
		ctx.JSON(200, gin.H{
			"Message": "The todo data has wrong!",
			"code":    model.ERR_TODO_WRONG,
			"data": gin.H{
				"todo": tododata,
			},
		})
		return
	}

	// Add Todo to DB
	db.Create(&tododata)

	// Return the todo data
	ctx.JSON(200, gin.H{
		"Message": "Todo Added!!!",
		"code":    model.SUCCESS,
		"data":    tododata,
	})
}

// Query Todo
func QueryTodo(ctx *gin.Context, db *gorm.DB) {
	var tododata []model.Todo
	var total int64

	// Current User
	curuser := auth.GetCurrentUser(ctx, db)

	pageSize, _ := strconv.Atoi(ctx.Query("pageSize"))
	pageNum, _ := strconv.Atoi(ctx.Query("pageNum"))

	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}

	offsetVal := (pageNum - 1) * pageSize
	if pageSize == -1 && pageNum == -1 {
		offsetVal = -1
	}

	db.Model(&model.Todo{}).Where("user_id = ?", curuser.ID).Count(&total).Limit(pageSize).Offset(offsetVal).Find(&tododata)

	if len(tododata) == 0 {
		ctx.JSON(200, gin.H{
			"Message": "Query Failed!!!",
			"code":    model.ERR_TODO_NOT_FOUND,
			"data":    gin.H{},
		})
	} else {
		ctx.JSON(200, gin.H{
			"Message": "Query Successfully!",
			"code":    model.SUCCESS,
			"data": gin.H{
				"list":     tododata,
				"total":    total,
				"pageSize": pageSize,
				"pageNum":  pageNum,
			},
		})
	}

}

// Update Todo
func UpdateTodo(ctx *gin.Context, db *gorm.DB) {
	// Get the new todo data
	var new model.Todo
	err := ctx.ShouldBindBodyWithJSON(&new)
	if err != nil {
		// Failed with JSON
		ctx.JSON(200, gin.H{
			"Message": "Failed with JSON!!!",
			"code":    model.ERR_JSON,
			"data":    new,
		})
	}

	// Is the new todo data has wrong
	userdata := auth.GetCurrentUser(ctx, db)
	if new.Title == "" || new.UserID == 0 || new.UserID != userdata.ID || new.ID == 0 {
		// The new todo data has wrong
		ctx.JSON(200, gin.H{
			"Message": "The new todo data has wrong!",
			"code":    model.ERR_TODO_WRONG,
			"data": gin.H{
				"new": new,
			},
		})
		return
	}

	// Update the todo data
	db.Model(&model.Todo{}).Where("user_id = ? AND id = ?", userdata.ID, new.ID).Updates(&new)
	ctx.JSON(200, gin.H{
		"Message": "Update Successfully!",
		"code":    model.SUCCESS,
		"data":    new,
	})
}

// Delete Todo
func DeleteTodo(ctx *gin.Context, db *gorm.DB) {
	// Get the todotodel ID
	todotodel := ctx.Param("todoid")

	// Get the Current User
	curuser := auth.GetCurrentUser(ctx, db)

	// Delete the todo data
	db.Model(&model.Todo{}).Where("user_id = ? AND id = ?", curuser.ID, todotodel).Delete(&model.Todo{})
	ctx.JSON(200, gin.H{
		"Message": "Delete Successfully!",
		"code":    model.SUCCESS,
		"data":    gin.H{},
	})
}

// Marked Todo
func MarkedTodo(ctx *gin.Context, db *gorm.DB) {
	// Get the todotodel ID
	todoid := ctx.Param("todoid")

	// Transform the todotodel ID to int
	markid, _ := strconv.Atoi(todoid)

	// Get the Current User
	curuser := auth.GetCurrentUser(ctx, db)

	// Marked the todo data
	db.Model(&model.Todo{}).Where("user_id = ? AND id = ?", curuser.ID, markid).Update("is_completed", true)

	ctx.JSON(200, gin.H{
		"Message": "Marked Successfully!",
		"code":    model.SUCCESS,
		"data":    gin.H{},
	})
}

// Unmarked Todo
func UnmarkedTodo(ctx *gin.Context, db *gorm.DB) {
	todoid := ctx.Param("todoid")

	// Transform the todotodel ID to int
	unmarkid, _ := strconv.Atoi(todoid)

	// Get the Current User
	curuser := auth.GetCurrentUser(ctx, db)

	// Unmarked the todo data
	db.Model(&model.Todo{}).Where("user_id = ? AND id = ?", curuser.ID, unmarkid).Update("is_completed", false)

	ctx.JSON(200, gin.H{
		"Message": "Unmarked Successfully!",
		"code":    model.SUCCESS,
		"data":    gin.H{},
	})
}
