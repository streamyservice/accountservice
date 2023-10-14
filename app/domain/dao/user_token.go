package dao

import (
	"gorm.io/gorm"
	"time"
)

type UserToken struct {
	Id    uint      `json:"id"`
	Email string    `json:"email"`
	Exp   time.Time `json:"exp"`
	Code  string    `json:"code"`

	gorm.Model
}
