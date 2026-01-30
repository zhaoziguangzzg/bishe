package mysql

import (
	"bishe/internal/app/knowledge_sharing/model"

	"gorm.io/gorm"
)

// create 用户对文章点赞
func CreateUserEssayLike(newUserEssayLike *model.UserEssayLike) (err error) {
	err = DB.Model(&model.UserEssayLike{}).Create(newUserEssayLike).Error
	return
}

// get 用户对文章点赞
func GetUserEssayInteract(uid int, circleId int, essayId int) (userEssayInteract *model.UserEssayLike, err error) {
	err = DB.Model(&model.UserEssayLike{}).Where("user_id=? and circle_id=? and essay_id=?", uid, circleId, essayId).Find(userEssayInteract).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound { //没查到数据返回空
			return nil, nil
		}

		return nil, err
	}

	return userEssayInteract, nil
}

// 根据uid,eid获取文章点赞
func GetUserEssayLike(uid int, eid int) (userEssayLike *model.UserEssayLike, err error) {
	userEssayLike = new(model.UserEssayLike)
	err = DB.Model(&model.UserEssayLike{}).Where("user_id=? and essay_id=?", uid, eid).
		First(&userEssayLike).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound { //没查到数据返回空
			return nil, nil
		}

		return nil, err
	}

	return userEssayLike, nil
}

// get 用户全部点赞
func GetUserAllLikeByUid(uid int, page int, pageSize int) (userEssayInteracts []model.UserEssayLike, err error) {
	offset := (page - 1) * pageSize
	err = DB.Model(&model.UserEssayLike{}).Where("user_id and is_deleted=?", uid, model.LIKE_NOT_DELETED).
		Order("id ASC").Offset(offset).Limit(pageSize).Find(&userEssayInteracts).Error
	return
}

// 取消喜欢
func UpdateUserEssayLikeIsToNot(uid int, eid int) (int64, error) {
	result := DB.Model(&model.UserEssayLike{}).Where("user_id=? and essay_id=?", uid, eid).
		Update("is_deleted", model.LIKE_IS_DELETED)
	return result.RowsAffected, result.Error
}

// 进行喜欢
func UpdateUserEssayLikeNotToIs(uid int, eid int) (int64, error) {
	result := DB.Model(&model.UserEssayLike{}).Where("user_id=? and essay_id=?", uid, eid).
		Update("is_deleted", model.LIKE_NOT_DELETED)
	return result.RowsAffected, result.Error
}

// // 更新 用户对文章点赞
// func UpdateUserEssayInteractLike(userEssayInteract *model.UserEssayLike, uid int, circleId int, essayId int) (result *gorm.DB) {
// 	if userEssayInteract.LikeStatus == model.UserEssayInteractNotJoin {
// 		result = DB.Model(&model.UserEssayLike{}).Where("user_id=? and circle_id=? and essay_id=?", uid, circleId, essayId).Updates(model.UserEssayInteract{LikeStatus: model.UserEssayInteractJoin})
// 	} else if userEssayInteract.LikeStatus == model.UserEssayInteractJoin {
// 		result = DB.Model(&model.UserEssayLike{}).Where("user_id=? and circle_id=? and essay_id=?", uid, circleId, essayId).Updates(model.UserEssayInteract{LikeStatus: model.UserEssayInteractNotJoin})
// 	}

// 	return
// }

// // 更新 用户对文章收藏
// func UpdateUserEssayInteractCollect(userEssayInteract *model.UserEssayLike, uid int, circleId int, essayId int, favorite string) (result *gorm.DB) {
// 	if userEssayInteract.CollectStatus == model.UserEssayInteractNotJoin {
// 		result = DB.Model(&model.UserEssayLike{}).Where("user_id=? and circle_id=? and essay_id=?", uid, circleId, essayId).Updates(model.UserEssayInteract{CollectStatus: model.UserEssayInteractJoin, Favorite: favorite})

// 	} else if userEssayInteract.CollectStatus == model.UserEssayInteractJoin {
// 		result = DB.Model(&model.UserEssayLike{}).Where("user_id=? and circle_id=? and essay_id=?", uid, circleId, essayId).Updates(model.UserEssayInteract{CollectStatus: model.UserEssayInteractNotJoin})
// 	}

// 	return
// }
