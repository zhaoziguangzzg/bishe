package service

import (
	"bishe/internal/app/knowledge_sharing/dao/mysql"
	"bishe/internal/app/knowledge_sharing/model"
)

// 添加或更新联系人
func ChatContactInsertUpdate(chatContact *model.ChatContact) (err error) {
	return mysql.ChatContactInsertUpdate(chatContact)
}

// 获取联系人列表
func GetChatContactList(uid int, page int, pageSize int) (chatContacts []model.ChatContact, err error) {
	return mysql.GetChatContactList(uid, page, pageSize)
}

// create 用户联系人
func CreateUserContact(newContect *model.Contact) (err error) {
	return mysql.CreateUserContact(newContect)
}

// 根据uid,receiveId获取联系人
func GetUserContact(uid int, receiveId int) (contact *model.Contact, err error) {
	return mysql.GetUserContact(uid, receiveId)
}

// 根据id获取联系人
func GetUserContactById(id int) (contact *model.Contact, err error) {
	return mysql.GetUserContactById(id)
}

// 获取用户全部联系人
func GetUserAllContact(uid int, page int, pagesize int) (contacts []model.Contact, err error) {
	return mysql.GetUserAllContact(uid, page, pagesize)
}

// 删除联系人
func DeleteUserContactByReceiveId(uid int, receiveId int) (int64, error) {
	return mysql.DeleteUserContactByReceiveId(uid, receiveId)
}
