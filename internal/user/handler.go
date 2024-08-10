/*
user package contains the user model and handler
*/
package user

import (
	"github.com/gin-gonic/gin"
	"github.com/lin-snow/dooo/internal/model"
	"gorm.io/gorm"
)

// Create a new user
func CreateUser(ctx *gin.Context, db *gorm.DB) {
	var userdata model.User
	err := ctx.ShouldBindBodyWithJSON(&userdata)
	if err != nil {
		ctx.JSON(200, gin.H{
			"Message": "Failed with JSON!",
			"code":    400,
			"data":    gin.H{},
		})
	} else {
		db.Create(&userdata)

		ctx.JSON(200, gin.H{
			"Message": "Create Successfully!",
			"code":    200,
			"data":    userdata,
		})
	}
}
