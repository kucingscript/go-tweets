package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

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

func ValidateToken(tokenString string, secretKey string) (*Claims, error) {
	claims := &Claims{}
	key := []byte(secretKey)

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
