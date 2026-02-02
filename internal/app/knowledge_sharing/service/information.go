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

// 获取用户消息
func GetInformationByUname(uname string) (information *model.Information, err error) {
	return mysql.GetInformationByUname(uname)
}

// 创建和发送消息
func MakeAndSendNotice(sendId int, uname string, content string, t time.Time) (err error) {
	notice := &model.Information{
		SendId:      sendId,
		ReceiveName: uname,
		Content:     content,
		CreateAt:    &t,
		UpdateAt:    &t,
	}

	err = AddUserNotice(notice)
	return
}

// 更新IsDeleted删除information
func UpdateInformationIsDeleted(iid int) (int64, error) {
	return mysql.UpdateInformationIsDeleted(iid)
}
