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
			"code":    model.ERR_JSON,
			"data":    gin.H{},
		})
	} else {
		// Check the user data
		if userdata.Username == "" || userdata.Password == "" || userdata.Email == "" || userdata.Nickname == "" {
			ctx.JSON(200, gin.H{
				"Message": "The user data is blank!",
				"code":    model.ERR_JSON,
				"data":    gin.H{},
			})
			return
		}

		// is the user exist
		var tempuser model.User
		db.Where("username = ?", userdata.Username).First(&tempuser)
		if tempuser.ID != 0 {
			ctx.JSON(200, gin.H{
				"Message": "The user is exist!",
				"code":    model.ERR_USERNAME_EXIST,
				"data":    gin.H{},
			})
			return
		}
		db.Where("email = ?", userdata.Email).First(&tempuser)
		if tempuser.ID != 0 {
			ctx.JSON(200, gin.H{
				"Message": "The email is exist!",
				"code":    model.ERR_EMAIL_EXIST,
				"data":    gin.H{},
			})
			return
		}

		db.Create(&userdata)

		ctx.JSON(200, gin.H{
			"Message": "Create Successfully!",
			"code":    model.SUCCESS,
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
		ctx.JSON(400, gin.H{
			"Message": "User not found!",
			"code":    model.ERR_USER_NOT_FOUND,
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
				"code":    model.SUCCESS,
				"data": gin.H{
					"username": tempuser.Username,
					"nickname": tempuser.Nickname,
					"email":    tempuser.Email,
					"token":    tokenString, // Send the token to the client
				},
			})

		} else {
			// Password is wrong
			ctx.JSON(400, gin.H{
				"Message": "Oh no! Password is wrong!",
				"code":    model.ERR_PASSWORD_WRONG,
				"data":    userdata,
			})
		}
	}
}

// Get a user data
func GetUserInfo(ctx *gin.Context, db *gorm.DB) {
	// Get the logined user data
	userdata := auth.GetCurrentUser(ctx, db)

	// return the user data
	ctx.JSON(200, gin.H{
		"Message": "Get Successfully!",
		"code":    model.SUCCESS,
		"data": gin.H{
			"userid":   userdata.ID,
			"username": userdata.Username,
			"nickname": userdata.Nickname,
			"email":    userdata.Email,
		},
	})
}

// Update a user
func UpdateUsername(ctx *gin.Context, db *gorm.DB) {
	// Get the new username
	var new model.UserToUpdate
	err := ctx.ShouldBindBodyWithJSON(&new)
	if err != nil {
		ctx.JSON(200, gin.H{
			"Message": "Failed with JSON!!!",
			"code":    model.ERR_JSON,
			"data":    gin.H{},
		})
		panic(err)
	}

	// Is new username is blank
	if new.Username == "" {
		// The new username is blank
		ctx.JSON(200, gin.H{
			"Message": "The new username is blank!",
			"code":    model.ERR_USERNAME_BLANK,
			"data":    gin.H{},
		})
		return
	}

	// Is new username is the same
	curname, _ := ctx.Get("username")
	if new.Username == curname {
		// The new username is the same
		ctx.JSON(200, gin.H{
			"Message": "The new username is the same!",
			"code":    model.ERR_USERNAME_SAME,
			"data":    gin.H{},
		})
		return
	}

	// Is the new username has existed
	var tempuser model.User
	db.Where("username = ?", new.Username).First(&tempuser)
	if tempuser.ID != 0 {
		// The new username is exist
		ctx.JSON(200, gin.H{
			"Message": "The new username is exist!",
			"code":    model.ERR_USERNAME_EXIST,
			"data":    gin.H{},
		})
		return
	}

	// Update the username
	db.Model(&model.User{}).Where("username = ?", curname).Update("username", new.Username)
	ctx.JSON(200, gin.H{
		"Message": "Update Successfully!",
		"code":    model.SUCCESS,
		"data": gin.H{
			"OldUsername": curname,
			"NewUsername": new.Username,
		},
	})

	// Then the client need to re-login!!!
}

