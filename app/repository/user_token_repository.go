package repository

import (
	"accountservice/app/domain/dao"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserTokenRepository interface {
	SaveUserToken(token *dao.UserToken) (*dao.UserToken, error)
	CodeExists(code string) bool
	GetUserToken(email string) (*dao.UserToken, error)
}

type UserTokenRepositoryImpl struct {
	db *gorm.DB
}

func (u UserTokenRepositoryImpl) SaveUserToken(token *dao.UserToken) (*dao.UserToken, error) {
	var err = u.db.Save(token).Error
	if err != nil {
		log.Error("Got Error while saving userToken : ", err)
		return nil, err
	}
	return token, nil
}

func (u UserTokenRepositoryImpl) CodeExists(code string) bool {

	userTokenData := dao.UserToken{
		Code: code,
	}
	err := u.db.Where("code = ?", userTokenData.Code).First(&userTokenData).Error

	if err != nil {
		return false
	}
	return true

}

func (u UserTokenRepositoryImpl) GetUserToken(email string) (*dao.UserToken, error) {
	userTokenData := dao.UserToken{
		Email: email,
	}

	err := u.db.Where("email = ?", userTokenData.Email).First(&userTokenData).Error
	if err != nil {
		log.Error("Got and error when finding token by email. Error: ", err)
		return nil, err
	}
	return &userTokenData, nil
}

func UserTokenRepositoryInit(db *gorm.DB) *UserTokenRepositoryImpl {
	err := db.AutoMigrate(&dao.UserToken{})
	if err != nil {
		log.Fatalf("Unable to migrate model changes %s", err)
	}
	return &UserTokenRepositoryImpl{
		db: db,
	}
}
