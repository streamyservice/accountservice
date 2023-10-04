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
}

type UserControllerImpl struct {
	svc service.UserService
}
