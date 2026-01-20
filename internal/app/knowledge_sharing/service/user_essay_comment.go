package service

import (
	"bishe/internal/app/knowledge_sharing/dao/mysql"
	"bishe/internal/app/knowledge_sharing/model"
)

// create 用户对文章评论
func CreateUserEssayComment(newUserEssayComment *model.UserEssayComment) (err error) {
	return mysql.CreateUserEssayComment(newUserEssayComment)
}

// get 文章的评论
func GetCircleEssayComment(circleId int, essayId int) (comments *model.UserEssayComment, err error) {
	return mysql.GetCircleEssayComment(circleId, essayId)
}
