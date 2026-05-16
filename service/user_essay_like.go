package service

import (
	"bishe/dao/mysql"
	"bishe/model"
)

// create 用户对文章点赞
func CreateUserEssayLike(newUserEssayLike *model.UserEssayLike) (err error) {
	return mysql.CreateUserEssayLike(newUserEssayLike)
}

// get 用户全部点赞
func GetUserAllLikeEssayByUid(uid int, page int, pageSize int) (userEssays []model.UserEssay, err error) {
	eids, err := mysql.GetUserAllLikeEssayByUid(uid, page, pageSize)
	if err != nil {
		return
	}

	if len(eids) == 0 {
		return
	}
	var essays []model.Essay
	essays, err = mysql.GetEssayByEids(eids)
	if err != nil {
		return
	}

	if len(essays) == 0 {
		return
	}

	uids := make([]int, 0)
	for _, v := range essays {
		uids = append(uids, v.AuthorId)
	}

	userMap, err := mysql.GetUserMapByUids(uids)
	if err != nil {
		return
	}

	if len(userMap) == 0 {
		return
	}
	userEssays = make([]model.UserEssay, 0)

	for _, v := range essays {
		var userEssay model.UserEssay
		if _, ok := userMap[v.AuthorId]; ok {
			userEssay.Author = userMap[v.AuthorId]
			userEssay.Essay = v
		}
		userEssays = append(userEssays, userEssay)
	}

	return
}

// 根据uid,eid获取文章点赞
func GetUserEssayLike(uid int, eid int) (userEssayLike *model.UserEssayLike, err error) {
	return mysql.GetUserEssayLike(uid, eid)
}

// 根据eids获取essayLikeMap
func GetUserEssayLikeMapByEids(uid int, eids []int) (essayLikeMap map[int]model.UserEssayLike, err error) {
	return mysql.GetUserEssayLikeMapByEids(uid, eids)
}

// 取消喜欢
func UpdateUserEssayLikeIsToNot(uid int, eid int) (int64, error) {
	return mysql.UpdateUserEssayLikeIsToNot(uid, eid)
}

// 进行喜欢
func UpdateUserEssayLikeNotToIs(uid int, eid int) (int64, error) {
	return mysql.UpdateUserEssayLikeNotToIs(uid, eid)
}

// // 更新 用户对文章点赞
// func UpdateUserEssayInteractLike(userEssayInteract *model.UserEssayLike, uid int, circleId int, essayId int) (result *gorm.DB) {
// 	return mysql.UpdateUserEssayInteractLike(userEssayInteract, uid, circleId, essayId)
// }

// func UpdateUserEssayInteractCollect(userEssayInteract *model.UserEssayLike, uid int, circleId int, essayId int, favorite string) (result *gorm.DB) {
// 	return mysql.UpdateUserEssayInteractCollect(userEssayInteract, uid, circleId, essayId, favorite)
// }
