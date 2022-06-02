package service

import (
	"github.com/alexeymyakinin/ruck/app/src/core/env"
	"github.com/golang-jwt/jwt"
	"time"
)

type JWTClaims struct {
	jwt.StandardClaims
	ID       uint64 `json:"id"`
	Username string `json:"username"`
}

func GenerateJWT(userId uint64, username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, &JWTClaims{
		StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Add(24 * time.Hour).Unix()},
		ID:             userId,
		Username:       username,
	})

	return token.SignedString([]byte(env.JWTSecret))
}
