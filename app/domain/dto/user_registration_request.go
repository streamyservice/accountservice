package dto

type UserRegistrationRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Fullname string `json:"fullname" binding:"required"`
	Password string `json:"password" binding:"required"`
}
