package user

import (
	"github.com/gin-gonic/gin"
)

func (controller *UserController) Route(app *gin.Engine) {
	app.POST("/register", controller.Register)
}
