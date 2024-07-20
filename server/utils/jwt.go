package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "incrediblySuperSecret"

func CreateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"expiry": time.Now().Add(time.Hour * 2).Unix(),
	}) //new token with data

	return token.SignedString([]byte(secretKey))
}
