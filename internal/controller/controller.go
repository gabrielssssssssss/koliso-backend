package controller

import (
	"github.com/gabrielssssssssss/koliso-backend.git/config"
	"github.com/gabrielssssssssss/koliso-backend.git/internal/controller/user"
	"github.com/gabrielssssssssss/koliso-backend.git/internal/repository"
	"github.com/gabrielssssssssss/koliso-backend.git/internal/service"
	"github.com/gin-gonic/gin"
)

func Controller() {
	databaseSupabase := config.NewSupabaseClient()

	userRepository := repository.NewUserRepository(databaseSupabase)
	userService := service.NewUserService(&userRepository)
	userController := user.NewUserController(&userService)

	app := gin.Default()
	userController.Route(app)
	app.Run(":8080")
}
