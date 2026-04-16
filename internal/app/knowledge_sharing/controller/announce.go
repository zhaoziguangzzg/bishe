package controller

import (
	"bishe/internal/app/knowledge_sharing/model"
	"bishe/internal/app/knowledge_sharing/service"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 添加公告
func AddAnnounceHandler(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")

	titleLen := len(title)
	if titleLen > model.ANNOUNCE_MAX_TITLE || titleLen == 0 {
		MakeApiResponseError(c, CODE_ANNOUNCE_TITLE_LEN_INVASLID)
		return
	}

	contentLen := len(content)
	if contentLen > model.ANNOUNCE_MAX_CONTENT || contentLen == 0 {
		MakeApiResponseError(c, CODE_ANNOUNCE_CONTENT_LEN_INVASLID)
		return
	}

	startTimeStr := c.PostForm("start_time")
	if startTimeStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	startTime, err := time.ParseInLocation("2006-01-02 15:04:05", startTimeStr, time.Local)
	if err != nil {
		MakeApiResponseErrorParams(c)
		return
	}

	endTimeStr := c.PostForm("end_time")
	if endTimeStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	//转为本时区
	endTime, err := time.ParseInLocation("2006-01-02 15:04:05", endTimeStr, time.Local)
	if err != nil {
		MakeApiResponseErrorParams(c)
		return
	}

	createTime := time.Now()

	//开始时间大于当前时间
	// a after b    a再b后
	if createTime.After(startTime) {
		MakeApiResponseErrorParams(c)
		return
	}

	//开始时间小于结束时间
	if startTime.After(endTime) {
		MakeApiResponseErrorParams(c)
		return
	}

	// 构造公告
	announce := &model.Announce{ //其中包含自动生成的id
		Title:     title,
		Content:   content,
		StartTime: &startTime,
		EndTime:   &endTime,
		CreateAt:  &createTime,
		UpdateAt:  &createTime,
		IsDeleted: model.IS_DELETED_NO,
	}

	err = service.CreateAnnounce(announce)
	if err != nil {
		service.Logger.Error("CreateAnnounce err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	MakeApiResponseSuccessDefault(c)
}

// 获取全部显示公告列表
func GetAllAnnounceByTimeHandler(c *gin.Context) {
	cTime := time.Now()

	pageStr := c.Query("page")
	page := GetPage(pageStr)

	pagesize := 10

	//获取全部公告
	announces, err := service.GetAllAnnounceByTime(cTime, page, pagesize)
	if err != nil {
		service.Logger.Error("GetAllAnnounceByTime", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if len(announces) == 0 {
		announces = make([]model.Announce, 0)
	}

	MakeApiResponseSuccess(c, map[string]interface{}{
		"announces": announces,
	})
}

// 获取全部显示公告列表
func GetAllAnnounceHandler(c *gin.Context) {

	pageStr := c.Query("page")
	page := GetPage(pageStr)

	pagesize := 10

	//获取全部公告
	announces, err := service.GetAllAnnounce(page, pagesize)
	if err != nil {
		service.Logger.Error("GetAllAnnounce", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if len(announces) == 0 {
		announces = make([]model.Announce, 0)
	}

	MakeApiResponseSuccess(c, map[string]interface{}{
		"announces": announces,
	})
}

// 获取某公告
func GetAnnounceHandler(c *gin.Context) {
	//获取公告id
	announceIdStr := c.Query("announce_id")
	if announceIdStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	announceId, err := strconv.Atoi(announceIdStr)
	if err != nil {
		MakeApiResponseErrorParams(c)
		return
	}

	//根据id获取公告
	announce, err := service.GetAnnounceById(announceId)
	if err != nil {
		service.Logger.Error("GetAnnounceById", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if announce == nil {
		MakeApiResponseError(c, CODE_ANNOUNCE_NOT_EXIST)
		return
	}

	data := map[string]interface{}{
		"announce": announce,
	}

	MakeApiResponseSuccess(c, data)
}

// 更新公告信息
func UpdateAnnounceHandler(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")

	titleLen := len(title)
	if titleLen > model.ANNOUNCE_MAX_TITLE || titleLen == 0 {
		MakeApiResponseError(c, CODE_ANNOUNCE_TITLE_LEN_INVASLID)
		return
	}

	contentLen := len(content)
	if contentLen > model.ANNOUNCE_MAX_CONTENT || contentLen == 0 {
		MakeApiResponseError(c, CODE_ANNOUNCE_CONTENT_LEN_INVASLID)
		return
	}

	startTimeStr := c.PostForm("start_time")
	if startTimeStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	startTime, err := time.ParseInLocation("2006-01-02 15:04:05", startTimeStr, time.Local)
	if err != nil {
		MakeApiResponseErrorParams(c)
		return
	}

	endTimeStr := c.PostForm("end_time")
	if endTimeStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	//转为本时区
	endTime, err := time.ParseInLocation("2006-01-02 15:04:05", endTimeStr, time.Local)
	if err != nil {
		MakeApiResponseErrorParams(c)
		return
	}

	announceIdStr := c.PostForm("announce_id")
	if announceIdStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	announceId, err := strconv.Atoi(announceIdStr)
	if err != nil {
		MakeApiResponseErrorParams(c)
		return
	}

	//根据id获取公告
	announce, err := service.GetAnnounceById(announceId)
	if err != nil {
		service.Logger.Error("GetAnnounceById", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if announce == nil {
		MakeApiResponseError(c, CODE_ANNOUNCE_NOT_EXIST)
		return
	}

	newAnnounce := map[string]interface{}{
		"title":      title,
		"content":    content,
		"start_time": &startTime,
		"end_time":   &endTime,
	}

	//根据id更新公告
	affectRows, err := service.UpdateAnnounceById(announceId, newAnnounce)
	if err != nil || affectRows == 0 {
		service.Logger.Error("UpdateAnnounceById err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	MakeApiResponseSuccessDefault(c)
}

// 删除公告
func DeletedAnnounceByUpdateIsDeletedHandler(c *gin.Context) {
	//获取公告id
	announceIdStr := c.PostForm("announce_id")
	if announceIdStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	announceId, err := strconv.Atoi(announceIdStr)
	if err != nil {
		MakeApiResponseErrorParams(c)
		return
	}

	//根据id删除公告
	affectRows, err := service.UpdateAnnounceIsDeleted(announceId)
	if err != nil || affectRows == 0 {
		service.Logger.Error("UpdateAnnounceIsDeleted err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	MakeApiResponseSuccessDefault(c)
}
