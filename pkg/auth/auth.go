package auth

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lin-snow/dooo/internal/model"
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

// Create Token
func GenerateToken(claim jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString(model.MySigningKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// Parse Token
func ParseToken(claim jwt.Claims, tokenString string) {
	token, err := jwt.ParseWithClaims(tokenString, claim, func(token *jwt.Token) (interface{}, error) {
		return model.MySigningKey, nil
	})
	if err != nil {
		log.Fatal(err)
	} else if claims, ok := token.Claims.(*model.MyClaims); ok {
		fmt.Println(claims.Username, claims.RegisteredClaims.Issuer)
	} else {
		log.Fatal("unknown claims type, cannot proceed")
	}
}
