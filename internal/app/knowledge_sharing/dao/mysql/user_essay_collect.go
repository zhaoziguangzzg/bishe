package mysql

import "bishe/internal/app/knowledge_sharing/model"

// create 用户对文章收藏
func CreateUserEssayCollect(newUserEssayCollect *model.UserEssayCollect) (err error) {
	err = DB.Model(&model.UserEssayCollect{}).Create(newUserEssayCollect).Error
	return
}
