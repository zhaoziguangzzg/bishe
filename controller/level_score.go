package controller

import (
	"bishe/model"
	"bishe/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 获取用户在圈子全部等级详情
func GetUserCircleLevelAllRecordHandler(c *gin.Context) {
	uid := service.GetUidFromContext(c)

	cidStr := c.Query("cid")
	if cidStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	cid, err := strconv.Atoi(cidStr)
	if err != nil {
		MakeApiResponseErrorParams(c)
		return
	}

	// 获取用户在圈子全部等级详情
	levelScoreRecords, err := service.GetUserOfCircleLevelRecordList(uid, cid)
	if err != nil {
		service.Logger.Error("GetUserStatList", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if len(levelScoreRecords) == 0 {
		levelScoreRecords = make([]model.LevelScoreRecord, 0)
	}

	data := map[string]interface{}{
		"levelScoreRecords": levelScoreRecords,
	}

	MakeApiResponseSuccess(c, data)
}

// 获取用户在圈子的等级分数
func GetUserLevelScoreHandler(c *gin.Context) {
	uid := service.GetUidFromContext(c)

	cidStr := c.Query("cid")
	if cidStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	cid, err := strconv.Atoi(cidStr)
	if err != nil {
		MakeApiResponseErrorParams(c)
		return
	}

	levelScore, err := service.GetUserLevelScoreByUidCid(uid, cid)
	if err != nil {
		service.Logger.Error("GetUserLevelScoreByUidCid", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	score := 0
	if levelScore != nil {
		score = levelScore.Score
	}
	level := score / 1000

	data := map[string]interface{}{
		"score": score,
		"level": level,
	}

	MakeApiResponseSuccess(c, data)
}
