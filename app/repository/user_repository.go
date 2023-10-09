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
	UserExists(email string) bool
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func (u UserRepositoryImpl) CreateUser(user *dao.User) (*dao.User, error) {
	var err = u.db.Save(user).Error
	if err != nil {
		log.Error("Got Error while saving user : ", err)
		return nil, err
	}
	return user, nil

}

func (u UserRepositoryImpl) UpdateUser(email string, user *dao.User) (*dao.User, error) {
	return nil, nil
}

func (u UserRepositoryImpl) GetUser(email string) (*dao.User, error) {
	user := dao.User{
		Email: email,
	}

	err := u.db.Where("email = ?", user.Email).First(&user).Error
	if err != nil {
		log.Error("Got and error when finding user by email. Error: ", err)
		return nil, err
	}
	return &user, nil

}

func (u UserRepositoryImpl) UserExists(email string) bool {
	user := dao.User{
		Email: email,
	}
	err := u.db.Where("email = ?", user.Email).First(&user).Error
	if err != nil {
		return false
	}
	return true
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
