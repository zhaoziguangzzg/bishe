package controller

import (
	"bishe/internal/app/knowledge_sharing/model"
	"bishe/internal/app/knowledge_sharing/service"
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
		MakeApiResponseError(c, CODE_INFORMATION_CONTENT_LEN_INVASLID)
		return
	}

	uname := c.GetString("uname")
	unameLen := len(uname)
	if unameLen > model.INFORMATION_MAX_RECEIVE_NAME || unameLen == 0 {
		MakeApiResponseError(c, CODE_USER_NAME_LEN_INVALID)
		return
	}

	uid, _ := service.GetUserFromCookie(c)
	if uid == 0 {
		MakeApiResponseError(c, CODE_USER_NOT_LOGIN)
		return
	}

	createTime := time.Now()

	// 构造消息
	newInformation := &model.Information{ //其中包含自动生成的id
		SendId:      uid,
		ReceiveName: uname,
		Content:     content,
		CreateAt:    &createTime,
		UpdateAt:    &createTime,
	}

	// 插入数据库
	err := service.UserAddInformation(newInformation)
	if err != nil {
		service.Logger.Error("CreateInformation err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	// 返回成功响应
	MakeApiResponseSuccessDefault(c)
}

// 用户获取消息
func GetUserInformationHandler(c *gin.Context) {
	//获取uid，uname
	uid, uname := service.GetUserFromCookie(c)
	if uid == 0 || uname == "" {
		MakeApiResponseError(c, CODE_USER_NOT_LOGIN)
		return
	}

	//根据uname获取消息
	information, err := service.GetInformationByUname(uname)
	if err != nil {
		service.Logger.Error("GetInformationByUname", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if information == nil {
		MakeApiResponseError(c, CODE_INFORMATION_NOT_EXIST)
		return
	}

	data := map[string]interface{}{
		"information": information,
	}

	MakeApiResponseSuccess(c, data)
}

// 删除发送的消息
func DeletedInformationByUpdateIsDeletedHandler(c *gin.Context) {
	//更新字段
	iidStr := c.Query("iid")
	if iidStr == "" {
		service.Logger.Error("Getiid err", zap.String("err", "get iid err"))
		MakeApiResponseErrorParams(c)
		return
	}

	iid, err := strconv.Atoi(iidStr)
	if err != nil {
		service.Logger.Error("Atoi iidStr err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
	}

	affectRows, err := service.UpdateInformationIsDeleted(iid)
	if err != nil || affectRows == 0 {
		service.Logger.Error("UpdateInformationIsDeleted err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	MakeApiResponseSuccessDefault(c)
}
