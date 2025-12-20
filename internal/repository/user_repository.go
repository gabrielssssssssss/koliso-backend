package repository

import (
	"github.com/gabrielssssssssss/koliso-backend.git/internal/entity"
	"github.com/supabase-community/supabase-go"
)

type UserRepository interface {
	Register(users *entity.UserEntity) ([]entity.UserEntity, error)
}

type userImplementation struct {
	db *supabase.Client
}

func NewUserRepository(client *supabase.Client) UserRepository {
	return &userImplementation{
		db: client,
	}
}
