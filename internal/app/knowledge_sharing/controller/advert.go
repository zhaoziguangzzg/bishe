package controller

import (
	"bishe/internal/app/knowledge_sharing/model"
	"bishe/internal/app/knowledge_sharing/service"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 添加广告
func AddAdvertHandler(c *gin.Context) {
	position := c.PostForm("position")
	content := c.PostForm("content")
	advertAddr := c.PostForm("addvert_addr")

	positionLen := len(position)
	if positionLen > model.ADVERT_MAX_POSITION || positionLen == 0 {
		MakeApiResponseError(c, CODE_ADVERT_POSITION_LEN_INVASLID)
		return
	}

	contentLen := len(content)
	if contentLen > model.ADVERT_MAX_CONTENT || contentLen == 0 {
		MakeApiResponseError(c, CODE_ADVERT_CONTENT_LEN_INVASLID)
		return
	}

	advertAddrLen := len(advertAddr)
	if advertAddrLen > model.ADVERT_MAX_ADDR || advertAddrLen == 0 {
		MakeApiResponseError(c, CODE_ADVERT_ADDR_LEN_INVASLID)
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

	// 构造广告
	advert := &model.Advert{ //其中包含自动生成的id
		Position:   position,
		AdvertAddr: advertAddr,
		Content:    content,
		StartTime:  &startTime,
		EndTime:    &endTime,
		CreateAt:   &createTime,
		UpdateAt:   &createTime,
		IsDeleted:  model.IS_DELETED_NO,
	}

	fileType := service.FILE_TYPE_ADVERT_IMG

	// 处理广告图片上传
	avatarPath := ""
	file, header, err := c.Request.FormFile("avatar")
	//判断错误不等于无文件
	if err != nil && err != http.ErrMissingFile {
		service.Logger.Error("FormFile err", zap.Error(err))
		MakeApiResponseErrorParams(c)
		return
	}

	//判断size不是空
	if err == nil && header.Size != 0 {
		avatarPath, err = service.FileSave(file, header, fileType, createTime)
		if err != nil {
			MakeApiResponseErrorDefault(c)
			return
		}
		advert.Img = avatarPath
	}

	err = service.CreateAdvert(advert)
	if err != nil {
		service.Logger.Error("CreateAdvert err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	MakeApiResponseSuccessDefault(c)
}

// 获取全部广告列表
func GetAllAdvertByTimeHandler(c *gin.Context) {
	position := c.Query("position")

	positionLen := len(position)
	if positionLen > model.ADVERT_MAX_POSITION || positionLen == 0 {
		MakeApiResponseError(c, CODE_ADVERT_POSITION_LEN_INVASLID)
		return
	}
	cTime := time.Now()

	pageStr := c.Query("page")
	page := GetPage(pageStr)

	pagesize := 10

	//获取全部广告
	adverts, err := service.GetAllAdvertByTime(cTime, position, page, pagesize)
	if err != nil {
		service.Logger.Error("GetAllAdvertByTime", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if len(adverts) == 0 {
		adverts = make([]model.Advert, 0)
	}

	MakeApiResponseSuccess(c, map[string]interface{}{
		"adverts": adverts,
	})
}

// 获取全部广告列表
func GetAllAdvertHandler(c *gin.Context) {

	pageStr := c.Query("page")
	page := GetPage(pageStr)

	pagesize := 10

	//获取全部广告
	adverts, err := service.GetAllAdvert(page, pagesize)
	if err != nil {
		service.Logger.Error("GetAllAdvert", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if len(adverts) == 0 {
		adverts = make([]model.Advert, 0)
	}

	MakeApiResponseSuccess(c, map[string]interface{}{
		"adverts": adverts,
	})
}

// 获取某广告
func GetAdvertHandler(c *gin.Context) {
	//获取广告id
	advertIdStr := c.Query("advert_id")
	if advertIdStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	advertId, err := strconv.Atoi(advertIdStr)
	if err != nil {
		MakeApiResponseErrorParams(c)
		return
	}

	//根据id获取广告
	advert, err := service.GetAdvertById(advertId)
	if err != nil {
		service.Logger.Error("GetAdvertById", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if advert == nil {
		MakeApiResponseError(c, CODE_ANNOUNCE_NOT_EXIST)
		return
	}

	data := map[string]interface{}{
		"advert": advert,
	}

	MakeApiResponseSuccess(c, data)
}

// 更新广告信息
func UpdateAdvertHandler(c *gin.Context) {
	position := c.PostForm("position")
	content := c.PostForm("content")
	advertAddr := c.PostForm("addvert_addr")

	positionLen := len(position)
	if positionLen > model.ADVERT_MAX_POSITION || positionLen == 0 {
		MakeApiResponseError(c, CODE_ADVERT_POSITION_LEN_INVASLID)
		return
	}

	contentLen := len(content)
	if contentLen > model.ADVERT_MAX_CONTENT || contentLen == 0 {
		MakeApiResponseError(c, CODE_ADVERT_CONTENT_LEN_INVASLID)
		return
	}

	advertAddrLen := len(advertAddr)
	if advertAddrLen > model.ADVERT_MAX_ADDR || advertAddrLen == 0 {
		MakeApiResponseError(c, CODE_ADVERT_ADDR_LEN_INVASLID)
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

	if startTime.After(endTime) {
		MakeApiResponseErrorParams(c)
		return
	}

	advertIdStr := c.PostForm("advert_id")
	if advertIdStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	advertId, err := strconv.Atoi(advertIdStr)
	if err != nil {
		MakeApiResponseErrorParams(c)
		return
	}

	//根据id获取广告
	advert, err := service.GetAdvertById(advertId)
	if err != nil {
		service.Logger.Error("GetAdvertById", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if advert == nil {
		MakeApiResponseError(c, CODE_ANNOUNCE_NOT_EXIST)
		return
	}

	newAdvert := map[string]interface{}{
		"position":    position,
		"advert_addr": advertAddr,
		"content":     content,
		"start_time":  &startTime,
		"end_time":    &endTime,
	}

	fileType := service.FILE_TYPE_ADVERT_IMG
	nowTime := time.Now()

	// 处理广告图片上传
	avatarPath := ""
	file, header, err := c.Request.FormFile("avatar")
	//判断错误不等于无文件
	if err != nil && err != http.ErrMissingFile {
		service.Logger.Error("FormFile err", zap.Error(err))
		MakeApiResponseErrorParams(c)
		return
	}

	//判断size不是空
	if err == nil && header.Size != 0 {
		avatarPath, err = service.FileSave(file, header, fileType, nowTime)
		if err != nil {
			MakeApiResponseErrorDefault(c)
			return
		}
		newAdvert["img"] = avatarPath
	}

	//根据id更新广告
	affectRows, err := service.UpdateAdvertById(advertId, newAdvert)
	if err != nil || affectRows == 0 {
		service.Logger.Error("UpdateAdvertById err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	MakeApiResponseSuccessDefault(c)
}

// 删除广告
func DeletedAdvertByUpdateIsDeletedHandler(c *gin.Context) {
	//获取广告id
	advertIdStr := c.PostForm("advert_id")
	if advertIdStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	advertId, err := strconv.Atoi(advertIdStr)
	if err != nil {
		MakeApiResponseErrorParams(c)
		return
	}

	//根据id删除广告
	affectRows, err := service.UpdateAdvertIsDeleted(advertId)
	if err != nil || affectRows == 0 {
		service.Logger.Error("UpdateAdvertIsDeleted err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	MakeApiResponseSuccessDefault(c)
}
