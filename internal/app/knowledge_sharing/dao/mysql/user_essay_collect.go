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

// 根据uid,eid获取文章收藏
func GetUserEssayCollect(uid int, eid int) (userEssayCollect *model.UserEssayCollect, err error) {
	userEssayCollect = new(model.UserEssayCollect)
	err = DB.Model(&model.UserEssayCollect{}).Where("user_id=? and essay_id=?", uid, eid).First(&userEssayCollect).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound { //没查到数据返回空
			return nil, nil
		}

		return nil, err
	}

	return userEssayCollect, nil
}

// get 用户收藏夹全部收藏
func GetUserAllCollectByUidFid(uid int, fid int, page int, pageSize int) (eids []int, err error) {
	offset := (page - 1) * pageSize

	err = DB.Model(&model.UserEssayCollect{}).
		Where("user_id=? and favorite_id=? and collect_status=?", uid, fid, model.COLLECT_STATUS_NORMAL).
		Order("id ASC").Offset(offset).Limit(pageSize).Pluck("essay_id", &eids).Error
	if err != nil {
		return
	}

	return
}

// 取消收藏
func UpdateUserEssayCollectIsToNot(uid int, eid int) (int64, error) {
	result := DB.Model(&model.UserEssayCollect{}).Where("user_id=? and essay_id=?", uid, eid).
		Update("collect_status", model.COLLECT_STATUS_REVIEW)

	return result.RowsAffected, result.Error
}

// 进行收藏
func UpdateUserEssayCollectNotToIs(uid int, eid int, fid int) (int64, error) {
	collect := map[string]interface{}{
		"favorite_id":    fid,
		"collect_status": model.COLLECT_STATUS_NORMAL,
	}

	result := DB.Model(&model.UserEssayCollect{}).Where("user_id=? and essay_id=?", uid, eid).Updates(collect)
	return result.RowsAffected, result.Error
}
