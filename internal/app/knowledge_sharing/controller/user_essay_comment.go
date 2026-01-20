package controller

import (
	"bishe/internal/app/knowledge_sharing/model"
	"bishe/internal/app/knowledge_sharing/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func AddUserEssayCommentHandle(c *gin.Context) {
	uidStr := c.PostForm("user_id")
	circleIdStr := c.PostForm("circle_id")
	essayIdStr := c.PostForm("essay_id")
	comment := c.PostForm("comment")

	if uidStr == "" {
		MakeApiResponseError(c, CODE_PARAMS_ERROR)
		return
	}

	if circleIdStr == "" {
		MakeApiResponseError(c, CODE_PARAMS_ERROR)
		return
	}

	if essayIdStr == "" {
		MakeApiResponseError(c, CODE_PARAMS_ERROR)
		return
	}

	uid, err := strconv.Atoi(uidStr)
	if err != nil {
		MakeApiResponseError(c, CODE_PARAMS_ERROR)
		return
	}

	circleId, err := strconv.Atoi(circleIdStr)
	if err != nil {
		MakeApiResponseError(c, CODE_PARAMS_ERROR)
		return
	}

	essayId, err := strconv.Atoi(essayIdStr)
	if err != nil {
		MakeApiResponseError(c, CODE_PARAMS_ERROR)
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

	MakeApiResponseSuccess(c, newUserEssayComment)
}

func GetUserEssayCommentHandle(c *gin.Context) {
	circleIdStr := c.PostForm("circle_id")
	essayIdStr := c.PostForm("essay_id")

	if circleIdStr == "" {
		MakeApiResponseError(c, CODE_PARAMS_ERROR)
		return
	}

	if essayIdStr == "" {
		MakeApiResponseError(c, CODE_PARAMS_ERROR)
		return
	}

	circleId, err := strconv.Atoi(circleIdStr)
	if err != nil {
		MakeApiResponseError(c, CODE_PARAMS_ERROR)
		return
	}

	essayId, err := strconv.Atoi(essayIdStr)
	if err != nil {
		MakeApiResponseError(c, CODE_PARAMS_ERROR)
		return
	}

	comments, err := service.GetCircleEssayComment(circleId, essayId)
	if err != nil {
		MakeApiResponseError(c, CODE_PARAMS_ERROR)
		return
	}

	MakeApiResponseSuccess(c, comments)
}
