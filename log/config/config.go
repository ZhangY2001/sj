package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	DBUser     string //`yaml:"DB_USER"`
	DBPassword string //`yaml:"DB_PASSWORD"`
	DBName     string //`yaml:"DB_NAME"`
	DBHost     string //`yaml:"DB_HOST"`
	DBPort     string //`yaml:"DB_PORT"`
	JWTSecret  string //`yaml:"JWT_SECRET"`
}

func LoadConfig() *Config {
	// 加载 .env 文件
	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatal("无法加载 .env 文件:", err)
	}
	return &Config{
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		JWTSecret:  os.Getenv("JWT_SECRET"),
	}
}

/*
func InitConfig() {
	//设置默认值
	viper.SetDefault("port", 8080)
	viper.SetDefault("host", "localhost")

	//设置配置文件名（不带扩展名）
	viper.SetConfigName("config")
	//设置文件类型
	viper.SetConfigType("yaml")
	//设置配置文件搜索路径
	viper.AddConfigPath(".")

	//从环境变量读取
	viper.AutomaticEnv()

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("No config file found, using defaults and environment variables")
		} else {
			log.Fatalf("Error reading config file: %v", err)
		}
	}

	//监听配置变化
	viper.WatchConfig()
}
*/
