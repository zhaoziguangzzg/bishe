package mysql

import (
	"bishe/model"

	"gorm.io/gorm"
)

// 用户创建私信
func ChatAdd(chat *model.Chat) (err error) {
	err = DB.Model(&model.Chat{}).Create(chat).Error
	return
}

// 获取私信列表（分页）
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

func chatWhere(uid, chatUid int) *gorm.DB {
	return DB.Model(&model.Chat{}).Where(
		"(send_uid=? AND receive_uid=? AND is_deleted=?) OR (send_uid=? AND receive_uid=? AND is_deleted=?)",
		uid, chatUid, model.IS_DELETED_NO,
		chatUid, uid, model.IS_DELETED_NO,
	)
}

// 获取最新的N条私信（降序取前N，再反转成升序）
func GetChatListLatest(uid, chatUid, limit int) ([]model.Chat, error) {
	var chats []model.Chat
	err := chatWhere(uid, chatUid).Order("id DESC").Limit(limit).Find(&chats).Error
	if err != nil {
		return nil, err
	}
	// 反转成 id ASC
	for i, j := 0, len(chats)-1; i < j; i, j = i+1, j-1 {
		chats[i], chats[j] = chats[j], chats[i]
	}
	return chats, nil
}

// 获取baseId之前的N条私信（更早的）
func GetChatListBefore(uid, chatUid, baseId, limit int) ([]model.Chat, error) {
	var chats []model.Chat
	err := chatWhere(uid, chatUid).Where("id < ?", baseId).Order("id DESC").Limit(limit).Find(&chats).Error
	if err != nil {
		return nil, err
	}
	for i, j := 0, len(chats)-1; i < j; i, j = i+1, j-1 {
		chats[i], chats[j] = chats[j], chats[i]
	}
	return chats, nil
}

// 获取baseId之后的N条私信（更新的）
func GetChatListAfter(uid, chatUid, baseId, limit int) ([]model.Chat, error) {
	var chats []model.Chat
	err := chatWhere(uid, chatUid).Where("id > ?", baseId).Order("id ASC").Limit(limit).Find(&chats).Error
	if err != nil {
		return nil, err
	}
	return chats, nil
}
