package controller

import (
	"bishe/internal/app/knowledge_sharing/model"
	"bishe/internal/app/knowledge_sharing/service"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 获取用户全部数据详情
func GetUserStatDetailsListByTimeHandler(c *gin.Context) {
	uid, _ := service.GetUserFromCookie(c)
	if uid == 0 {
		MakeApiResponseError(c, CODE_USER_NOT_LOGIN)
		return
	}

	stimeStr := c.Query("stime")
	if stimeStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	stimeInt, err := strconv.Atoi(stimeStr)
	if err != nil {
		MakeApiResponseErrorParams(c)
		return
	}

	now := time.Now()
	stime := now.AddDate(0, 0, -stimeInt)

	// 获取近期各类型数据
	results, err := service.GetStatDetailsByType(uid, stime)
	if err != nil {
		service.Logger.Error("GetStatDetailsByType", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if len(results) == 0 {
		results = make([]model.StatDetailsTypeCount, 0)
	}

	data := map[string]interface{}{
		"results": results,
	}

	MakeApiResponseSuccess(c, data)
}

// 通过uid获取用户数据总数Map
func GetUserStatMapByUidHandler(c *gin.Context) {
	uidStr := c.Query("uid")
	if uidStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	uid, err := strconv.Atoi(uidStr)
	if err != nil {
		MakeApiResponseErrorParams(c)
		return
	}

	//根据uid,type获取UserStatMap
	userStatMap, err := service.GetUserStatMapByType(uid)
	if err != nil {
		service.Logger.Error("GetUserStatMapByType err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if len(userStatMap) == 0 {
		userStatMap = make(map[int]int, 0)
	}

	data := map[string]interface{}{
		"userStatMap": userStatMap,
	}

	MakeApiResponseSuccess(c, data)
}

// 获取用户数据总数Map
func GetUserStatMapHandler(c *gin.Context) {
	uid, _ := service.GetUserFromCookie(c)
	if uid == 0 {
		MakeApiResponseError(c, CODE_USER_NOT_LOGIN)
		return
	}

	//根据uid,type获取UserStatMap
	userStatMap, err := service.GetUserStatMapByType(uid)
	if err != nil {
		service.Logger.Error("GetUserStatMapByType err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if len(userStatMap) == 0 {
		userStatMap = make(map[int]int, 0)
	}

	data := map[string]interface{}{
		"userStatMap": userStatMap,
	}

	MakeApiResponseSuccess(c, data)
}
