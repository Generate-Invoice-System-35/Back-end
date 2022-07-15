package helper

import (
	"log"
	"time"

	"Back-end/config"

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

func ExtractToken(tokenString string) (jwt.MapClaims, bool) {
	data := config.Config{}
	secretString := data.JWT_KEY
	secret := []byte(secretString)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		log.Printf("Invalid JWT Token")
		return nil, false
	}
}
