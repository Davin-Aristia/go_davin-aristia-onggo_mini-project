package middleware

import (
	"time"

	"github.com/golang-jwt/jwt"
	
	"go-mini-project/config"
)

func CreateToken(userId int, name, email, role string) (string, error) {
	claims := jwt.MapClaims{}
	// token kedua (payload)
	claims["authorized"] = true
	claims["userId"] = userId
	claims["name"] = name
	claims["email"] = email
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //Token expires after 1 hour
	// token pertama (header)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// return bersama token ketiga (dengan secret key)
	return token.SignedString([]byte(config.JWT_KEY))
}