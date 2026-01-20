package service

import (
	"bishe/internal/app/knowledge_sharing/dao/mysql"
	"bishe/internal/app/knowledge_sharing/model"

	"gorm.io/gorm"
)

// create 用户对文章点赞收藏
func CreateUserEssayInteract(newUserEssayInterAct *model.UserEssayInteract) (err error) {
	return mysql.CreateUserEssayInteract(newUserEssayInterAct)
}

// get 用户对文章点赞收藏
func GetUserEssayInteract(uid int, circleId int, essayId int) (userEssayInteract *model.UserEssayInteract, err error) {
	return mysql.GetUserEssayInteract(uid, circleId, essayId)
}

// 更新 用户对文章点赞
func UpdateUserEssayInteractLike(userEssayInteract *model.UserEssayInteract, uid int, circleId int, essayId int) (result *gorm.DB) {
	return mysql.UpdateUserEssayInteractLike(userEssayInteract, uid, circleId, essayId)
}

func UpdateUserEssayInteractCollect(userEssayInteract *model.UserEssayInteract, uid int, circleId int, essayId int, favorite string) (result *gorm.DB) {
	return mysql.UpdateUserEssayInteractCollect(userEssayInteract, uid, circleId, essayId, favorite)
}
