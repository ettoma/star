package auth

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateTokenString(username string) (string, error) {

	key := []byte(os.Getenv("JWT_SECRET_KEY"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":        jwt.NewNumericDate(time.Now().Add(time.Minute * 10)),
		"authorized": true,
		"user":       username,
	})

	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return tokenString, nil

}

func ValidateToken(tokenString string) (bool, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		log.Print(tokenString)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		key := []byte(os.Getenv("SECRET_KEY"))
		return key, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		var date time.Time
		switch exp := claims["exp"].(type) {
		case float64:
			date = time.Unix(int64(exp), 0)
		case json.Number:
			v, _ := exp.Int64()
			date = time.Unix(v, 0)
		}
		fmt.Printf(" username: %s \n authorised: %v \n expiresAt: %s \n", claims["user"], claims["authorized"], date)
		return true, nil
	} else {
		return false, err
	}
}
