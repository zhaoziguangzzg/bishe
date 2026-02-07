package service

import (
	"bishe/internal/app/knowledge_sharing/dao/mysql"
	"bishe/internal/app/knowledge_sharing/model"
)

// create 用户对文章收藏
func CreateUserEssayCollect(newUserEssayCollect *model.UserEssayCollect) (err error) {
	return mysql.CreateUserEssayCollect(newUserEssayCollect)
}

// 根据uid,eid获取文章收藏
func GetUserEssayCollect(uid int, eid int) (userEssayCollect *model.UserEssayCollect, err error) {
	return mysql.GetUserEssayCollect(uid, eid)
}

// get 用户全部收藏
func GetUserAllCollectByUid(uid int, page int, pageSize int) (userEssayCollects []model.UserEssayCollect, err error) {
	return mysql.GetUserAllCollectByUid(uid, page, pageSize)
}

// 取消收藏
func DeleteCollectById(id int) (int64, error) {
	return mysql.DeleteCollectById(id)
}

// 进行收藏
func UpdateUserEssayCollectNotToIs(uid int, eid int, fid int) (int64, error) {
	return mysql.UpdateUserEssayCollectNotToIs(uid, eid, fid)
}
