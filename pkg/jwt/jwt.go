package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(id int64, email, secretKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    id,
		"email": email,
		"exp":   time.Now().Add(60 * time.Minute).Unix(),
	})

	key := []byte(secretKey)
	tokenStr, err := token.SignedString(key)

	return tokenStr, err
}
