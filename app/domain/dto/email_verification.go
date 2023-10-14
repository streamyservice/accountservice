package dto

type EmailVerification struct {
	Code  string `json:"code" binding:"required"`
	Email string `json:"email" binding:"required"`
}
