package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func VerifyJwtToken(c *gin.Context) {
	c.Next()
	jwtToken := c.Request.Header.Get("JWT_TOKEN")
	fmt.Println(jwtToken)
}
