package service

import (
	"bishe/dao/mysql"
	"bishe/model"
)

// 用户创建私信
func ChatAdd(chat *model.Chat) (err error) {
	return mysql.ChatAdd(chat)
}

// 获取私信列表（分页）
func GetChatList(uid int, chatUid int, page int, pageSize int) (chats []model.Chat, err error) {
	return mysql.GetChatList(uid, chatUid, page, pageSize)
}

// 获取最新的N条私信
func GetChatListLatest(uid, chatUid, limit int) ([]model.Chat, error) {
	return mysql.GetChatListLatest(uid, chatUid, limit)
}

// 获取baseId之前的N条私信
func GetChatListBefore(uid, chatUid, baseId, limit int) ([]model.Chat, error) {
	return mysql.GetChatListBefore(uid, chatUid, baseId, limit)
}

// 获取baseId之后的N条私信
func GetChatListAfter(uid, chatUid, baseId, limit int) ([]model.Chat, error) {
	return mysql.GetChatListAfter(uid, chatUid, baseId, limit)
}
