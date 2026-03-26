package service

import (
	"bishe/internal/app/knowledge_sharing/dao/mysql"
	"bishe/internal/app/knowledge_sharing/model"
)

// create 用户联系人
func CreateUserContact(newContect *model.Contact) (err error) {
	return mysql.CreateUserContact(newContect)
}

// 根据uid,receiveId获取联系人
func GetUserContact(uid int, receiveId int) (contact *model.Contact, err error) {
	return mysql.GetUserContact(uid, receiveId)
}
