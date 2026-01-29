package service

import (
	"bishe/internal/app/knowledge_sharing/dao/mysql"
	"bishe/internal/app/knowledge_sharing/model"
)

// create 用户对文章收藏
func CreateUserEssayCollect(newUserEssayCollect *model.UserEssayCollect) (err error) {
	return mysql.CreateUserEssayCollect(newUserEssayCollect)
}