func UpdatePassword(ctx *gin.Context, db *gorm.DB) {
	// Get the new password
	var new model.UserToUpdate
	err := ctx.ShouldBindBodyWithJSON(&new)
	if err != nil {
		ctx.JSON(200, gin.H{
			"Message": "Failed with JSON!!!",
			"code":    model.ERR_JSON,
			"data":    gin.H{},
		})
		panic(err)
	}

	// Is new password is blank
	if new.Password == "" {
		// The new password is blank
		ctx.JSON(200, gin.H{
			"Message": "The new password is blank!",
			"code":    model.ERR_PASSWORD_BLANK,
			"data":    gin.H{},
		})
		return
	}

	// Is new password
	tempuser := auth.GetCurrentUser(ctx, db)
	if new.Password == tempuser.Password {
		// The new password is the same
		ctx.JSON(200, gin.H{
			"Message": "The new password is the same!",
			"code":    model.ERR_PASSWORD_SAME,
			"data":    gin.H{},
		})
		return
	}

	// Update the password
	db.Model(&model.User{}).Where("username = ?", tempuser.Username).Update("password", new.Password)
	ctx.JSON(200, gin.H{
		"Message": "Update Successfully!",
		"code":    model.SUCCESS,
		"data": gin.H{
			"NewPassword": new.Password,
		},
	})

	// Then the client need to re-login!!!
}

func UpdateEmail(ctx *gin.Context, db *gorm.DB) {
	// Get the new email
	var new model.UserToUpdate
	err := ctx.ShouldBindBodyWithJSON(&new)
	if err != nil {
		panic(err)
	}

	// Is new email is blank
	if new.Email == "" {
		// The new email is blank
		ctx.JSON(200, gin.H{
			"Message": "The new email is blank!",
			"code":    model.ERR_EMAIL_BLANK,
			"data":    gin.H{},
		})
		return
	}

	// Is new email exist or the same
	tempuser := auth.GetCurrentUser(ctx, db)
	if new.Email == tempuser.Email {
		// The new email is the same
		ctx.JSON(200, gin.H{
			"Message": "The new email is the same!",
			"code":    model.ERR_EMAIL_SAME,
			"data":    gin.H{},
		})
		return
	}

	// Is the new email has existed
	var tempuser2 model.User
	db.Where("email = ?", new.Email).First(&tempuser2)
	if tempuser2.ID != 0 {
		// The new email is exist
		ctx.JSON(200, gin.H{
			"Message": "The new email is exist!",
			"code":    model.ERR_EMAIL_EXIST,
			"data":    gin.H{},
		})
		return
	}

	// Update the email
	db.Model(&model.User{}).Where("username = ?", tempuser.Username).Update("email", new.Email)
	ctx.JSON(200, gin.H{
		"Message": "Update Successfully!",
		"code":    model.SUCCESS,
		"data": gin.H{
			"NewEmail": new.Email,
		},
	})
}

func UpdateNickname(ctx *gin.Context, db *gorm.DB) {
	// Get the new nickname
	var new model.UserToUpdate
	err := ctx.ShouldBindBodyWithJSON(&new)
	if err != nil {
		panic(err)
	}

	// Is new nickname is blank
	if new.Nickname == "" {
		// The new nickname is blank
		ctx.JSON(200, gin.H{
			"Message": "The new nickname is blank!",
			"code":    model.ERR_NICKNAME_BLANK,
			"data":    gin.H{},
		})
		return
	} else {
		// Update the nickname
		curname, _ := ctx.Get("username")
		db.Model(&model.User{}).Where("username = ?", curname).Update("nickname", new.Nickname)
		ctx.JSON(200, gin.H{
			"Message": "Update Successfully!",
			"code":    model.SUCCESS,
			"data":    new.Nickname,
		})
	}
}

// Delete a user
func DeleteAccount(ctx *gin.Context, db *gorm.DB) {
	// Get the delete detail
	var deletedata model.DeleteDetail
	err := ctx.ShouldBindBodyWithJSON(&deletedata)
	if err != nil {
		panic(err)
	}

	// Get the username
	username := deletedata.Usertodel
	// Check the user
	curuser, exist := ctx.Get("username")
	// Get the UserID from the username
	var tempuser model.User
	db.Model(&model.User{}).Where("username = ?", curuser).First(&tempuser)
	userid := tempuser.ID
	if !exist {
		// Current user is not exist
		ctx.JSON(200, gin.H{
			"Message": "The user is not exist!",
			"code":    model.ERR_USER_NOT_FOUND,
			"data":    gin.H{},
		})
		return
	}

	if username != curuser {
		// The user is not the same
		ctx.JSON(200, gin.H{
			"Message": "The user is not the same!",
			"code":    model.ERR_USER_NOT_SAME,
			"data":    gin.H{},
		})
		return
	} else {
		// Delete the user
		db.Model(&model.User{}).Where("username = ?", username).Delete(&model.User{})

		// Whether to delete the data
		if deletedata.Deldata {
			// Delete the ALL TODDO By UserID
			db.Model(&model.Todo{}).Where("user_id = ?", userid).Delete(&model.Todo{})
		}

		ctx.JSON(200, gin.H{
			"Message": "Delete Successfully!",
			"code":    model.SUCCESS,
			"data": gin.H{
				"DeletedUser": username,
				"DelData":     deletedata.Deldata,
			},
		})
	}
}
