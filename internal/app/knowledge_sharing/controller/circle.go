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

	// 用户创建圈子之前，判断isdelete

	// 删除唯一键，通过获取判断只有一个
	circle, err := service.GetCircleByTitle(title)
	if err != nil {
		service.Logger.Error("GetCircleByTitle err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	// 有正常状态返回，没有就创建

	if circle != nil {
		MakeApiResponseError(c, CODE_CIRCLE_EXIST)
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
		CircleStatus:  model.CIRCLE_STATUS_NORMAL,
		IsDeleted:     model.CIRCLE_NOT_DELETED,
	}

	err = service.CreateCircle(newCircle)
	if err != nil {
		service.Logger.Error("CreateCircle err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	MakeApiResponseSuccessDefault(c)

}

// 更新圈子信息
func UpdateCircleHandler(c *gin.Context) {
	// cid
	cidStr := c.PostForm("cid")
	if cidStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	cid, err := strconv.Atoi(cidStr)
	if err != nil {
		MakeApiResponseErrorDefault(c)
		return
	}

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
		MakeApiResponseErrorDefault(c)
		return
	}

	if price > model.CIRCLE_MAX_PRICE {
		MakeApiResponseError(c, CODE_CIRCLE_PRICE_INVASLID)
		return
	}

	//根据cid获取圈子
	circle, err := service.GetCircleByCid(cid)
	if err != nil {
		service.Logger.Error("GetCircleByCid err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if circle == nil {
		MakeApiResponseError(c, CODE_CIRCLE_NOT_EXIST)
		return
	}

	//更新圈子信息
	affectRows, err := service.UpdateCircleByCid(cid, title, price, introduction)
	if err != nil || affectRows == 0 {
		service.Logger.Error("UpdateCircleByCid err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	MakeApiResponseSuccessDefault(c)
}

// 获取圈子列表
func GetAllCircleHandler(c *gin.Context) {
	pageStr := c.Query("page")
	page := GetPage(pageStr)
	pageSize := 10

	//获取全部circle，按joinnum倒叙
	circles, err := service.GetCircleAllByJoinNum(page, pageSize)
	if err != nil {
		service.Logger.Error("GetCircleAllByJoinNum", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if circles == nil {
		circles = make([]model.Circle, 0)
	}

	data := map[string]interface{}{
		"circles": circles,
	}

	MakeApiResponseSuccess(c, data)
}

// 获取圈子详情
func GetCircleHandler(c *gin.Context) {
	cidStr := c.Query("cid")
	if cidStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	cid, err := strconv.Atoi(cidStr)
	if err != nil {
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

	//是否已加入
	isJoin := false

	//如果用户已加入圈子
	if userCircleJoin != nil {
		isJoin = true
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
		"is_join": isJoin,
		"circle":  circle,
	})
}

// 获取用户创建的圈子列表
func GetUserCreateCircleHandler(c *gin.Context) {
	pageStr := c.Query("page")
	page := GetPage(pageStr)
	pagesize := 5

	uid, _ := service.GetUserFromCookie(c)
	if uid == 0 {
		MakeApiResponseError(c, CODE_USER_NOT_LOGIN)
		return
	}

	//获取用户创建的圈子
	circles, err := service.GetUserCreateCircleByUid(uid, page, pagesize)
	if err != nil {
		service.Logger.Error("GetUserCreateCircleByUid", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	//如果circles为nil，make保证circles不为空
	if circles == nil {
		circles = make([]model.Circle, 0)
	}

	data := map[string]interface{}{
		"circles": circles,
	}

	MakeApiResponseSuccess(c, data)

}

// 获取用户已加入圈子的列表
func GetUserJoinCircleHandler(c *gin.Context) {
	pageStr := c.Query("page")
	page := GetPage(pageStr)

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

	//make保证json化时为[],而不是null
	if circles == nil {
		circles = make([]model.Circle, 0)
	}

	MakeApiResponseSuccess(c, map[string]interface{}{
		"circles": circles,
	})
}

// 获取付费圈子排行
func GetChargeCircleRankHandler(c *gin.Context) {
	pageStr := c.Query("page")
	page := GetPage(pageStr)

	pagesize := 10

	//获取付费circle，按joinnum倒叙
	circles, err := service.GetCircleAllChargeOrderByJoinNum(page, pagesize)
	if err != nil {
		service.Logger.Error("GetCircleAllChargeOrderByJoinNum", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	// nil
	if circles == nil {
		circles = make([]model.Circle, 0)
	}

	MakeApiResponseSuccess(c, map[string]interface{}{
		"circles": circles,
	})
}

// 获取免费圈子排行
func GetFreeCircleRankHandler(c *gin.Context) {
	pageStr := c.Query("page")
	page := GetPage(pageStr)

	pagesize := 10

	//获取免费circle，按join num 倒叙
	circles, err := service.GetCricleAllFreeOrderByJoinNum(page, pagesize)
	if err != nil {
		service.Logger.Error("GetCricleAllFreeOrderByJoinNum", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	// nil
	if circles == nil {
		circles = make([]model.Circle, 0)
	}

	MakeApiResponseSuccess(c, map[string]interface{}{
		"circles": circles,
	})
}
