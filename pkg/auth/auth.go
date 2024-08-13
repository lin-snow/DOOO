package auth

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/lin-snow/dooo/internal/model"
	"gorm.io/gorm"
)

// Create claims
func CreateClaims(user model.User) jwt.Claims {
	claim := model.MyClaims{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "lin-snow",
			Subject:   "auth",
			Audience:  jwt.ClaimStrings{"dooo"},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	return claim
}

// Create Token (The token will be send to the client)
func GenerateToken(claim jwt.Claims) (string, error) {
	/*
		NewWithClaims contains Header & Payload
	*/
	part := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	/*
		SignedString contains Generate Signature
		and Link the Header & Payload & Signature together
		and return the tokenString
	*/
	token, err := part.SignedString(model.MySigningKey)
	if err != nil {
		return "", err
	}
	return token, nil
}

// Parse Token (The token will be received from the client)
func ParseToken(token string) (*model.MyClaims, error) {

	/*
		Parse the token to get the claims.
		The process contains getting the signature from the token
	*/
	parsedtoken, err := jwt.ParseWithClaims(token, &model.MyClaims{}, func(parsedtoken *jwt.Token) (interface{}, error) {
		return model.MySigningKey, nil
	})

	// Check the token
	if err != nil { // If the token is invalid
		log.Fatal(err)
	} else if claims, ok := parsedtoken.Claims.(*model.MyClaims); ok { // verify the signature
		fmt.Println(claims.Username, claims.RegisteredClaims.Issuer)
		return claims, nil
	} else {
		log.Fatal("unknown claims type, cannot proceed") // If the token is invalid
	}

	return nil, nil
}

// JWT Auth Middleware
func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get the Authorization from the header
		/*
			auth is token and contains the Header & Payload & Signature

		*/
		auth := ctx.Request.Header.Get("Authorization")

		// Check the auth
		if auth == "" {
			// auth is blank
			ctx.JSON(http.StatusOK, gin.H{
				"Message": "The Auth is Blank!!!!",
				"code":    model.ERR_AUTH_NOT_FOUND,
				"data":    gin.H{},
			})

			// Stop the request
			ctx.Abort()
			return
		} else {
			// Authorization is not blank
			parts := strings.SplitN(auth, " ", 2)
			if !(len(parts) == 2 && parts[0] == "Bearer") {
				ctx.JSON(http.StatusOK, gin.H{
					"Message": "The Auth is Invalid!!!!",
					"code":    model.ERR_AUTH_INVALID,
					"data":    gin.H{},
				})
				ctx.Abort()
				return
			}

			// Parse the token
			mc, err := ParseToken(parts[1])
			if err != nil {
				ctx.JSON(http.StatusOK, gin.H{
					"Message": "The Token is Invalid!!!!",
					"code":    model.ERR_TOKEN_INVALID,
					"data":    gin.H{},
				})
				ctx.Abort()
				return
			}

			// Auth is valid
			// Set the claims to the context
			ctx.Set("username", mc.Username)
			ctx.Next()
		}
	}
}

// GET CURRENT USER
func GetCurrentUser(ctx *gin.Context, db *gorm.DB) model.User {
	// Get the current user
	curuser, _ := ctx.Get("username")
	// Get the current user data
	var userdata model.User
	db.Model(&model.User{}).Where("username = ?", curuser).First(&userdata)

	// Return the current user data
	return userdata
}
