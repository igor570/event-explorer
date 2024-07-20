package utils

import (
	"errors"
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

func VerifyToken(token string) error {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok { //checking type signing method of token
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return errors.New("could not parse the token")
	}

	tokenIsValid := parsedToken.Valid

	if !tokenIsValid {
		return errors.New("invalid token")
	}

	/*
		//optional valdiation of properties in the token: email and user id

		claims, ok := parsedToken.Claims.(jwt.MapClaims) //checking the type of parsedToken

		if !ok {
			return errors.New("Invalid token claims")
		}

		//accessing the data in a map
		email := claims["email"].(string)
		userId := claims["userId"].(int64)

	*/

	return nil

}
