package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"hello/work/log/config"
	"log"
)

var DB *gorm.DB

func InitDB() {
	//dsn := "root:123456@qwe@tcp(127.0.0.1:3306)/zy?charset=utf8mb4&parseTime=True&loc=Local"
	cfg := config.LoadConfig()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {

		log.Fatalf("failed to connect database: %v\n ", err)
	}

	fmt.Println("Database connected successfully")
}
