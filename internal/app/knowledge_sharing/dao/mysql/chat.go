package mysql

import "bishe/internal/app/knowledge_sharing/model"

// 用户创建私信
func ChatAdd(chat *model.Chat) (err error) {
	err = DB.Model(&model.Chat{}).Create(chat).Error
	return
}

// 获取私信列表
func GetChatList(uid int, chatUid int, page int, pageSize int) (chats []model.Chat, err error) {
	offset := (page - 1) * pageSize

	err = DB.Model(&model.Chat{}).
		Where("send_uid=? and receive_uid=? and is_deleted=?", uid, chatUid, model.IS_DELETED_NO).
		Or("send_uid=? and receive_uid=? and is_deleted=?", chatUid, uid, model.IS_DELETED_NO).
		Order("id ASC").Offset(offset).Limit(pageSize).Find(&chats).Error
	if err != nil {
		return
	}

	return
}
