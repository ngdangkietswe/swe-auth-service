package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/ngdangkietswe/swe-auth-service/configs"
	"github.com/ngdangkietswe/swe-auth-service/data/ent"
	"time"
)

type JwtUserClaims struct {
	UserId   string `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// GenerateToken is a function that generates a JWT token.
func GenerateToken(entUser *ent.User, isRefresh bool) (string, error) {
	var tokenExp time.Duration
	if isRefresh {
		tokenExp = time.Second * configs.GlobalConfig.RefreshTokenExp
	} else {
		tokenExp = time.Second * configs.GlobalConfig.JwtExp
	}

	exp := time.Now().Add(tokenExp).Unix()

	mapClaims := make(jwt.MapClaims)
	mapClaims["sub"] = entUser.ID
	mapClaims["user"] = JwtUserClaims{
		UserId:   entUser.ID.String(),
		Username: entUser.Username,
		Email:    entUser.Email,
	}
	mapClaims["iat"] = time.Now().Unix()
	mapClaims["nbf"] = time.Now().Unix()
	mapClaims["exp"] = exp

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims).SignedString([]byte(configs.GlobalConfig.JwtSecret))

	if err != nil {
		return "", err
	}

	return token, nil
}

// ParseToken is a function that parses a JWT token.
func ParseToken(jwtToken string) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(configs.GlobalConfig.JwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !token.Valid || !ok {
		return nil, err
	}

	return &claims, nil
}
