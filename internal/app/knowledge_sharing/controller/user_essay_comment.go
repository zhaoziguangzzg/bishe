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
	uidStr, err := c.Cookie("uid")
	if err != nil {
		return
	}

	accountStr, err := c.Cookie("username")
	if err != nil {
		return
	}

	circleIdStr := c.PostForm("circle_id")
	essayIdStr := c.PostForm("essay_id")

	comment := c.PostForm("comment")

	uid, _ := strconv.Atoi(uidStr)
	if uid == 0 {
		MakeApiResponseErrorParams(c)
		return
	}

	account, _ := strconv.Atoi(accountStr)
	if account == 0 {
		MakeApiResponseErrorParams(c)
		return
	}

	circleId, _ := strconv.Atoi(circleIdStr)
	if circleId == 0 {
		MakeApiResponseErrorParams(c)
		return
	}

	essayId, _ := strconv.Atoi(essayIdStr)
	if essayId == 0 {
		MakeApiResponseErrorParams(c)
		return
	}

	if len(comment) > 200 {
		//评论过长
		MakeApiResponseError(c, CODE_COMMENT_TOO_LONG)
		return
	}

	newUserEssayComment := &model.UserEssayComment{ //其中包含自动生成的id
		UserId:   uid,
		CircleId: circleId,
		EssayId:  essayId,
		Comment:  comment,
	}

	err = service.CreateUserEssayComment(newUserEssayComment)
	if err != nil {
		service.Logger.Error("CreateUserEssayComment err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	createTime := time.Now()
	content := "对文章" + essayIdStr + "评论成功"
	err = service.MakeAndSendNotice(0, UserAccount, content, createTime)
	if err != nil {
		service.Logger.Error("MakeAndSendNotice err", zap.Int("uid", uid), zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	MakeApiResponseSuccess(c, newUserEssayComment)
}

// 获取文章评论列表
func GetUserEssayCommentHandle(c *gin.Context) {
	circleId := c.GetInt("circle_id")
	essayId := c.GetInt("essay_id")
	page := c.GetInt("page")
	pageSize := 10

	if page <= 0 {
		page = 1
	}

	if circleId == 0 || essayId == 0 {
		MakeApiResponseErrorParams(c)
		return
	}

	comments, err := service.GetCircleEssayComment(circleId, essayId, page, pageSize)
	if err != nil {
		MakeApiResponseError(c, CODE_PARAMS_ERROR)
		return
	}

	MakeApiResponseSuccess(c, comments)
}
