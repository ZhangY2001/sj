package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"time"
)

type User struct {
	gorm.Model
	Username      string //
	Password      string
	Phone         string
	Email         string
	Avatar        string //头像
	Identity      string
	ClientIp      string
	ClientPort    string
	Salt          string
	LoginTime     time.Time
	HeartbeatTine time.Time
	LastLoginTime time.Time
	IsLogOut      bool
	DeviceInfo    string
	PasswordHash  string `gorm:"not null"`
	Age           int
}

// 增加用户
//func (user *User) AddUser() *gorm.DB {
//return utils.DB.Create(&user)
//}

// 删除用户
//func (user *User) DeleteUser() *gorm.DB {
//return utils.DB.Delete(&user)
//}

// 修改用户信息
//func (user *User) UpdateUser() *gorm.DB {
//return utils.DB.Update(&user)
//}

// 查询用户信息
//func (user *User) CheckUser() *gorm.DB {
//return utils.DB.First(&user, "id = ?", user.ID)
//}

// 获取用户信息
//func GetUserList() []*User {
//data := make([]*User, 0, 10)
//从数据库中获取用户列表
//utils.DB.Find(&data)
//for _, v := range data {
//fmt.Println(v)
//}
//return data
//}

// 密码加密
func (u *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.PasswordHash = string(bytes)
	return nil
}

// 验证密码
func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
}
