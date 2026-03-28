package service

import (
	"bishe/internal/app/knowledge_sharing/dao/mysql"
	"bishe/internal/app/knowledge_sharing/model"
)

// create 用户关注
func CreateUserFollow(newFollow *model.Follow) (err error) {
	return mysql.CreateUserFollow(newFollow)
}

// 根据uid,followid获取关注
func GetUserFollow(uid int, followerId int) (follow *model.Follow, err error) {
	return mysql.GetUserFollow(uid, followerId)
}

// 根据uid,followid获取关注
func GetUserFollowByStatus(uid int, followerId int) (follow *model.Follow, err error) {
	return mysql.GetUserFollowByStatus(uid, followerId)
}

// 进行关注
func UpdateUserFollowNotToIs(uid int, followId int) (int64, error) {
	return mysql.UpdateUserFollowNotToIs(uid, followId)
}

// 取关
func UpdateUserFollowIsToNot(uid int, followerId int) (int64, error) {
	return mysql.UpdateUserFollowIsToNot(uid, followerId)
}

// get 用户的关注
func GetUserFollowListByUid(uid int, page int, pagesize int) (users []model.User, err error) {
	return mysql.GetUserFollowListByUid(uid, page, pagesize)
}

// get 用户粉丝
func GetUserFanListByUid(uid int, page int, pagesize int) (users []model.User, err error) {
	return mysql.GetUserFanListByUid(uid, page, pagesize)
}
