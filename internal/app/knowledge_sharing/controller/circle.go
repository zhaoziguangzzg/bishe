package controller

import (
	"bishe/internal/app/knowledge_sharing/model"
	"bishe/internal/app/knowledge_sharing/service"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 圈子
func AddCircleHandler(c *gin.Context) {

	// 从表单中获取圈子信息
	title := c.PostForm("title")
	priceStr := c.PostForm("price")

	// 数据验证
	if title == "" {
		MakeApiResponseError(c, CODE_PARAMS_ERROR)
		return
	}

	if priceStr == "" {
		MakeApiResponseError(c, CODE_PARAMS_ERROR)
		return
	}

	price, err := strconv.Atoi(priceStr)
	if err != nil {
		service.Logger.Error("priceStrAtoi err", zap.Error(err))
	}

	createTime := time.Now()

	// 构造圈子对象
	newCircle := &model.Circle{ //其中包含自动生成的id
		Title:         title,
		Price:         price,
		CircleOwnerId: UserId,
		CreateAt:      &createTime,
		UpdateAt:      &createTime,
		Status:        model.CircleNormal,
	}

	// 插入数据库
	err = service.CreateCircle(newCircle)
	if err != nil {
		service.Logger.Error("CreateCircle err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	// 返回成功响应
	MakeApiResponseSuccess(c, CODE_SUCCESS)
}

// 获取圈子
func GetCircleHandle(c *gin.Context) {
	chargeCircles, err := service.GetCircleAllCharge()
	if err != nil {
		service.Logger.Error("GetCircleAllCharge err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	freeCircles, err := service.GetCricleAllFree()
	if err != nil {
		service.Logger.Error("GetCricleAllFree err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	allCircles, err := service.GetCircleAllSortByJoinNum()
	if err != nil {
		service.Logger.Error("GetCircleAllSortByJoinNum err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	MakeApiResponse(c, CODE_SUCCESS, map[string]interface{}{
		"chargeCircles": chargeCircles,
		"freeCircles":   freeCircles,
		"allCircles":    allCircles,
	})

}
