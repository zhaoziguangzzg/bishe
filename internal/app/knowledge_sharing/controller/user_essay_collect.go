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
	uid, _ := service.GetUserFromCookie(c)
	if uid == 0 {
		MakeApiResponseError(c, CODE_USER_NOT_LOGIN)
		return
	}

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

	// TODO collect_status 直查询有效的
	collect, err := service.GetUserEssayCollect(uid, eid)
	if err != nil {
		service.Logger.Error("GetUserEssayCollect", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if collect == nil {
		createTime := time.Now()
		newUserEssayCollect := &model.UserEssayCollect{ //其中包含自动生成的id
			UserId:        uid,
			EssayId:       eid,
			FavoriteId:    fid,
			CreateAt:      &createTime,
			UpdateAt:      &createTime,
			CollectStatus: model.COLLECT_STATUS_NORMAL,
		}

		//TODO 去唯一键，可以重复

		err = service.CreateUserEssayCollect(newUserEssayCollect)
		if err != nil {
			service.Logger.Error("CreateUserEssayCollect err", zap.Error(err))
			MakeApiResponseErrorDefault(c)
			return
		}

		MakeApiResponseSuccessDefault(c)
		return

	} else {
		//未收藏转为收藏
		affectRows, err := service.UpdateUserEssayCollectNotToIs(uid, eid, fid)
		if err != nil || affectRows == 0 {
			service.Logger.Error("UpdateUserEssayCollectNotToIs err", zap.Error(err))
			MakeApiResponseErrorDefault(c)
			return
		}

		MakeApiResponseSuccessDefault(c)
		return
	}
}

// 取消用户收藏
func CancelEssayCollectHandler(c *gin.Context) {
	uid, _ := service.GetUserFromCookie(c)
	if uid == 0 {
		MakeApiResponseError(c, CODE_USER_NOT_LOGIN)
		return
	}

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

	MakeApiResponseSuccessDefault(c)

}

// 获取用户文章是否收藏
func GetEssayCollectHandler(c *gin.Context) {
	uid, _ := service.GetUserFromCookie(c)
	if uid == 0 {
		MakeApiResponseError(c, CODE_USER_NOT_LOGIN)
		return
	}

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
