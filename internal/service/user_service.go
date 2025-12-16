package service

import (
	"github.com/gabrielssssssssss/koliso-backend.git/internal/model"
	"github.com/gabrielssssssssss/koliso-backend.git/internal/repository"
)

type UserService interface {
	Register(users *model.RegisterUserPayload) (*model.RegisterUserResponse, error)
}

type userServiceimpl struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) UserService {
	return &userServiceimpl{
		userRepository: *userRepository,
	}
}
