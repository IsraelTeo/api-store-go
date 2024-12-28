package dto

type UserResponse struct {
	Email   string `json:"email"`
	IsAdmin bool   `json:"is_admin"`
}
