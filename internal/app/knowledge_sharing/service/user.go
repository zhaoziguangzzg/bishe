package service

import (
	"bishe/internal/app/knowledge_sharing/dao/mysql"
	"bishe/internal/app/knowledge_sharing/model"

	"gorm.io/gorm"
)

// create用户
func CreateUser(newUser *model.User) (err error) {
	return mysql.CreateUser(newUser)
}

// 获取用户密码
func GetPasswordFromUser(account int) (password string, err error) {
	return mysql.GetPasswordFromUser(account)
}

// 根据account获取用户
func GetUserByAccount(account int) (user *model.User, err error) {
	return mysql.GetUserByAccount(account)
}

// 根据id获取用户
func GetUserByUserId(UserId int) (user *model.User, err error) {
	return mysql.GetUserByUserId(UserId)
}

// 更新
func UpdateFromUser(user *model.User, password string, email string, age int, phone int) (result *gorm.DB) {
	return mysql.UpdateFromUser(user, password, email, age, phone)
}
