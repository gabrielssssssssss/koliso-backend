package user

import (
	"github.com/gabrielssssssssss/koliso-backend.git/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func (controller *UserController) Route(app *gin.Engine) {
	app.Use(middlewares.VerifyJwtToken)
	app.POST("/register", controller.Register)
}
