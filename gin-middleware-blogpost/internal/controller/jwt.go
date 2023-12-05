package controller

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"time"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env file: %v", err)
	}
}

// GenerateToken generates jwt token
func GenerateToken(userID int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Subject:   fmt.Sprintf("%d", userID),
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
	})
	var jwtSect = os.Getenv("PASSWORD_SALT")
	log.Println("config.JwtSecret=", jwtSect)
	tokenString, err := token.SignedString([]byte(jwtSect))
	return tokenString, err
}

// ParseToken - parses token
func ParseToken(accessToken string) (int, error) {
	var jwtSect = os.Getenv("PASSWORD_SALT")
	t, err := jwt.ParseWithClaims(accessToken, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSect), nil
	})
	if err != nil {
		log.Println("err := jwt.ParseWithClaims ==== ", err)
		return 0, err
	}

	if !t.Valid {
		return 0, errors.New("invalid token")
	}

	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid claims")
	}

	subject, ok := claims["sub"].(string)
	if !ok {
		return 0, errors.New("invalid subject")
	}

	userId, err := strconv.Atoi(subject)
	log.Println(userId)
	if err != nil {
		return 0, errors.New("invalid subject")
	}
	log.Println("parseToken - userId=", userId)
	return userId, nil

}
