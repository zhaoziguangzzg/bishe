package mysql

import (
	"bishe/internal/app/knowledge_sharing/model"

	"gorm.io/gorm"
)

// create 用户关注
func CreateUserFollow(newFollow *model.Follow) (err error) {
	err = DB.Model(&model.Follow{}).Create(newFollow).Error
	return
}

// 根据uid,followid获取关注
func GetUserFollow(uid int, followerId int) (follow *model.Follow, err error) {
	follow = new(model.Follow)
	err = DB.Model(&model.Follow{}).Where("fan_id=? and follower_id=? and follow_status=?", uid, followerId, model.FOLLOW_STATUS_NORMAL).First(&follow).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound { //没查到数据返回空
			return nil, nil
		}

		return nil, err
	}

	return follow, nil
}

// 进行关注
func UpdateUserFollowNotToIs(uid int, followerId int) (int64, error) {
	result := DB.Model(&model.Follow{}).Where("fan_id=? and follower_id=?", uid, followerId).
		Update("follow_status", model.FOLLOW_STATUS_NORMAL)
	return result.RowsAffected, result.Error
}

// 取关
func UpdateUserFollowIsToNot(uid int, followerId int) (int64, error) {
	result := DB.Model(&model.Follow{}).Where("fan_id=? and follower_id=?", uid, followerId).
		Update("follow_status", model.FOLLOW_STATUS_REVIEW)
	return result.RowsAffected, result.Error
}

// get 用户的关注
func GetUserFollowListByUid(uid int, page int, pagesize int) (users []model.User, err error) {
	var uids []int
	offset := (page - 1) * pagesize

	err = DB.Model(&model.Follow{}).Where("fan_id=? and follow_status=?", uid, model.FOLLOW_STATUS_NORMAL).
		Order("follow_time DESC").Offset(offset).Limit(pagesize).Pluck("follower_id", &uids).Error
	if err != nil {
		return
	}

	err = DB.Model(&model.User{}).Where("id IN (?)", uids).Find(&users).Error
	if err != nil {
		return
	}

	return

}

// get 用户粉丝
func GetUserFanListByUid(uid int, page int, pagesize int) (users []model.User, err error) {
	var uids []int
	offset := (page - 1) * pagesize

	err = DB.Model(&model.Follow{}).Where("follower_id=? and follow_status=?", uid, model.FOLLOW_STATUS_NORMAL).
		Order("follow_time DESC").Offset(offset).Limit(pagesize).Pluck("fan_id", &uids).Error
	if err != nil {
		return
	}

	err = DB.Model(&model.User{}).Where("id IN (?)", uids).Find(&users).Error
	if err != nil {
		return
	}

	return

}
