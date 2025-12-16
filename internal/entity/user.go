package entity

import "github.com/golang/protobuf/ptypes/timestamp"

type UserEntity struct {
	Uuid               string
	FirstName          string
	LastName           string
	Email              string
	Password           string
	Phone              string
	Bio                string
	ProfilPicture      string
	Role               string
	Average_rating     float32
	IsEmailVerified    bool
	IsPhoneVerified    bool
	IsIdentityVerified bool
	StripeCustomerId   string
	StripeAccountId    string
	CreatedAt          timestamp.Timestamp
}
