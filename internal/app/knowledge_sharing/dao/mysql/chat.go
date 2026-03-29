package mysql

import "bishe/internal/app/knowledge_sharing/model"

// 用户创建私信
func ChatAdd(chat *model.Chat) (err error) {
	err = DB.Model(&model.Chat{}).Create(chat).Error
	return
}
