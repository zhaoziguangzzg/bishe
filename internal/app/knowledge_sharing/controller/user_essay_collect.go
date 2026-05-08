package controller

import (
	"bishe/internal/app/knowledge_sharing/model"
	"bishe/internal/app/knowledge_sharing/service"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 用户收藏的文章
func AddUserEssayCollectHandler(c *gin.Context) {
	uid := service.GetUidFromContext(c)

	eidStr := c.PostForm("eid")
	if eidStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	eid, err := strconv.Atoi(eidStr)
	if err != nil {
		MakeApiResponseErrorDefault(c)
		return
	}

	fidStr := c.PostForm("fid")
	if fidStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	fid, err := strconv.Atoi(fidStr)
	if err != nil {
		MakeApiResponseErrorDefault(c)
		return
	}

	collect, err := service.GetUserEssayCollect(uid, eid)
	if err != nil {
		service.Logger.Error("GetUserEssayCollect", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	createTime := time.Now()
	if collect != nil {
		MakeApiResponseError(c, CODE_COLLECT_EXIST)
		return
	}

	newUserEssayCollect := &model.UserEssayCollect{ //其中包含自动生成的id
		UserId:        uid,
		EssayId:       eid,
		FavoriteId:    fid,
		CreateAt:      &createTime,
		UpdateAt:      &createTime,
		CollectStatus: model.COLLECT_STATUS_NORMAL,
	}

	err = service.CreateUserEssayCollect(newUserEssayCollect)
	if err != nil {
		service.Logger.Error("CreateUserEssayCollect err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	typei := model.STAT_TYPE_COLLECT

	//添加或更新用户统计数
	err = service.StatInsertUpdate(uid, 1, typei, createTime)
	if err != nil {
		service.Logger.Error("StatInsertUpdate err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	//添加文章收藏数据详情
	err = service.StatDetailsInsert(uid, typei, model.STAT_DETAILS_STATUS_INCR, createTime)
	if err != nil {
		service.Logger.Error("StatDetailsInsert err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	//更新文章收藏数
	affectRows, err := service.UpdateEssayCollectNum(eid, 1)
	if err != nil || affectRows == 0 {
		service.Logger.Error("UpdateEssayCollectNum err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	MakeApiResponseSuccessDefault(c)
}

// 取消用户收藏
func CancelEssayCollectHandler(c *gin.Context) {
	uid := service.GetUidFromContext(c)

	eidStr := c.PostForm("eid")
	if eidStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	eid, err := strconv.Atoi(eidStr)
	if err != nil {
		MakeApiResponseErrorDefault(c)
		return
	}

	//取消收藏
	affectRows, err := service.UpdateUserEssayCollectIsToNot(uid, eid)
	if err != nil || affectRows == 0 {
		service.Logger.Error("UpdateUserEssayCollectIsToNot err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	//更新文章收藏数
	affectRows, err = service.UpdateEssayCollectNum(eid, -1)
	if err != nil || affectRows == 0 {
		service.Logger.Error("UpdateEssayCollectNum err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	typei := model.STAT_TYPE_COLLECT
	nowTime := time.Now()

	//添加或更新用户统计数
	err = service.StatInsertUpdate(uid, -1, typei, nowTime)
	if err != nil {
		service.Logger.Error("StatInsertUpdate err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	//更新用户统计数
	err = service.UpdateStatAndStatDetail(uid, typei, model.STAT_DETAILS_STATUS_DECR, nowTime)
	if err != nil {
		service.Logger.Error("UpdateStatAndStatDetail err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	MakeApiResponseSuccessDefault(c)

}

// 获取用户文章是否收藏
func GetEssayCollectHandler(c *gin.Context) {
	uid := service.GetUidFromContext(c)

	eidStr := c.Query("eid")
	if eidStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	eid, err := strconv.Atoi(eidStr)
	if err != nil {
		MakeApiResponseErrorDefault(c)
		return
	}

	// TODO 仅查询有效的
	collect, err := service.GetUserEssayCollect(uid, eid)
	if err != nil {
		service.Logger.Error("GetUserEssayCollect", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if collect == nil {
		MakeApiResponseError(c, CODE_COLLECT_NOT_EXIST)
		return
	}

	data := map[string]interface{}{
		"collect": collect,
	}

	MakeApiResponseSuccess(c, data)

}

// 获取用户收藏夹全部收藏
func GetUserAllCollectHandler(c *gin.Context) {
	// 获取收藏夹参数
	fidStr := c.Query("fid")
	if fidStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	fid, err := strconv.Atoi(fidStr)
	if err != nil {
		MakeApiResponseErrorDefault(c)
		return
	}

	uid := service.GetUidFromContext(c)

	pageStr := c.Query("page")
	page := GetPage(pageStr)

	pageSize := 10

	//获取用户收藏夹全部collect
	collects, err := service.GetUserAllCollectByUidFid(uid, fid, page, pageSize)
	if err != nil {
		service.Logger.Error("GetUserAllCollectByUidFid", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	MakeApiResponseSuccess(c, map[string]interface{}{
		"collects": collects,
	})
}
