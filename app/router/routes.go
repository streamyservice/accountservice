package router

import (
	"accountservice/config"
	"github.com/gin-gonic/gin"
)

func Init(init *config.Initialization) *gin.Engine {

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api := router.Group("/api")
	{
		user := api.Group("/user")
		user.GET("", init.UserCtrl.RegisterUser)
		user.POST("", init.UserCtrl.LoginUser)
		user.GET("/:userID", init.UserCtrl.UpdateUser)
		user.PUT("/:userID", init.UserCtrl.RefreshAuthToken)
		user.DELETE("/:userID", init.UserCtrl.DeleteUser)
	}

	return router
}
