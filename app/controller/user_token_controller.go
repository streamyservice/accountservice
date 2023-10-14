package controller

import "github.com/gin-gonic/gin"

type UserTokenController interface {
	GetTokenByEmail(c *gin.Context)
}
