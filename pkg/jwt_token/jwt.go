package jwt_token

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kooroshh/fiber-boostrap/pkg/env"
)

type ClaimToken struct {
	Username string `json:"username"`
	FullName string `json:"full_name"`
	jwt.RegisteredClaims
}

var MapTypeToken = map[string]time.Duration{
	"token": time.Hour * 3,
	"refresh_token": time.Hour * 72,
}

func GenerateToken(ctx context.Context, username string, fullname string, tokenType string) (string, error) {
	secret := []byte(env.GetEnv("APP_SECRET", ""))

	claimToken := ClaimToken{
		Username: username,
		FullName: fullname,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: env.GetEnv("APP_NAME", ""),
			IssuedAt: jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(MapTypeToken[tokenType])),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimToken)

	resultToken, err := token.SignedString(secret)
	if err != nil {
		return resultToken, fmt.Errorf("failed to generate toke: %v", err)
	}

	return resultToken, nil
}