package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"hello/work/new_log/models"
	"hello/work/new_log/utils"
	"net/http"
)

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

// 登录请求
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// 通用响应
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type AuthController struct {
	DB *gorm.DB
}

func NewAuthController(db *gorm.DB) *AuthController {
	return &AuthController{DB: db}
}

// 注册
func (ac *AuthController) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, Response{Code: 400, Message: err.Error()})
		return
	}

	// 检查用户名是否已存在
	var existingUser models.User
	if err := ac.DB.Where("username = ?", req.Username).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, Response{Code: 409, Message: "Username already exists"})
		return
	}

	//创建用户
	user := models.User{
		Username: req.Username,
		Email:    req.Email,
	}

	//此处原为密码加密
	//if err := user.HashPassword(req.Password); err != nil {
	//		c.JSON(http.StatusInternalServerError, Response{Code: 500, Message: "Could not hash password"})
	//		return
	//	}

	if err := ac.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, Response{Code: 500, Message: "Could not create user!"})
		return
	}

	c.JSON(http.StatusCreated, Response{
		Code:    201,
		Message: "User registered successfully!",
		Data:    gin.H{"user_id": user.ID},
	})
}

// 登录
func (ac *AuthController) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, Response{Code: 400, Message: err.Error()})
		return
	}

	//查看用户
	var user models.User
	if err := ac.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, Response{Code: 404, Message: "User not found"})
		return
	}

	//此处为验证密码

	//生成jwt令牌
	token, err := utils.GenerateJWT(user.ID, user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Code:    200,
			Message: "Could not generate JWT token!",
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "Successfully logged in!",
		Data:    gin.H{"token": token},
	})

}

func (ac *AuthController) Profile(c *gin.Context) {
	Userid, _ := c.Get("user_id")

	var user models.User
	if err := ac.DB.Where("id = ?", Userid).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, Response{Code: 404, Message: "User not found"})
	}

	c.JSON(http.StatusOK, gin.H{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
	})
}
