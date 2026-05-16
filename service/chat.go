package service

import (
	"bishe/dao/mysql"
	"bishe/model"
)

// 用户创建私信
func ChatAdd(chat *model.Chat) (err error) {
	return mysql.ChatAdd(chat)
}

// 获取私信列表
func GetChatList(uid int, chatUid int, page int, pageSize int) (chats []model.Chat, err error) {
	return mysql.GetChatList(uid, chatUid, page, pageSize)
}
