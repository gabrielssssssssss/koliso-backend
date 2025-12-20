package service

import (
	"github.com/gabrielssssssssss/koliso-backend.git/internal/entity"
	"github.com/gabrielssssssssss/koliso-backend.git/internal/model"
)

func (user userServiceimpl) Register(request *model.RegisterUserPayload) (*model.RegisterUserResponse, error) {

	input := entity.UserEntity{
		FirstName:     request.FirstName,
		LastName:      request.LastName,
		Email:         request.Email,
		Password:      request.Password,
		Phone:         request.Phone,
		Bio:           request.Bio,
		ProfilPicture: request.ProfilPicture,
	}

	user.userRepository.Register(&input)
	response := model.RegisterUserResponse{
		Uuid:  "",
		Role:  "",
		Token: "",
	}

	return &response, nil
}
