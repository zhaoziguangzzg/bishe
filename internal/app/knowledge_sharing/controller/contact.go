package controller

import (
	"bishe/internal/app/knowledge_sharing/model"
	"bishe/internal/app/knowledge_sharing/service"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 添加联系人
func AddUserContactHandler(c *gin.Context) {
	uid, _ := service.GetUserFromCookie(c)
	if uid == 0 {
		MakeApiResponseError(c, CODE_USER_NOT_LOGIN)
		return
	}

	// model.Contact

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

	//查询用户的联系人
	contact, err := service.GetUserContact(uid, receiveId)
	if err != nil {
		service.Logger.Error("GetUserContact", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	//仅 不存在，存在状态为删除 两种
	if contact == nil {
		createTime := time.Now()

		newContact := &model.Contact{ //其中包含自动生成的id
			SendId:        uid,
			ReceiveId:     receiveId,
			CreateAt:      &createTime,
			UpdateAt:      &createTime,
			ContactStatus: model.CONTACT_STATUS_NORMAL,
		}

		err = service.CreateUserContact(newContact)
		if err != nil {
			service.Logger.Error("CreateUserContact err", zap.Error(err))
			MakeApiResponseErrorDefault(c)
			return
		}

		MakeApiResponseSuccessDefault(c)
		return
	}
}

// 删除联系人
func DeleteUserContactHandler(c *gin.Context) {
	uid, _ := service.GetUserFromCookie(c)
	if uid == 0 {
		MakeApiResponseError(c, CODE_USER_NOT_LOGIN)
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

	//更新删除字段来删除联系人
	affectRows, err := service.DeleteUserContactByReceiveId(uid, receiveId)
	if err != nil || affectRows == 0 {
		service.Logger.Error("DeleteUserContactByReceiveId err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	MakeApiResponseSuccessDefault(c)
}

// 获取全部联系人
func GetUserAllContactHandler(c *gin.Context) {
	uid, _ := service.GetUserFromCookie(c)
	if uid == 0 {
		MakeApiResponseError(c, CODE_USER_NOT_LOGIN)
		return
	}

	page := c.GetInt("page")
	if page < 1 {
		page = 1
	}

	pagesize := 10

	contacts, err := service.GetUserAllContact(uid, page, pagesize)
	if err != nil {
		service.Logger.Error("GetUserAllContact err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if contacts == nil {
		contacts = make([]model.Contact, 0)
	}

	data := map[string]interface{}{
		"contacts": contacts,
	}

	MakeApiResponseSuccess(c, data)

}

// 获取联系人
func GetUserContactHandler(c *gin.Context) {
	idStr := c.Query("id")
	if idStr != "" {
		MakeApiResponseErrorParams(c)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		MakeApiResponseErrorParams(c)
		return
	}

	//根据id获取联系人
	contact, err := service.GetUserContactById(id)
	if err != nil {
		service.Logger.Error("GetUserContactById err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if contact != nil {
		MakeApiResponseError(c, CODE_CONTACT_NOT_EXIST)
		return
	}

	data := map[string]interface{}{
		"contact": contact,
	}

	MakeApiResponseSuccess(c, data)
}
