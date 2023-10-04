package repository

import (
	"accountservice/app/domain/dao"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *dao.User) (dao.User, error)
	UpdateUser(email string) (dao.User, error)
	GetUser(email string) error
	DeleteUser(email string) error
}

type UserRepositoryImpl struct {
	db *gorm.DB
}
