package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"hello/work/new_log/controllers"
	"hello/work/new_log/middlewares"
)

func SetupAuthRouter(router *gin.Engine, db *gorm.DB) {
	authController := controllers.NewAuthController(db)

	authGroup := router.Group("/auth")
	{
		authGroup.POST("/register", authController.Register)
		authGroup.POST("/login", authController.Login)

		// 需要认证的路由
		authGroup.Use(middlewares.AuthMiddleware())

	}

}
