package utils

import (
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
func GenerateToken(entUser *ent.User) (string, error) {
	exp := time.Now().Add(time.Second * configs.GlobalConfig.JwtExp).Unix()

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
