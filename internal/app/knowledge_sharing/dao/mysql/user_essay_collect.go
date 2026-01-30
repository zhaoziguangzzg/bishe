package mysql

import (
	"bishe/internal/app/knowledge_sharing/model"

	"gorm.io/gorm"
)

// create 用户对文章收藏
func CreateUserEssayCollect(newUserEssayCollect *model.UserEssayCollect) (err error) {
	err = DB.Model(&model.UserEssayCollect{}).Create(newUserEssayCollect).Error
	return
}

// 根据uid,eid获取文章点赞
func GetUserEssayCollect(uid int, eid int) (userEssayCollect *model.UserEssayCollect, err error) {
	userEssayCollect = new(model.UserEssayCollect)
	err = DB.Model(&model.UserEssayCollect{}).Where("user_id=? and essay_id=?", uid, eid).
		First(&userEssayCollect).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound { //没查到数据返回空
			return nil, nil
		}

		return nil, err
	}

	return userEssayCollect, nil
}
