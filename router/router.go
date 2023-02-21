package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kkomissarov/beggar/controllers"
	"github.com/kkomissarov/beggar/middlewares"
)

func Router() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("v1")
	{
		auth := v1.Group("auth")
		{
			auth.POST("/signup/", controllers.SignUp)
			auth.POST("/login/", controllers.Login)
			auth.POST("/logout/", middlewares.AuthRequiredMiddleware(), controllers.Logout)
		}
		accounts := v1.Group("accounts")
		accounts.Use(middlewares.AuthRequiredMiddleware())
		{
			accounts.POST("/", controllers.CreateAccount)
			accounts.GET("/", controllers.GetAccountList)
			accounts.GET("/:id/", controllers.GetAccount)
			accounts.PUT("/:id/", controllers.UpdateAccount)
			accounts.DELETE("/:id/", controllers.DeleteAccount)
		}
	}

	return router
}
