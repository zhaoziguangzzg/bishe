package service

import (
	"bishe/internal/app/knowledge_sharing/dao/mysql"
	"bishe/internal/app/knowledge_sharing/model"
	"time"
)

// 创建用户消息
func UserAddInformation(information *model.Information) (err error) {
	return mysql.UserAddInformation(information)
}

// 通知
func AddUserNotice(notice *model.Information) (err error) {
	return mysql.AddUserNotice(notice)
}

// 创建和发送消息
func MakeAndSendNotice(sendId int, account int, content string, t time.Time) (err error) {
	notice := &model.Information{
		SendId:         sendId,
		ReceiveAccount: account,
		Content:        content,
		CreateAt:       &t,
	}

	err = AddUserNotice(notice)
	return
}
