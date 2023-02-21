package jwtManager

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/kkomissarov/beggar/db"
	"github.com/kkomissarov/beggar/models"
	"os"
	"time"
)

func decodeToken(token *jwt.Token) (interface{}, error) {
	_, ok := token.Method.(*jwt.SigningMethodHMAC)
	if !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
	return []byte(os.Getenv("JWT_SECRET")), nil
}

func findUser(claims jwt.MapClaims) *models.User {
	userID := claims["sub"]
	var user models.User
	db.DB.First(&user, userID)
	return &user
}

func tokenExpired(claims jwt.MapClaims) bool {
	return float64(time.Now().Unix()) > claims["exp"].(float64)
}

func tokenRevoked(token string) bool {
	var revokedToken models.RevokedToken
	db.DB.
		Where("token = ?", token).
		Where("expired_at > ?", time.Now()).
		First(&revokedToken)
	return revokedToken.ID != 0
}

func ExtractClaims(token string) (jwt.MapClaims, error) {
	tokenObject, err := jwt.Parse(token, decodeToken)
	if err != nil {
		return nil, err
	}

	claims, ok := tokenObject.Claims.(jwt.MapClaims)
	if !ok || !tokenObject.Valid {
		return nil, errors.New("auth token is invalid")
	}

	return claims, nil
}

func FindUserByToken(token string) (*models.User, error) {
	if tokenRevoked(token) {
		return nil, errors.New("auth token is revoked")
	}

	claims, err := ExtractClaims(token)
	if err != nil {
		return nil, err
	}

	if tokenExpired(claims) {
		return nil, errors.New("auth token is expired")
	}

	user := findUser(claims)
	if user.ID == 0 {
		return nil, errors.New("user specified in auth token not found")
	}

	return user, nil
}
