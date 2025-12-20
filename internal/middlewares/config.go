package middlewares

import (
	"fmt"

	"github.com/golang-jwt/jwt/v4"

	"github.com/gin-gonic/gin"
)

func VerifyJwtToken(c *gin.Context) {
	c.Next()
	jwtToken := c.Request.Header.Get("JWT_TOKEN")
	token, _ := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return JWT_SECRET_KEY, nil
	})

	if token.Valid == true {
		c.JSON(401, fmt.Errorf("JWT validation error: signature verification failed, invalid signing key"))
	}

}
