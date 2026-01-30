package service

import (
	"bishe/internal/app/knowledge_sharing/dao/mysql"
	"bishe/internal/app/knowledge_sharing/model"
)

// create 用户对文章收藏
func CreateUserEssayCollect(newUserEssayCollect *model.UserEssayCollect) (err error) {
	return mysql.CreateUserEssayCollect(newUserEssayCollect)
}

// 根据uid,eid获取文章点赞
func GetUserEssayCollect(uid int, eid int) (userEssayCollect *model.UserEssayCollect, err error) {
	return mysql.GetUserEssayCollect(uid, eid)
}
