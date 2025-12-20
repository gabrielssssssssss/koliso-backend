package middlewares

import (
	"os"

	"github.com/golang-jwt/jwt/v4"
)

var (
	JWT_SECRET_KEY = []byte(os.Getenv("JWT_KEY"))
)

type JWTClaim struct {
	Uuid  string
	Email string
	Role  string
	jwt.RegisteredClaims
}

func GetRole(claims *JWTClaim) string {
	return claims.Role
}

func GetClaims(claims *JWTClaim) *JWTClaim {
	return claims
}
