package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const tokenDuration int64 = 120 // token expiration time in minutes
const secretKey = "supersecret"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Minute * time.Duration(tokenDuration)).Unix(),
	})
	return token.SignedString([]byte(secretKey))
}
