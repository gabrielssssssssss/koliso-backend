package service

import "github.com/gabrielssssssssss/koliso-backend.git/internal/model"

func (user userServiceimpl) Register(request *model.RegisterUserPayload) (*model.RegisterUserResponse, error) {
	response := model.RegisterUserResponse{
		Uuid:  "",
		Role:  "",
		Token: "",
	}

	return &response, nil
}
