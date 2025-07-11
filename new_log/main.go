package main

import (
	"github.com/gin-gonic/gin"

	"hello/work/new_log/db"

	"hello/work/new_log/models"
	"hello/work/new_log/routes"
	"log"

	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID       uint
	Username string
	Age      int
}

func main() {
	//初始化数据库
	db.InitDB()

	//自动迁移
	db.AutoMigrate(&models.User{})

	// 创建记录

	user := User{Username: "Bob", Age: 30}
	result := db.Create(&user)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Printf("Inserted ID: %d\n", user.ID)

	// 查询
	var retrievedUser User
	db.First(&retrievedUser, "Name = ?", "Bob")
	fmt.Printf("Retrieved user: %+v\n", retrievedUser)

	//初始化Gin
	router := gin.Default()

	//设置路由
	routes.SetupAuthRouter(router, db)

	//启动服务器
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server :", err)
	}
}
