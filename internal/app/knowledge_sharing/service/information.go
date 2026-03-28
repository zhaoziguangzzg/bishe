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

// 创建通知举报违规消息
func AddAccusationInformation(content string, receiveId int) (err error) {
	return mysql.AddAccusationInformation(content, receiveId)
}

// 创建通知反馈消息
func AddFeedbackInformation(content string, receiveId int) (err error) {
	return mysql.AddFeedbackInformation(content, receiveId)
}

// 获取消息用户列表
func GetUserInformation(uid int, page int, pageSize int) (informations []model.Information, err error) {
	return mysql.GetUserInformation(uid, page, pageSize)
}

// 通知
func AddUserNotice(notice *model.Information) (err error) {
	return mysql.AddUserNotice(notice)
}

// 获取用户消息
func GetInformationByUname(uname string) (information *model.Information, err error) {
	return mysql.GetInformationByUname(uname)
}

// 获取用户接收消息
func GetReceiveInformationByUid(uid int, sendId int, page int, pageSize int) (informations []model.Information, err error) {
	return mysql.GetReceiveInformationByUid(uid, sendId, page, pageSize)
}

// 获取用户发送消息
func GetSendInformationByUid(uid int, receiveId int, page int, pageSize int) (informations []model.Information, err error) {
	return mysql.GetSendInformationByUid(uid, receiveId, page, pageSize)
}

// 获取用户消息
func GetUserAllInformation(uid int, page int, pageSize int) (informations []model.Information, err error) {
	return mysql.GetUserAllInformation(uid, page, pageSize)
}

// 创建和发送消息
func MakeAndSendNotice(sendId int, receiveId int, content string, t time.Time) (err error) {
	notice := &model.Information{
		SendId:    sendId,
		ReceiveId: receiveId,
		Content:   content,
		CreateAt:  &t,
		UpdateAt:  &t,
	}

	err = AddUserNotice(notice)
	return
}

// 更新IsDeleted删除information
func UpdateInformationIsDeleted(iid int) (int64, error) {
	return mysql.UpdateInformationIsDeleted(iid)
}
