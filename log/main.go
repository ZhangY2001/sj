package main

import (
	"github.com/gin-gonic/gin"
	"hello/work/log/controllers"
	"hello/work/log/database"
	"hello/work/log/middleware"
	"hello/work/log/models"
)

func main() {

	// 初始化数据库
	database.InitDB()

	// 自动迁移模型
	db := database.DB

	db.AutoMigrate(&models.User{})

	// 创建Gin路由
	r := gin.Default()

	// 路由组
	authController := controllers.AuthController{}

	// 公共路由
	r.POST("/register", authController.Register)
	r.POST("/login", authController.Login)

	// 需要认证的路由
	auth := r.Group("/auth")
	auth.Use(middleware.JWTAuth())
	{
		auth.GET("/profile", authController.Profile)
	}

	// 启动服务器
	r.Run(":8080")
}
