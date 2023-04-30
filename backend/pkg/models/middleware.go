package models

type HasAccessModel struct {
	Email     string `json:"email"`
	UserId    string `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
