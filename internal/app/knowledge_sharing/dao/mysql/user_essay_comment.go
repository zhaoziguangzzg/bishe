package mysql

import "bishe/internal/app/knowledge_sharing/model"

//create 用户对文章评论
func CreateUserEssayComment(newUserEssayComment *model.UserEssayComment) (err error) {
	err = DB.Model(&model.UserEssayComment{}).Create(newUserEssayComment).Error
	return
}

//get 文章的评论
func GetEssayAllComment(eid int, page int, pageSize int) (comments []model.UserEssayComment, err error) {
	offset := (page - 1) * pageSize
	err = DB.Model(&model.UserEssayComment{}).Where("essay_id and is_deleted=?", eid, model.COMMENT_NOT_DELETED).
		Order("id ASC").Offset(offset).Limit(pageSize).Find(&comments).Error
	return
}

//get 用户全部comment
func GetEssayAllCommentByUid(uid int, page int, pageSize int) (comments []model.UserEssayComment, err error) {
	offset := (page - 1) * pageSize
	err = DB.Model(&model.UserEssayComment{}).Where("user_id and is_deleted=?", uid, model.COMMENT_NOT_DELETED).
		Order("id ASC").Offset(offset).Limit(pageSize).Find(&comments).Error
	return
}

//update isdeleted
func UpdateIsDeletedByCommentId(commentId int) (int64, error) {
	result := DB.Model(&model.UserEssayComment{}).Where("id=? and is_deleted=?", commentId, model.COMMENT_NOT_DELETED).
		Update("is_deleted", model.COMMENT_IS_DELETED)
	return result.RowsAffected, result.Error
}
