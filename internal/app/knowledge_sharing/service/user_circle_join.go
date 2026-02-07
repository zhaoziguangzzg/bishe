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
func GetUserCircleJoinByUidCid(uid int, cid int) (join *model.UserCircleJoin, err error) {
	return mysql.GetUserCircleJoinByUidCid(uid, cid)
}

// 修改用户参与圈子状态
func UpdateUserCircleJoinStatusByJid(jid int, joinStatus int) (int64, error) {
	return mysql.UpdateUserCircleJoinStatusByJid(jid, joinStatus)
}
