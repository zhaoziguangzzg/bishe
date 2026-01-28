package controller

import (
	"bishe/internal/app/knowledge_sharing/model"
	"bishe/internal/app/knowledge_sharing/service"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 用户在文章的评论
func AddUserEssayCommentHandle(c *gin.Context) {
	uid, _ := service.GetUserFromCookie(c)
	if uid == 0 {
		MakeApiResponseError(c, CODE_USER_NOT_LOGIN)
		return
	}

	eid := c.GetInt("eid")
	if eid == 0 {
		service.Logger.Error("geteid err", zap.String("err", "get eid"))
		MakeApiResponseErrorDefault(c)
		return
	}

	content := c.PostForm("content")

	contentLen := len(content)
	if contentLen > model.COMMENT_MAX_CONTENT || contentLen == 0 {
		MakeApiResponseError(c, CODE_COMMENT_CONTENT_LEN_INVASLID)
		return
	}

	createTime := time.Now()

	newUserEssayComment := &model.UserEssayComment{ //其中包含自动生成的id
		UserId:   uid,
		EssayId:  eid,
		Content:  content,
		CreateAt: &createTime,
		UpdateAt: &createTime,
	}

	err := service.CreateUserEssayComment(newUserEssayComment)
	if err != nil {
		service.Logger.Error("CreateUserEssayComment err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	MakeApiResponseSuccessDefault(c)
}

// 获取文章评论列表
func GetUserEssayCommentHandle(c *gin.Context) {
	eid := c.GetInt("eid")
	if eid == 0 {
		service.Logger.Error("GetInt eid err", zap.String("err", "get eid err"))
		MakeApiResponseErrorDefault(c)
		return
	}

	page := c.GetInt("page")
	if page == 0 {
		service.Logger.Error("GetInt page err", zap.String("err", "get page err"))
		MakeApiResponseErrorDefault(c)
		return
	}

	pageSize := 10

	comments, err := service.GetEssayAllComment(eid, page, pageSize)
	if err != nil {
		service.Logger.Error("GetEssayAllComment", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	MakeApiResponseSuccess(c, map[string]interface{}{
		"comments": comments,
	})
}

// 获取用户全部评论列表
func GetUserAllCommentHandler(c *gin.Context) {
	uid, _ := service.GetUserFromCookie(c)
	if uid == 0 {
		MakeApiResponseError(c, CODE_USER_NOT_LOGIN)
		return
	}

	page := c.GetInt("page")
	if page == 0 {
		service.Logger.Error("GetInt page err", zap.String("err", "get page err"))
		MakeApiResponseErrorDefault(c)
		return
	}

	pageSize := 10

	//获取用户全部comment
	comments, err := service.GetEssayAllCommentByUid(uid, page, pageSize)
	if err != nil {
		service.Logger.Error("GetEssayAllCommentByUid", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	MakeApiResponseSuccess(c, map[string]interface{}{
		"comments": comments,
	})
}

// 删除评论
func DeletedCommentByUpdateIsDeletedHandler(c *gin.Context) {
	commentId := c.GetInt("comment_id")
	if commentId == 0 {
		service.Logger.Error("GetInt comment_id err", zap.String("err", "get comment_id"))
		MakeApiResponseErrorDefault(c)
		return
	}

	//更新isDeleted
	affectRows, err := service.UpdateIsDeletedByCommentId(commentId)
	if err != nil || affectRows == 0 {
		service.Logger.Error("UpdateIsDeletedByCommentId err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

}
