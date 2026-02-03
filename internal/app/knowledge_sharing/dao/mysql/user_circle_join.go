package mysql

import (
	"bishe/internal/app/knowledge_sharing/model"

	"gorm.io/gorm"
)

// create 用户加入圈子
func CreateUserCircleJoin(newUserCircleJoin *model.UserCircleJoin) (err error) {
	err = DB.Model(&model.UserCircleJoin{}).Create(newUserCircleJoin).Error
	return
}

// 根据uid，cid查询用户加入圈子
func GetUserJoinCircleByUidCid(uid int, cid int) (userCircleJoin *model.UserCircleJoin, err error) {
	userCircleJoin = new(model.UserCircleJoin)
	err = DB.Model(&model.UserCircleJoin{}).Where("user_id=? and circle_id=? and not_join_status=?", uid, cid, model.USER_CIRCLE_NOT_NO_JOIN).First(&userCircleJoin).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound { //没查到数据返回空
			return nil, nil
		}

		return nil, err
	}

	return userCircleJoin, nil
}

// 获取用户圈子
func GetUserCircleJoinByJoin(uid int, cid int) (join *model.UserCircleJoin, err error) {
	join = new(model.UserCircleJoin)
	err = DB.Model(&model.UserCircleJoin{}).Where("user_id=? and circle_id=?", uid, cid).First(&join).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound { //没查到数据返回空
			return nil, nil
		}

		return nil, err
	}

	return join, nil
}

// 用户退出圈子
func UpdateUserCircleNotJoinStatusByUidCid(uid int, cid int) (int64, error) {
	result := DB.Model(&model.UserCircleJoin{}).Where("user_id=? and circle_id=?", uid, cid).UpdateColumn("not_join_status", model.USER_CIRCLE_NOT_JOIN)
	return result.RowsAffected, result.Error
}

// 用户加入圈子
func UpdateUserCircleJoinStatusByUidCid(uid int, cid int) (int64, error) {
	result := DB.Model(&model.UserCircleJoin{}).Where("user_id=? and circle_id=?", uid, cid).UpdateColumn("not_join_status", model.USER_CIRCLE_NOT_NO_JOIN)
	return result.RowsAffected, result.Error
}
