package mysql

import "bishe/internal/app/knowledge_sharing/model"

//create 用户对文章评论
func CreateUserEssayComment(newUserEssayComment *model.UserEssayComment) (err error) {
	err = DB.Model(&model.UserEssayComment{}).Create(newUserEssayComment).Error
	return
}

//get 文章的评论
func GetCircleEssayComment(circleId int, essayId int) (comments *model.UserEssayComment, err error) {
	err = DB.Model(&model.UserEssayComment{}).Where("circle_id=? and essay_id", circleId, essayId).Find(comments).Error
	return
}
