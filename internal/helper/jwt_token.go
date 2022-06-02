package helper

import (
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateToken(id int, email, secret string) (string, error) {
	claims := jwt.MapClaims{}
	claims["exp"] = time.Now().Add((6 * time.Hour)).Unix()
	claims["iat"] = time.Now().Unix()
	claims["email"] = email
	claims["id"] = id
	claims["aud"] = "rahano.service"

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
