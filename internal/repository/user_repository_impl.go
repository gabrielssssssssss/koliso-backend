package repository

import (
	"encoding/json"

	"github.com/gabrielssssssssss/koliso-backend.git/config"
	"github.com/gabrielssssssssss/koliso-backend.git/internal/entity"
)

func (user *userImplementation) Register(users *entity.UserEntity) ([]entity.UserEntity, error) {
	_, cancel := config.NewSupabaseContext()
	defer cancel()

	var Results []entity.UserEntity

	response, _, err := user.db.From("Users").Insert(map[string]interface{}{
		"first_name":          users.FirstName,
		"last_name":           users.LastName,
		"email":               users.Email,
		"password":            users.Password,
		"phone":               users.Phone,
		"profile_picture_url": users.ProfilPicture,
	}, false, "", "representation", "").Execute()

	if err != nil {
		return Results, err
	}

	err = json.Unmarshal(response, &Results)
	if err != nil {
		return Results, err
	}
	return Results, nil
}
