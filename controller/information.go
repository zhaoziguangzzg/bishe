package controller

import (
	"bishe/model"
	"bishe/service"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 消息通知
func CreateInformationHandle(c *gin.Context) {
	// 从表单中获取用户信息
	content := c.PostForm("content")

	contentLen := len(content)
	if contentLen > model.INFORMATION_MAX_CONTENT || contentLen == 0 {
		MakeApiResponseError(c, CODE_CHAT_CONTENT_LEN_INVASLID)
		return
	}

	receiveIdStr := c.PostForm("receive_id")
	if receiveIdStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	receiveId, err := strconv.Atoi(receiveIdStr)
	if err != nil {
		MakeApiResponseErrorParams(c)
		return
	}

	uid := service.GetUidFromContext(c)

	createTime := time.Now()

	// 构造消息
	newInformation := &model.Information{ //其中包含自动生成的id
		SendId:    uid,
		ReceiveId: receiveId,
		Content:   content,
		CreateAt:  &createTime,
		UpdateAt:  &createTime,
	}

	// 插入数据库
	err = service.UserAddInformation(newInformation)
	if err != nil {
		service.Logger.Error("UserAddInformation err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	// 返回成功响应
	MakeApiResponseSuccessDefault(c)
}

// 获取消息联系人列表
func GetInformationUsersHandler(c *gin.Context) {
	uid := service.GetUidFromContext(c)

	pageStr := c.Query("page")
	page := GetPage(pageStr)

	pageSize := 10

	//获取消息用户列表
	informations, err := service.GetUserInformation(uid, page, pageSize)
	if err != nil {
		service.Logger.Error("GetUserInformation", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if informations == nil {
		informations = make([]model.Information, 0)
	}

	data := map[string]interface{}{
		"informations": informations,
	}

	MakeApiResponseSuccess(c, data)

}

// 获取用户接收消息
func GetUserReceiveInformationHandler(c *gin.Context) {
	uid := service.GetUidFromContext(c)

	sendIdStr := c.Query("send_id")
	if sendIdStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	sendId, err := strconv.Atoi(sendIdStr)
	if err != nil {
		MakeApiResponseErrorDefault(c)
		return
	}

	pageStr := c.Query("page")
	page := GetPage(pageStr)

	pageSize := 10

	//根据uid获取接收的消息
	informations, err := service.GetReceiveInformationByUid(uid, sendId, page, pageSize)
	if err != nil {
		service.Logger.Error("GetReceiveInformationByUid", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if informations == nil {
		informations = make([]model.Information, 0)
	}

	data := map[string]interface{}{
		"informations": informations,
	}

	MakeApiResponseSuccess(c, data)
}

// 获取与某人消息
func GetUserSendInformationHandler(c *gin.Context) {
	uid := service.GetUidFromContext(c)

	receiveIdStr := c.Query("receive_id")
	if receiveIdStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	receiveId, err := strconv.Atoi(receiveIdStr)
	if err != nil {
		MakeApiResponseErrorDefault(c)
		return
	}

	pageStr := c.Query("page")
	page := GetPage(pageStr)

	pageSize := 10

	//根据uid获取发送的消息
	informations, err := service.GetSendInformationByUid(uid, receiveId, page, pageSize)
	if err != nil {
		service.Logger.Error("GetSendInformationByUid", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if informations == nil {
		informations = make([]model.Information, 0)
		return
	}

	data := map[string]interface{}{
		"informations": informations,
	}

	MakeApiResponseSuccess(c, data)
}

// 获取用户的消息记录
func GetUserAllInformationHandler(c *gin.Context) {
	uid := service.GetUidFromContext(c)

	pageStr := c.Query("page")
	page := GetPage(pageStr)

	pageSize := 10

	informations, err := service.GetUserAllInformation(uid, page, pageSize)
	if err != nil {
		service.Logger.Error("GetUserAllInformation err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if informations == nil {
		MakeApiResponseError(c, CODE_INFORMATION_NOT_EXIST)
		return
	}

}

// 删除发送的消息
func DeletedInformationByUpdateIsDeletedHandler(c *gin.Context) {
	//更新字段
	iidStr := c.PostForm("iid")
	if iidStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	iid, err := strconv.Atoi(iidStr)
	if err != nil {
		MakeApiResponseErrorDefault(c)
		return
	}

	affectRows, err := service.UpdateInformationIsDeleted(iid)
	if err != nil || affectRows == 0 {
		service.Logger.Error("UpdateInformationIsDeleted err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	MakeApiResponseSuccessDefault(c)
}
