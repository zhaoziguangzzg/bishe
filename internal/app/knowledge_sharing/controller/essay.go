package controller

import (
	"bishe/internal/app/knowledge_sharing/model"
	"bishe/internal/app/knowledge_sharing/service"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 文章
func AddEssayHandler(c *gin.Context) { //c

	// 从表单中获取用户信息
	title := c.PostForm("title")
	content := c.PostForm("content")
	circleIdStr := c.PostForm("circleId")

	// 数据验证
	if title == "" {
		MakeApiResponseError(c, CODE_PARAMS_ERROR)
		return
	}

	if content == "" {
		MakeApiResponseError(c, CODE_PARAMS_ERROR)
		return
	}

	if circleIdStr == "" {
		MakeApiResponseError(c, CODE_PARAMS_ERROR)
		return
	}

	circleId, err := strconv.Atoi(circleIdStr)
	if err != nil {
		service.Logger.Error("circleIdAtoi err", zap.Error(err))
		MakeApiResponseError(c, CODE_PARAMS_ERROR)
		return
	}

	createTime := time.Now()

	// 构造文章
	newEssay := &model.Essay{ //其中包含自动生成的id
		Title:       title,
		Content:     content,
		CircleId:    circleId,
		AuthorId:    UserId,
		CreateAt:    &createTime,
		UpdateAt:    &createTime,
		EssayStatus: model.EssayNormal,
	}

	// 插入数据库
	err = service.CreateEssay(newEssay)
	if err != nil {
		service.Logger.Error("CreateEssay err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	// 返回成功响应
	MakeApiResponseSuccess(c, CODE_SUCCESS)
}
