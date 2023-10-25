package auth

import (
	"Gin/Basics/configs"
	"time"

	"github.com/golang-jwt/jwt"
)

var Key = []byte(configs.JWT_SECRET())

func GenerateJWT(email string) (tokenStr string, err error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(30 * time.Minute).Unix()
	claims["authorized"] = true
	claims["user"] = "username"

	tokenStr, err = token.SignedString(Key)

	return tokenStr, err
}

func ValidateJWT(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return Key, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
