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
		user.POST("/register", init.UserCtrl.RegisterUser)
		user.POST("/login", init.UserCtrl.LoginUser)
		user.POST("/refreshToken", init.UserCtrl.RefreshAuthToken)
		user.POST("/verifyEmail", init.UserCtrl.VerifyEmail)
		user.GET("/:userEmail", init.UserCtrl.GetUserByEmail)
		user.PUT("/update/:userEmail", init.UserCtrl.UpdateUser)
		user.DELETE("/delete/:userEmail", init.UserCtrl.DeleteUser)

	}

	return router
}
