package controller

import (
	"bishe/internal/app/knowledge_sharing/model"
	"bishe/internal/app/knowledge_sharing/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 获取通知列表
func GetNoticeListHandler(c *gin.Context) {
	uid := service.GetUidFromContext(c)

	pageStr := c.Query("page")
	page := GetPage(pageStr)
	pageSize := 10

	//获取通知列表
	notices, err := service.GetNoticeList(uid, page, pageSize)
	if err != nil {
		service.Logger.Error("GetNoticeList", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if len(notices) == 0 {
		notices = make([]model.Notice, 0)

	}

	data := map[string]interface{}{
		"notices": notices,
	}

	MakeApiResponseSuccess(c, data)
}

// 获取某类型通知列表
func GetNoticeListByTypeHandler(c *gin.Context) {
	uid := service.GetUidFromContext(c)

	typeStr := c.Query("type")
	if typeStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	typei, err := strconv.Atoi(typeStr)
	if err != nil {
		MakeApiResponseErrorParams(c)
		return
	}

	pageStr := c.Query("page")
	page := GetPage(pageStr)
	pageSize := 10

	//根据类型获取通知列表
	notices, err := service.GetNoticeListByType(uid, typei, page, pageSize)
	if err != nil {
		service.Logger.Error("GetNoticeListByType", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if len(notices) == 0 {
		notices = make([]model.Notice, 0)
		data := map[string]interface{}{
			"notices": notices,
		}

		MakeApiResponseSuccess(c, data)
		return
	}

	var uids []int
	for _, v := range notices {
		uids = append(uids, v.NoticeUid)
	}

	//根据uids获取userMap
	userMap, err := service.GetUserMapByUids(uids)
	if err != nil {
		service.Logger.Error("GetUserMapByUids", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if len(userMap) == 0 {
		service.Logger.Error("GetUserMapByUids len(userMap) == 0")
		MakeApiResponseErrorDefault(c)
		return
	}

	userNotices := make([]model.UserNotice, 0)

	for _, v := range notices {
		vUid := v.NoticeUid

		vUser, ok := userMap[vUid]
		if !ok {
			service.Logger.Error("set uids err")
			MakeApiResponseErrorDefault(c)
			return
		}

		var userNotice model.UserNotice

		updateAt := v.UpdateAt.Format("2006-01-02 15:04:05")
		userNotice.Uid = vUid
		userNotice.Name = vUser.Name
		userNotice.Content = v.Content
		userNotice.UpdateAt = updateAt
		userNotice.Type = v.Type
		userNotices = append(userNotices, userNotice)
	}

	data := map[string]interface{}{
		"userNotices": userNotices,
	}

	MakeApiResponseSuccess(c, data)
}
