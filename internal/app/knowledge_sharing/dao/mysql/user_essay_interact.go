package mysql

import (
	"bishe/internal/app/knowledge_sharing/model"

	"gorm.io/gorm"
)

// create 用户对文章点赞收藏
func CreateUserEssayInteract(newUserEssayInterAct *model.UserEssayInteract) (err error) {
	err = DB.Model(&model.UserEssayInteract{}).Create(newUserEssayInterAct).Error
	return
}

// get 用户对文章点赞收藏
func GetUserEssayInteract(uid int, circleId int, essayId int) (userEssayInteract *model.UserEssayInteract, err error) {
	err = DB.Model(&model.UserEssayInteract{}).Where("user_id=? and circle_id=? and essay_id=?", uid, circleId, essayId).Find(userEssayInteract).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound { //没查到数据返回空
			return nil, nil
		}

		return nil, err
	}

	return userEssayInteract, nil
}

// 更新 用户对文章点赞
func UpdateUserEssayInteractLike(userEssayInteract *model.UserEssayInteract, uid int, circleId int, essayId int) (result *gorm.DB) {
	if userEssayInteract.LikeStatus == model.UserEssayInteractNotJoin {
		result = DB.Model(&model.UserEssayInteract{}).Where("user_id=? and circle_id=? and essay_id=?", uid, circleId, essayId).Updates(model.UserEssayInteract{LikeStatus: model.UserEssayInteractJoin})
	} else if userEssayInteract.LikeStatus == model.UserEssayInteractJoin {
		result = DB.Model(&model.UserEssayInteract{}).Where("user_id=? and circle_id=? and essay_id=?", uid, circleId, essayId).Updates(model.UserEssayInteract{LikeStatus: model.UserEssayInteractNotJoin})
	}

	return
}

// 更新 用户对文章收藏
func UpdateUserEssayInteractCollect(userEssayInteract *model.UserEssayInteract, uid int, circleId int, essayId int, favorite string) (result *gorm.DB) {
	if userEssayInteract.CollectStatus == model.UserEssayInteractNotJoin {
		result = DB.Model(&model.UserEssayInteract{}).Where("user_id=? and circle_id=? and essay_id=?", uid, circleId, essayId).Updates(model.UserEssayInteract{CollectStatus: model.UserEssayInteractJoin, Favorite: favorite})

	} else if userEssayInteract.CollectStatus == model.UserEssayInteractJoin {
		result = DB.Model(&model.UserEssayInteract{}).Where("user_id=? and circle_id=? and essay_id=?", uid, circleId, essayId).Updates(model.UserEssayInteract{CollectStatus: model.UserEssayInteractNotJoin})
	}

	return
}
