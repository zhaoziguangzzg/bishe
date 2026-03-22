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
func GetEssayAllComment(eid int, page int, pageSize int) (comments []model.UserEssayComment, err error) {
	return mysql.GetEssayAllComment(eid, page, pageSize)
}

// get 用户全部comment
func GetUserAllCommentIdByUid(uid int, page int, pageSize int) (essays []model.Essay, err error) {
	comments, err := mysql.GetUserAllCommentIdByUid(uid, page, pageSize)
	if err != nil {
		return
	}

	var eids []int
	for _, v := range comments {
		eids = append(eids, v.EssayId)
	}

	essays, err = mysql.GetEssayByEids(eids)
	if err != nil {
		return
	}

	// 组装commentEssays

	return

}

// update isdeleted
func UpdateIsDeletedByCommentId(commentId int) (int64, error) {
	return mysql.UpdateIsDeletedByCommentId(commentId)
}
