package main

import (
	"github.com/gabrielssssssssss/koliso-backend.git/internal/controller"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	controller.Controller()
}
