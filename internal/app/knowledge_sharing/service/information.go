package service

import (
	"bishe/internal/app/knowledge_sharing/dao/mysql"
	"bishe/internal/app/knowledge_sharing/model"
)

// 创建用户消息
func UserAddInformation(information *model.Information) (err error) {
	return mysql.UserAddInformation(information)
}

// 通知
func AddUserNotice(notice *model.Information) (err error) {
	return mysql.AddUserNotice(notice)
}
