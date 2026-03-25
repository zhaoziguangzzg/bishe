package controller

import (
	"bishe/internal/app/knowledge_sharing/model"
	"bishe/internal/app/knowledge_sharing/service"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 创建举报
func AddUserAccusationEssayHandler(c *gin.Context) {
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

	content := c.PostForm("content")

	contentLen := len(content)
	if contentLen == 0 || contentLen > model.ACCUSATION_MAX_CONTENT {
		MakeApiResponseError(c, CODE_ACCUSATION_CONTENT_LEN_INVASLID)
		return
	}

	//查询用户的举报
	accusation, err := service.GetUserAccusationEssay(uid, eid)
	if err != nil {
		service.Logger.Error("GetUserAccusationEssay", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if accusation != nil {
		MakeApiResponseSuccess(c, CODE_ACCUSATION_EXIST)
		return
	}

	createTime := time.Now()

	newAccusation := &model.Accusation{ //其中包含自动生成的id
		UserId:           uid,
		EssayId:          eid,
		AccusationTime:   &createTime,
		CreateAt:         &createTime,
		UpdateAt:         &createTime,
		AccusationStatus: model.ACCUSATION_STATUS_WAIT,
		IsDeleted:        model.ACCUSATION_NOT_DELETED,
	}

	err = service.CreateUserAccusation(newAccusation)
	if err != nil {
		service.Logger.Error("CreateUserAccusation err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	MakeApiResponseSuccessDefault(c)
}

func GetAllAccusationEssayHandler(c *gin.Context) {
	page := c.GetInt("page")
	if page < 1 {
		page = 1
	}

	pagesize := 10

	accusations, err := service.GetAllAccusationEssay(page, pagesize)
	if err != nil {
		service.Logger.Error("GetAllAccusationEssay", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if accusations == nil {
		accusations = make([]model.Accusation, 0)
	}

	data := map[string]interface{}{
		"accusations": accusations,
	}

	MakeApiResponseSuccess(c, data)
}

func GetEssayContentByAccusation(c *gin.Context) {
	aidStr := c.Query("accusation_id")
	if aidStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	aid, err := strconv.Atoi(aidStr)
	if err != nil {
		MakeApiResponseErrorParams(c)
		return
	}

	//根据aid获取 举报
	accusation, err := service.GetAccusationByAid(aid)
	if err != nil {
		service.Logger.Error("", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if accusation == nil {
		MakeApiResponseError(c, CODE_ACCUSATION_NOT_EXIST)
		return
	}

	eid := accusation.EssayId

	//根据eid获取essay
	essay, err := service.GetEssayByEid(eid)
	if err != nil {
		service.Logger.Error("GetEssayByEid err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if essay == nil {
		MakeApiResponseError(c, CODE_ESSAY_NOT_EXIST)
		return
	}

	data := map[string]interface{}{
		"accusation": accusation,
		"essay":      essay,
	}

	MakeApiResponseSuccess(c, data)
}
