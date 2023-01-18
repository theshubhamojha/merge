package utils

import (
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/megre/dto"
)

func GenerateJWTToken(email string, role dto.RoleType, accountID string, jwtSecret string, expiry int) (token string, err error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":      email,
		"role":       role,
		"account_id": accountID,
		"iat":        time.Now().Unix(),
		"expiry":     time.Now().Add(time.Minute * time.Duration(expiry)),
	})

	secret := []byte(jwtSecret)

	token, err = claims.SignedString(secret)
	if err != nil {
		return
	}

	return
}

func VerifyJWTToken(token string, jwtSecret string) (jwt.MapClaims, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (data interface{}, jwtErr error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			jwtErr = errors.New("unexpected jwt method")
			return
		}
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		return claims, nil
	}

	return nil, errors.New("error verifying jwt token")
}
