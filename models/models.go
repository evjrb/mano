package models

type User struct {
	UserId      string `json:"user_id"`
	UserToken   string `json:"user_token"`
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
}
