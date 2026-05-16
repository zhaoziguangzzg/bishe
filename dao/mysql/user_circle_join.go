package mysql

import (
	"bishe/model"
	"time"

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
	err = DB.Model(&model.UserCircleJoin{}).
		Where("user_id=? and circle_id=? and not_join_status=?", uid, cid, model.USER_CIRCLE_JOIN_STATUS_JOIN).
		First(&userCircleJoin).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound { //没查到数据返回空
			return nil, nil
		}

		return nil, err
	}

	return userCircleJoin, nil
}

// 获取用户圈子
func GetUserCircleJoinByUidCid(uid int, cid int) (join *model.UserCircleJoin, err error) {
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

// 根据uid查询用户加入圈子
func GetUserJoinCircleByUid(uid int, page int, pagesize int) (userCircleJoins []model.UserCircleJoin, err error) {
	offset := (page - 1) * pagesize

	err = DB.Model(&model.UserCircleJoin{}).
		Where("user_id=?  and not_join_status=?", uid, model.USER_CIRCLE_JOIN_STATUS_JOIN).
		Order("id DESC").Offset(offset).Limit(pagesize).Find(&userCircleJoins).Error
	if err != nil {
		return
	}

	return
}

// 修改用户参与圈子时间
func UpdateUserCircleJoinTimeByUidCid(uid int, cid int, startTime time.Time, endTime time.Time) (int64, error) {
	userCircleJoin := &model.UserCircleJoin{
		StartTime: &startTime,
		EndTime:   &endTime,
	}

	result := DB.Model(&model.UserCircleJoin{}).Where("user_id=? and circle_id=?", uid, cid).
		Updates(userCircleJoin)
	return result.RowsAffected, result.Error
}

// 修改用户参与圈子状态
func UpdateUserCircleJoinStatusByJid(jid int, joinStatus int) (int64, error) {
	result := DB.Model(&model.UserCircleJoin{}).Where("id=?", jid).
		UpdateColumn("not_join_status", joinStatus)
	return result.RowsAffected, result.Error
}
