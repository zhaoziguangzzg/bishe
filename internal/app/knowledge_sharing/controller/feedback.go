package controller

import (
	"bishe/internal/app/knowledge_sharing/model"
	"bishe/internal/app/knowledge_sharing/service"
	"strconv"
	"time"
	"unicode/utf8"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 创建反馈
func AddUserFeedbackHandler(c *gin.Context) {
	uid := service.GetUidFromContext(c)

	content := c.PostForm("content")

	contentLen := utf8.RuneCountInString(content)
	if contentLen == 0 || contentLen > model.FEEDBACK_MAX_CONTENT {
		MakeApiResponseError(c, CODE_FEEDBACK_CONTENT_LEN_INVASLID)
		return
	}

	createTime := time.Now()

	newFeedback := &model.Feedback{ //其中包含自动生成的id
		UserId:         uid,
		Content:        content,
		FeedbackTime:   &createTime,
		CreateAt:       &createTime,
		UpdateAt:       &createTime,
		FeedbackStatus: model.FEEDBACK_STATUS_OPEN,
		IsDeleted:      model.FEEDBACK_NOT_DELETED,
	}

	err := service.CreateUserFeedback(newFeedback)
	if err != nil {
		service.Logger.Error("CreateUserFeedback err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	MakeApiResponseSuccessDefault(c)
}

// 获取全部未处理反馈
func GetAllFeedbackHandler(c *gin.Context) {
	pageStr := c.Query("page")
	page := GetPage(pageStr)

	pagesize := 10

	feedbacks, err := service.GetAllFeedback(page, pagesize)
	if err != nil {
		service.Logger.Error("GetAllFeedback", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if feedbacks == nil {
		feedbacks = make([]model.Feedback, 0)
	}

	data := map[string]interface{}{
		"feedbacks": feedbacks,
	}

	MakeApiResponseSuccess(c, data)
}

// 获取反馈内容
func GetFeedbackContentHandler(c *gin.Context) {
	idStr := c.Query("id")
	if idStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		MakeApiResponseErrorParams(c)
		return
	}

	//根据id获取反馈
	feedback, err := service.GetFeedbackById(id)
	if err != nil {
		service.Logger.Error("", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if feedback == nil {
		MakeApiResponseError(c, CODE_FEEDBACK_NOT_EXIST)
		return
	}

	data := map[string]interface{}{
		"feedback": feedback,
	}

	MakeApiResponseSuccess(c, data)
}

// 更新反馈状态
func UpdateFeedbackStatusHandler(c *gin.Context) {
	idStr := c.PostForm("id")
	if idStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		MakeApiResponseErrorParams(c)
		return
	}

	reply := c.PostForm("reply")
	if reply == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	//根据id获取反馈
	feedback, err := service.GetFeedbackById(id)
	if err != nil {
		service.Logger.Error("", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if feedback == nil {
		MakeApiResponseError(c, CODE_FEEDBACK_NOT_EXIST)
		return
	}

	nowTime := time.Now()

	//保存回复，更新状态
	affectRows, err := service.UpdateFeedbackStatusReplyById(id, reply, nowTime)
	if err != nil || affectRows == 0 {
		service.Logger.Error("UpdateFeedbackStatusReplyById err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	// 给用户发通知
	typei := model.NOTICE_TYPE_FEEDBACK
	//TODO 异步处理
	noticeMsg := &model.NoticeMsg{
		Type: typei,
		Uid:  feedback.UserId,
		Time: nowTime.Unix(),
	}

	_, _, err = service.ProduceKafkaNoticeMessage(noticeMsg)
	if err != nil {
		service.Logger.Error("ProduceKafkaNoticeMessage err", zap.Error(err))
		err = nil
	}

	MakeApiResponseSuccessDefault(c)
}

// 获取当前用户的反馈列表
func GetFeedbackByUidHandler(c *gin.Context) {
	uid := service.GetUidFromContext(c)
	if uid == 0 {
		MakeApiResponseError(c, CODE_USER_NOT_LOGIN)
		return
	}

	pageStr := c.Query("page")
	page := GetPage(pageStr)

	pagesize := 10

	feedbacks, err := service.GetFeedbackByUid(uid, page, pagesize)
	if err != nil {
		service.Logger.Error("GetFeedbackByUid", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if feedbacks == nil {
		feedbacks = make([]model.Feedback, 0)
	}

	data := map[string]interface{}{
		"feedbacks": feedbacks,
	}

	MakeApiResponseSuccess(c, data)
}
