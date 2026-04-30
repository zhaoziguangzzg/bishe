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
	uid := c.GetInt("uid")

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
		Content:          content,
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

// 获取全部未审核举报
func GetAllAccusationEssayHandler(c *gin.Context) {
	pageStr := c.Query("page")
	page := GetPage(pageStr)
	pageSize := 10

	accusations, err := service.GetAllAccusationEssay(page, pageSize)
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

// 获取举报内容文章
func GetEssayContentByAccusationHandler(c *gin.Context) {
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

// 更新accusation状态
func UpdateAccusationStatusHandler(c *gin.Context) {
	aidStr := c.PostForm("accusation_id")
	if aidStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	aid, err := strconv.Atoi(aidStr)
	if err != nil {
		MakeApiResponseErrorParams(c)
		return
	}

	statusStr := c.PostForm("accusation_status")
	if statusStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	status, err := strconv.Atoi(statusStr)
	if err != nil {
		MakeApiResponseErrorParams(c)
		return
	}

	userIdStr := c.PostForm("user_id")
	if userIdStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		MakeApiResponseErrorParams(c)
		return
	}

	authorIdStr := c.PostForm("author_id")
	if authorIdStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	authorId, err := strconv.Atoi(authorIdStr)
	if err != nil {
		MakeApiResponseErrorParams(c)
		return
	}

	if status == model.ACCUSATION_STATUS_NORMAL {
		// 更新举报信息为无违规
		affectRows, err := service.UpdateAccusationNormalByAid(aid)
		if err != nil || affectRows == 0 {
			service.Logger.Error("UpdateAccusationNormalByAid err", zap.Error(err))
			MakeApiResponseErrorDefault(c)
			return
		}

		content := "该举报文章无违规"
		receiveId := userId

		// 插入数据库
		err = service.AddAccusationInformation(content, receiveId)
		if err != nil {
			service.Logger.Error("AddAccusationInformation err", zap.Error(err))
			MakeApiResponseErrorDefault(c)
			return
		}

		MakeApiResponseSuccessDefault(c)

		return
	} else if status == model.ACCUSATION_STATUS_VIOLATE {
		// 更新举报信息为有违规
		affectRows, err := service.UpdateAccusationViolateByAid(aid)
		if err != nil || affectRows == 0 {
			service.Logger.Error("UpdateAccusationViolateByAid err", zap.Error(err))
			MakeApiResponseErrorDefault(c)
			return
		}

		content := "该文章有违规，以被封禁"
		receiveId := authorId

		// 插入数据库
		err = service.AddAccusationInformation(content, receiveId)
		if err != nil {
			service.Logger.Error("AddAccusationInformation err", zap.Error(err))
			MakeApiResponseErrorDefault(c)
			return
		}
		MakeApiResponseSuccessDefault(c)
	}
}
