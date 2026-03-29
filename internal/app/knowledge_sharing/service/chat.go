package service

import (
	"bishe/internal/app/knowledge_sharing/dao/mysql"
	"bishe/internal/app/knowledge_sharing/model"
)

// 用户创建私信
func ChatAdd(chat *model.Chat) (err error) {
	return mysql.ChatAdd(chat)
}

// 获取私信列表
func GetChatList(uid int, chatUid int, page int, pageSize int) (chats []model.Chat, err error) {
	return mysql.GetChatList(uid, chatUid, page, pageSize)
}
