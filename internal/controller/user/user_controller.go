package user

import (
	"net/http"

	"github.com/gabrielssssssssss/koliso-backend.git/internal/model"
	"github.com/gabrielssssssssss/koliso-backend.git/internal/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(UserService *service.UserService) UserController {
	return UserController{UserService: *UserService}
}

func (controller *UserController) Register(c *gin.Context) {
	var request model.RegisterUserPayload

	err := c.ShouldBindBodyWithJSON(&request)
	if err != nil {
		return
	}

	response, err := controller.UserService.Register(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 400,
			"msg":  "Bad request",
			"data": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "OK",
		"data": response,
	})
}
