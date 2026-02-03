package service

import (
	"bishe/internal/app/knowledge_sharing/dao/mysql"
	"bishe/internal/app/knowledge_sharing/model"
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
func GetUserCircleJoinByJoin(uid int, cid int) (join *model.UserCircleJoin, err error) {
	return mysql.GetUserCircleJoinByJoin(uid, cid)
}

// 用户退出圈子
func UpdateUserCircleNotJoinStatusByUidCid(uid int, cid int) (int64, error) {
	return mysql.UpdateUserCircleNotJoinStatusByUidCid(uid, cid)
}

// 用户加入圈子
func UpdateUserCircleJoinStatusByUidCid(uid int, cid int) (int64, error) {
	return mysql.UpdateUserCircleJoinStatusByUidCid(uid, cid)
}
