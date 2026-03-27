package controller

import (
	"bishe/internal/app/knowledge_sharing/model"
	"bishe/internal/app/knowledge_sharing/service"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 创建反馈
func AddUserFeedbackHandler(c *gin.Context) {
	uid, _ := service.GetUserFromCookie(c)
	if uid == 0 {
		MakeApiResponseError(c, CODE_USER_NOT_LOGIN)
		return
	}

	content := c.PostForm("content")

	contentLen := len(content)
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
		FeedbackStatus: model.FEEDBACK_STATUS_WAIT,
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
	page := c.GetInt("page")
	if page < 1 {
		page = 1
	}

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

	statusStr := c.PostForm("feedback_status")
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

	receiveId := userId
	var content string

	if status == model.FEEDBACK_STATUS_NORMAL {
		// 更新反馈状态为无问题
		affectRows, err := service.UpdateFeedbackNormalById(id)
		if err != nil || affectRows == 0 {
			service.Logger.Error("UpdateFeedbackNormalById err", zap.Error(err))
			MakeApiResponseErrorDefault(c)
			return
		}

		content = "该反馈未查询到问题，期待您的再次反馈"

	} else if status == model.FEEDBACK_STATUS_QUESTIONABLE {
		// 更新反馈状态为有问题
		affectRows, err := service.UpdateFeedbackViolateById(id)
		if err != nil || affectRows == 0 {
			service.Logger.Error("UpdateFeedbackViolateById err", zap.Error(err))
			MakeApiResponseErrorDefault(c)
			return
		}

		content = "该反馈已核实，感谢您的反馈"

	}

	// 插入数据库
	err = service.AddFeedbackInformation(content, receiveId)
	if err != nil {
		service.Logger.Error("AddFeedbackInformation err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	MakeApiResponseSuccessDefault(c)

}
