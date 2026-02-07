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
	err = DB.Model(&model.UserEssayCollect{}).
		Where("user_id=? and essay_id=? and is_deleted=?", uid, eid, model.COLLECT_NOT_DELETED).
		First(&userEssayCollect).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound { //没查到数据返回空
			return nil, nil
		}

		return nil, err
	}

	return userEssayCollect, nil
}

// get 用户全部收藏
func GetUserAllCollectByUid(uid int, page int, pageSize int) (userEssayCollects []model.UserEssayCollect, err error) {
	var eids []int
	offset := (page - 1) * pageSize

	err = DB.Model(&model.UserEssayCollect{}).Where("user_id and is_deleted=?", uid, model.COLLECT_NOT_DELETED).
		Order("id ASC").Offset(offset).Limit(pageSize).Pluck("essay_id", &eids).Error
	if err != nil {
		return
	}

	err = DB.Where("id IN (?)", eids).Error
	return
}

// 取消收藏
func DeleteCollectById(id int) (int64, error) {
	result := DB.Model(&model.UserEssayCollect{}).Where("id=?", id).
		Update("is_deleted", model.COLLECT_IS_DELETED)

	return result.RowsAffected, result.Error
}

// 进行收藏
func UpdateUserEssayCollectNotToIs(uid int, eid int, fid int) (int64, error) {
	collect := model.UserEssayCollect{
		FavoriteId: fid,
		IsDeleted:  model.COLLECT_NOT_DELETED,
	}

	result := DB.Model(&model.UserEssayCollect{}).Where("user_id=? and essay_id=?", uid, eid).Updates(collect)
	return result.RowsAffected, result.Error
}
