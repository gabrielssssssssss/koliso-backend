package helper

import (
	"fmt"
	"time"

	"github.com/gabrielssssssssss/koliso-backend.git/internal/middlewares"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateJwtToken(uuid, email, role string) (string, error) {
	expiration := time.Now().Add(time.Minute * 300)
	claims := middlewares.JWTClaim{
		Uuid:  uuid,
		Email: email,
		Role:  role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "github.com/gabrielssssssssss",
			ExpiresAt: jwt.NewNumericDate(expiration),
		},
	}

	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tokenAlgo.SignedString(middlewares.JWT_SECRET_KEY)
	if err != nil {
		return "", fmt.Errorf("Unable to create JWT with the given parameters.")
	}
	return token, nil
}
