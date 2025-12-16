package repository

import (
	"github.com/gabrielssssssssss/koliso-backend.git/config"
	"github.com/gabrielssssssssss/koliso-backend.git/internal/entity"
)

func (user *userImplementation) Register(users *entity.UserEntity) error {
	_, cancel := config.NewSupabaseContext()
	defer cancel()
	// Supabase requests to add users here
	return nil
}
