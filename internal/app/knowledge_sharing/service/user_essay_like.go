package service

import (
	"bishe/internal/app/knowledge_sharing/dao/mysql"
	"bishe/internal/app/knowledge_sharing/model"
)

// create 用户对文章点赞
func CreateUserEssayLike(newUserEssayLike *model.UserEssayLike) (err error) {
	return mysql.CreateUserEssayLike(newUserEssayLike)
}

// get 用户对文章点赞收藏
func GetUserEssayInteract(uid int, circleId int, essayId int) (userEssayInteract *model.UserEssayLike, err error) {
	return mysql.GetUserEssayInteract(uid, circleId, essayId)
}

// get 用户全部点赞
func GetUserAllLikeByUid(uid int, page int, pageSize int) (userEssayInteracts []model.UserEssayLike, err error) {
	return GetUserAllLikeByUid(uid, page, pageSize)
}

// 根据uid,eid获取文章点赞
func GetUserEssayLike(uid int, eid int) (userEssayLike *model.UserEssayLike, err error) {
	return mysql.GetUserEssayLike(uid, eid)
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
