package repository

import (
	"accountservice/app/domain/dao"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *dao.User) (*dao.User, error)
	UpdateUser(email string, user *dao.User) (*dao.User, error)
	GetUser(email string) (*dao.User, error)
	DeleteUser(email string) error
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func (u UserRepositoryImpl) CreateUser(user *dao.User) (*dao.User, error) {
	return nil, nil
}

func (u UserRepositoryImpl) UpdateUser(email string, user *dao.User) (*dao.User, error) {
	return nil, nil
}

func (u UserRepositoryImpl) GetUser(email string) (*dao.User, error) {
	return nil, nil
}

func (u UserRepositoryImpl) DeleteUser(emil string) error {
	return nil
}

func UserRepositoryInit(db *gorm.DB) *UserRepositoryImpl {
	err := db.AutoMigrate(&dao.User{})
	if err != nil {
		log.Fatalf("Unable to migrate model changes %s", err)
	}
	return &UserRepositoryImpl{
		db: db,
	}
}
