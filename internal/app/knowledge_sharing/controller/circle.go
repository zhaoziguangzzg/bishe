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
	introduction := c.PostForm("introduction")

	// 数据验证
	titleLen := len(title)
	if titleLen > model.CIRCLE_MAX_TITLE || titleLen == 0 {
		MakeApiResponseError(c, CODE_CIRCLE_TITLE_LEN_INVASLID)
		return
	}

	introductionLen := len(introduction)
	if introductionLen > model.CIRCLE_MAX_INTRODUCTION || introductionLen == 0 {
		MakeApiResponseError(c, CODE_CIRCLE_INTRODUCTION_LEN_INVASLID)
		return
	}

	price, err := strconv.Atoi(priceStr)
	if err != nil {
		service.Logger.Error("priceStrAtoi err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if price > model.CIRCLE_MAX_PRICE {
		MakeApiResponseError(c, CODE_CIRCLE_PRICE_INVASLID)
		return
	}

	uid, _ := service.GetUserFromCookie(c)
	if uid == 0 {
		MakeApiResponseError(c, CODE_USER_NOT_LOGIN)
		return
	}

	createTime := time.Now()
	// 构造圈子对象
	newCircle := &model.Circle{ //其中包含自动生成的id
		Title:         title,
		Price:         price,
		Introduction:  introduction,
		CircleOwnerId: uid,
		CreateAt:      &createTime,
		UpdateAt:      &createTime,
	}

	// 插入数据库
	err = service.CreateCircle(newCircle)
	if err != nil {
		service.Logger.Error("CreateCircle err", zap.Error(err))
		MakeApiResponseError(c, CODE_CIRCLE_EXIST)
		return
	}

	// 返回成功响应
	MakeApiResponseSuccessDefault(c)
}

// 获取圈子列表
func GetAllCircleHandler(c *gin.Context) {
	page := c.GetInt("page")
	if page == 0 {
		service.Logger.Error("GetInt page err", zap.String("err", "get page err"))
		MakeApiResponseErrorDefault(c)
		return
	}

	pagesize := 10

	//获取全部circle，按joinnum倒叙
	circles, err := service.GetCircleAllByJoinNum(page, pagesize)
	if err != nil {
		service.Logger.Error("GetCircleAllByJoinNum", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	MakeApiResponseSuccess(c, map[string]interface{}{
		"circles": circles,
	})
}

// 获取圈子详情
func GetCircleHandler(c *gin.Context) {
	cid := c.GetInt("cid")
	if cid == 0 {
		service.Logger.Error("GetInt cid err", zap.String("err", "get cid err"))
		MakeApiResponseErrorDefault(c)
		return
	}

	uid, _ := service.GetUserFromCookie(c)
	if uid == 0 {
		MakeApiResponseError(c, CODE_USER_NOT_LOGIN)
		return
	}

	//根据uidcid获取用户加入
	userCircleJoin, err := service.GetUserJoinCircleByUidCid(uid, cid)
	if err != nil {
		service.Logger.Error("GetUserJoinCircleByUidCid", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	//如果用户已加入圈子
	if userCircleJoin != nil {
		//进入圈子
		return
	}

	//用户未加入圈子
	//根据cid获取圈子
	circle, err := service.GetCircleByCid(cid)
	if err != nil {
		service.Logger.Error("GetCircleByCid", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if circle == nil {
		MakeApiResponseError(c, CODE_CIRCLE_NOT_EXIST)
		return
	}

	MakeApiResponseSuccess(c, map[string]interface{}{
		"circle": circle,
	})
}

// 获取用户创建的圈子列表
func GetUserCreateCircleHandler(c *gin.Context) {
	page := c.GetInt("page")
	if page == 0 {
		service.Logger.Error("GetInt page err", zap.String("err", "get page err"))
		MakeApiResponseErrorDefault(c)
		return
	}

	pagesize := 5

	uid, _ := service.GetUserFromCookie(c)
	if uid == 0 {
		MakeApiResponseError(c, CODE_USER_NOT_LOGIN)
		return
	}

	circles, err := service.GetUserCreateCircleByUid(uid, page, pagesize)
	if err != nil {
		service.Logger.Error("GetUserCreateCircleByUid", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if circles == nil {
		MakeApiResponseError(c, CODE_USER_NOT_CREATE_CIRCLE)
		return
	}

	MakeApiResponseSuccess(c, map[string]interface{}{
		"circles": circles,
	})

}

// 获取用户已加入圈子的列表
func GetUserJoinCricleHandler(c *gin.Context) {
	page := c.GetInt("page")
	if page == 0 {
		service.Logger.Error("GetInt page err", zap.String("err", "get page err"))
		MakeApiResponseErrorDefault(c)
		return
	}

	pagesize := 5

	uid, _ := service.GetUserFromCookie(c)
	if uid == 0 {
		MakeApiResponseError(c, CODE_USER_NOT_LOGIN)
		return
	}

	//根据uid获取用户加入的圈子列表
	circles, err := service.GetUserJoinCircleListByUid(uid, page, pagesize)
	if err != nil {
		service.Logger.Error("GetUserJoinCircleListByUid", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	MakeApiResponseSuccess(c, map[string]interface{}{
		"circle": circle,
	})
}

// 更新圈子信息
func UpdateCircleHandler(c *gin.Context) {
	// cid

	// title

	// price

	// introduction

}
