package controller

import (
	"bishe/model"
	"bishe/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 获取通知列表
func GetNoticeListHandler(c *gin.Context) {
	uid := service.GetUidFromContext(c)
	noticeTypeStr := c.Query("type")
	noticeType, _ := strconv.Atoi(noticeTypeStr)

	pageStr := c.Query("page")
	page := GetPage(pageStr)
	pageSize := 10

	//获取通知列表
	var notices []model.Notice
	var err error
	if noticeType > 0 {
		notices, err = service.GetNoticeListByType(uid, noticeType, page, pageSize)
	} else {
		notices, err = service.GetNoticeList(uid, page, pageSize)
	}

	if err != nil {
		service.Logger.Error("GetNoticeList", zap.Error(err), zap.Int("uid", uid), zap.String("type", noticeTypeStr))
		MakeApiResponseErrorDefault(c)
		return
	}

	if len(notices) == 0 {
		notices = make([]model.Notice, 0)
	}

	data := map[string]interface{}{
		"notices":      notices,
		"typeNameList": service.GetNoticeTypeNameList(),
	}

	MakeApiResponseSuccess(c, data)
}
