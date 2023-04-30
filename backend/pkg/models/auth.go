package models

type RegisterUserRequest struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	RegisterType string `json:"register_type"` // EMAIL,GOOGLE,PHONE
}

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserWithAuth struct {
	User         *User  `json:"user"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type GetByCredentialsRequest struct {
	Email    string `json:"phone"`
	Password string `json:"password"`
}
