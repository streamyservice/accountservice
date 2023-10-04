package dao

import "gorm.io/gorm"

type User struct {
	Email         string `json:"email"`
	Username      string `json:"username"`
	Fullname      string `json:"fullname"`
	LastIp        string `json:"last_ip"`
	EmailVerified bool   `json:"email_verified"`
	Password      string `json:"password"`
	BirthDate     string `json:"birthDate"`
	Image         string `json:"image"`
	Gender        string `json:"gender"`
	gorm.Model
}
