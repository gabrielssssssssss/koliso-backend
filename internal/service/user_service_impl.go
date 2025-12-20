package service

import (
	"github.com/gabrielssssssssss/koliso-backend.git/internal/entity"
	"github.com/gabrielssssssssss/koliso-backend.git/internal/model"
	"github.com/gabrielssssssssss/koliso-backend.git/internal/validation"
)

func (user userServiceimpl) Register(request *model.RegisterUserPayload) (*model.RegisterUserResponse, error) {
	email, err := validation.EmailAdress(request.Email)
	if !email && err != nil {
		return nil, err
	}

	phone, err := validation.PhoneNumber(request.Phone)
	if !phone && err != nil {
		return nil, err
	}

	input := entity.UserEntity{
		FirstName:     request.FirstName,
		LastName:      request.LastName,
		Email:         request.Email,
		Password:      request.Password,
		Phone:         request.Phone,
		Bio:           request.Bio,
		ProfilPicture: request.ProfilPicture,
	}

	userEntity, err := user.userRepository.Register(&input)
	if err != nil {
		return nil, err
	}

	response := model.RegisterUserResponse{
		Uuid:  userEntity[0].Uuid,
		Role:  userEntity[0].Role,
		Token: "",
	}
	return &response, nil
}
