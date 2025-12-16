package model

type RegisterUserPayload struct {
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	Phone         string `json:"phone"`
	Bio           string `json:"bio"`
	ProfilPicture string `json:"profile_picture"`
}

type RegisterUserResponse struct {
	Uuid  string `json:"uuid"`
	Role  string `json:"role"`
	Token string `json:"token"`
}
