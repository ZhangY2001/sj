package config

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	DBUser     string
	DBPassword string
	DBName     string
	DBHost     string
	DBPort     string
	JWTSecret  string
}

func LoadConfig() *Config {
	return &Config{
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		JWTSecret:  os.Getenv("JWT_SECRET"),
	}
}

var GConfig Config

const configPath = "./config.yaml"

func InitConfig() {
	file, err := os.ReadFile(configPath)
	if err != nil {
		logrus.Fatal("配置文件失败：", err.Error())
	}
	c := new(Config)
	err = yaml.Unmarshal(file, c)
	//GConfig = c
	if err != nil {
		logrus.Fatal("配置文件解析失败：", err.Error())
	}
}
