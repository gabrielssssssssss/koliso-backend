package repository

import (
	"fmt"

	"github.com/gabrielssssssssss/koliso-backend.git/config"
	"github.com/gabrielssssssssss/koliso-backend.git/internal/entity"
)

func (user *userImplementation) Register(users *entity.UserEntity) error {
	_, cancel := config.NewSupabaseContext()
	defer cancel()

	fmt.Println(users)
	data, _, err := user.db.From("Users").Insert(map[string]interface{}{
		"first_name":          users.FirstName,
		"last_name":           users.LastName,
		"email":               users.Email,
		"password":            users.Password,
		"phone":               users.Phone,
		"profile_picture_url": users.ProfilPicture,
	}, false, "", "representation", "").Execute()

	fmt.Println(err)
	fmt.Println(data)
	return nil
}
