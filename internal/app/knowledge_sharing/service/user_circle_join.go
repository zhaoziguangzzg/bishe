package service

import (
	"bishe/internal/app/knowledge_sharing/dao/mysql"
	"bishe/internal/app/knowledge_sharing/model"
	"time"
)

// create 用户加入圈子
func CreateUserCircleJoin(newUserCircle *model.UserCircleJoin) (err error) {
	return mysql.CreateUserCircleJoin(newUserCircle)
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
