/*
user package contains the user model and handler
*/
package user

import (
	"github.com/gin-gonic/gin"
	"github.com/lin-snow/dooo/internal/model"
	"github.com/lin-snow/dooo/pkg/auth"
	"gorm.io/gorm"
)

// Create a new user
func Register(ctx *gin.Context, db *gorm.DB) {
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

// Login a user
func Login(ctx *gin.Context, db *gorm.DB) {
	// Get the user data
	var userdata model.User
	err := ctx.ShouldBindBodyWithJSON(&userdata)
	if err != nil {
		panic(err)
	}

	// Check the user data
	var tempuser model.User
	db.Where("username = ?", userdata.Username).First(&tempuser)
	if tempuser.ID == 0 {
		// User not found
		ctx.JSON(200, gin.H{
			"Message": "User not found!",
			"code":    400,
			"data":    gin.H{},
		})
	} else {
		// Check the password
		if userdata.Password == tempuser.Password {
			//  Password is correct
			// Create Claim and Generate Token
			tokenString, _ := auth.GenerateToken(auth.CreateClaims(tempuser))

			// Login Successfully && Return the token
			ctx.JSON(200, gin.H{
				"Message": "Login Successfully!",
				"code":    200,
				"data": gin.H{
					"token": tokenString, // Send the token to the client
				},
			})

		} else {
			// Password is wrong
			ctx.JSON(200, gin.H{
				"Message": "Oh no! Password is wrong!",
				"code":    400,
				"data":    userdata,
			})
		}
	}
}
