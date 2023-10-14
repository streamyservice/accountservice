package controller

import (
	"accountservice/app/service"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	RegisterUser(c *gin.Context)
	LoginUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	RefreshAuthToken(c *gin.Context)
	VerifyEmail(c *gin.Context)
	GetUserByEmail(c *gin.Context)
}

type UserControllerImpl struct {
	svc service.UserService
}

func (u UserControllerImpl) RegisterUser(c *gin.Context) {
	u.svc.CreateUser(c)

}

func (u UserControllerImpl) LoginUser(c *gin.Context) {
	u.svc.Login(c)
}
func (u UserControllerImpl) UpdateUser(c *gin.Context) {

}
func (u UserControllerImpl) DeleteUser(c *gin.Context) {

}
func (u UserControllerImpl) RefreshAuthToken(c *gin.Context) {

}
func (u UserControllerImpl) VerifyEmail(c *gin.Context) {

	u.svc.VerifyUserEmail(c)

}
func (u UserControllerImpl) GetUserByEmail(c *gin.Context) {

	u.svc.GetUser(c)

}

func UserControllerInit(userService service.UserService) *UserControllerImpl {
	return &UserControllerImpl{
		svc: userService,
	}
}
