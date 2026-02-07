package controller

import (
	"bishe/internal/app/knowledge_sharing/model"
	"bishe/internal/app/knowledge_sharing/service"
	"strconv"
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

	eidStr := c.PostForm("eid")
	if eidStr == "" {
		service.Logger.Error("Geteid err", zap.String("err", "get eid err"))
		MakeApiResponseErrorParams(c)
		return
	}

	eid, err := strconv.Atoi(eidStr)
	if err != nil {
		service.Logger.Error("Atoi eidStr err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
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

	//TODO 去唯一键

	err = service.CreateUserEssayComment(newUserEssayComment)
	if err != nil {
		service.Logger.Error("CreateUserEssayComment err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	MakeApiResponseSuccessDefault(c)
}

// 删除评论
func DeletedCommentByUpdateIsDeletedHandler(c *gin.Context) {
	commentIdStr := c.PostForm("comment_id")
	if commentIdStr == "" {
		service.Logger.Error("GetcommentId err", zap.String("err", "get commentId err"))
		MakeApiResponseErrorParams(c)
		return
	}

	commentId, err := strconv.Atoi(commentIdStr)
	if err != nil {
		service.Logger.Error("Atoi commentIdStr err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
	}

	//更新isDeleted
	affectRows, err := service.UpdateIsDeletedByCommentId(commentId)
	if err != nil || affectRows == 0 {
		service.Logger.Error("UpdateIsDeletedByCommentId err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	MakeApiResponseSuccessDefault(c)

}

// 获取文章评论列表
func GetEssayAllCommentHandle(c *gin.Context) {
	eidStr := c.Query("eid")
	if eidStr == "" {
		service.Logger.Error("Geteid err", zap.String("err", "get eid err"))
		MakeApiResponseErrorParams(c)
		return
	}

	eid, err := strconv.Atoi(eidStr)
	if err != nil {
		service.Logger.Error("Atoi eidStr err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
	}

	page := c.GetInt("page")
	if page < 1 {
		page = 1
	}

	pageSize := 10

	comments, err := service.GetEssayAllComment(eid, page, pageSize)
	if err != nil {
		service.Logger.Error("GetEssayAllComment", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if comments == nil {
		comments = make([]model.UserEssayComment, 0)
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
	if page < 1 {
		page = 1
	}

	pageSize := 10

	//获取用户全部评论文章
	commentEssays, err := service.GetUserAllCommentIdByUid(uid, page, pageSize)
	if err != nil {
		service.Logger.Error("GetEssayAllCommentByUid", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if commentEssays == nil {
		commentEssays = make([]model.CommentEssay, 0)
	}

	MakeApiResponseSuccess(c, commentEssays)
}
