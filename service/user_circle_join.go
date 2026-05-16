package service

import (
	"bishe/dao/mysql"
	"bishe/model"
	"time"
)

// create 用户加入圈子
func CreateUserCircleJoin(newUserCircle *model.UserCircleJoin) (err error) {
	return mysql.CreateUserCircleJoin(newUserCircle)
}

// 用户加入圈子并更新参与人数
func CreateUserJoinCircleAndUpdateJoinNum(uid int, cid int, joinTime time.Time, endTime time.Time) (joinid int, err error) {

	newUserCircle := &model.UserCircleJoin{
		UserId:        uid,
		CircleId:      cid,
		StartTime:     &joinTime,
		EndTime:       &endTime,
		JoinTime:      &joinTime,
		UpdateAt:      &joinTime,
		NotJoinStatus: model.USER_CIRCLE_JOIN_STATUS_JOIN,
	}

	err = CreateUserCircleJoin(newUserCircle)
	if err != nil {
		return
	}

	affectRows, err := IncrUpdateCircleJoinNumByCid(cid)
	if affectRows == 0 || err != nil {
		return
	}
	joinid = newUserCircle.Id
	return joinid, err
}

// 查询用户加入圈子
func GetUserJoinCircleByUidCid(uid int, cid int) (userCircleJoin *model.UserCircleJoin, err error) {
	return mysql.GetUserJoinCircleByUidCid(uid, cid)
}

// 获取用户圈子
func GetUserCircleJoinByUidCid(uid int, cid int) (join *model.UserCircleJoin, err error) {
	return mysql.GetUserCircleJoinByUidCid(uid, cid)
}

// 根据uid查询用户加入圈子
func GetUserJoinCircleByUid(uid int, page int, pagesize int) (userCircleJoins []model.UserCircleJoin, err error) {
	return mysql.GetUserJoinCircleByUid(uid, page, pagesize)
}

// 修改用户参与圈子时间
func UpdateUserCircleJoinTimeByUidCid(uid int, cid int, startTime time.Time, endTime time.Time) (int64, error) {
	return mysql.UpdateUserCircleJoinTimeByUidCid(uid, cid, startTime, endTime)
}

// 修改用户参与圈子状态
func UpdateUserCircleJoinStatusByJid(jid int, joinStatus int) (int64, error) {
	return mysql.UpdateUserCircleJoinStatusByJid(jid, joinStatus)
}
