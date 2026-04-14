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
func GetUserAllCommentIdByUid(uid int, page int, pageSize int) (commentEssays []model.CommentEssay, err error) {
	comments, err := mysql.GetUserAllCommentIdByUid(uid, page, pageSize)
	if err != nil {
		return
	}

	var eids []int
	for _, v := range comments {
		eids = append(eids, v.EssayId)
	}

	var essays []model.Essay

	essays, err = mysql.GetEssayByEids(eids)
	if err != nil {
		return
	}

	//根据eids的essaymap
	essayMap := make(map[int]model.Essay, 0)
	for _, v := range essays {
		essayMap[v.Id] = v
	}

	for _, v := range comments {
		var commentEssay model.CommentEssay
		essay, ok := essayMap[v.EssayId]
		if !ok {
			return
		}

		commentEssay.Essay = essay
		commentEssay.Comment = v
		commentEssays = append(commentEssays, commentEssay)
	}

	// 组装commentEssays

	return

}

// update isdeleted
func UpdateIsDeletedByCommentId(commentId int) (int64, error) {
	return mysql.UpdateIsDeletedByCommentId(commentId)
}
