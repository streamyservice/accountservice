package service

import (
	"accountservice/app/repository"
	"github.com/gin-gonic/gin"
)

type UserService interface {
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	GetUser(c *gin.Context)
	RefreshAuthToken(c *gin.Context)
	VerifyUserEmail(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type UserServiceImpl struct {
	userRepository repository.UserRepository
}

func (u UserServiceImpl) CreateUser(c *gin.Context) {

}
func (u UserServiceImpl) UpdateUser(c *gin.Context) {

}
func (u UserServiceImpl) GetUser(c *gin.Context) {

}
func (u UserServiceImpl) RefreshAuthToken(c *gin.Context) {

}
func (u UserServiceImpl) VerifyUserEmail(c *gin.Context) {

}
func (u UserServiceImpl) DeleteUser(c *gin.Context) {

}
func UserServiceInit(userRepository repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		userRepository: userRepository,
	}
}
