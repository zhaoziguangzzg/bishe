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

// get 用户收藏夹全部收藏
func GetUserAllCollectByUidFid(uid int, fid int, page int, pageSize int) (essays []model.Essay, err error) {
	eids, err := mysql.GetUserAllCollectByUidFid(uid, fid, page, pageSize)
	if err != nil {
		return
	}

	if eids == nil {
		return
	}

	essays, err = mysql.GetEssayByEids(eids)
	if err != nil {
		return
	}
	return
}

// 取消收藏
func UpdateUserEssayCollectIsToNot(uid int, eid int) (int64, error) {
	return mysql.UpdateUserEssayCollectIsToNot(uid, eid)
}

// 进行收藏
func UpdateUserEssayCollectNotToIs(uid int, eid int, fid int) (int64, error) {
	return mysql.UpdateUserEssayCollectNotToIs(uid, eid, fid)
}
