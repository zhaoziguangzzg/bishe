package controller

import (
	"bishe/internal/app/knowledge_sharing/model"
	"bishe/internal/app/knowledge_sharing/service"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 获取联系人列表
func GetChatContactListHandler(c *gin.Context) {
	uid, _ := service.GetUserFromCookie(c)
	if uid == 0 {
		MakeApiResponseError(c, CODE_USER_NOT_LOGIN)
		return
	}

	pageStr := c.Query("page")
	page := GetPage(pageStr)
	pageSize := 10

	chatContacts, err := service.GetChatContactList(uid, page, pageSize)
	if err != nil {
		service.Logger.Error("GetChatContactList", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if len(chatContacts) == 0 {
		chatContacts = make([]model.ChatContact, 0)
		data := map[string]interface{}{
			"chatContacts": chatContacts,
		}

		MakeApiResponseSuccess(c, data)
		return
	}

	var uids []int
	for _, v := range chatContacts {
		if v.SendUid == uid {
			uids = append(uids, v.ReceiveUid)
		} else {
			uids = append(uids, v.SendUid)
		}
	}

	//根据uids获取userMap
	userMap, err := service.GetUsersByUidMap(uids)
	if err != nil {
		service.Logger.Error("GetUsersByUids", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if len(userMap) == 0 {
		service.Logger.Error("GetUsersByUidMap len(userMap) == 0")
		MakeApiResponseErrorDefault(c)
		return
	}

	userChatContacts := make([]model.UserChatContact, 0)

	// for _, v := range uids {
	// 	var userChatContact model.UserChatContact
	// 	user, ok := userMap[v]
	// 	if ok != true {
	// 		service.Logger.Error("set uids err")
	// 		MakeApiResponseErrorDefault(c)
	// 		return
	// 	}
	// 	userChatContact.Name = user.Name
	// 	userChatContacts = append(userChatContacts, userChatContact)
	// }

	for _, v := range chatContacts {
		var vUid int
		if v.SendUid == uid {
			vUid = v.ReceiveUid
		} else {
			vUid = v.SendUid
		}

		vUser, ok := userMap[vUid]
		if !ok {
			service.Logger.Error("set uids err")
			MakeApiResponseErrorDefault(c)
			return
		}

		var userChatContact model.UserChatContact

		updateAt := v.UpdateAt.Format("2006-01-02 15:04:05")
		userChatContact.Uid = vUid
		userChatContact.Name = vUser.Name
		userChatContact.Content = v.Content
		userChatContact.UpdateAt = updateAt
		userChatContacts = append(userChatContacts, userChatContact)
	}

	data := map[string]interface{}{
		"userChatContacts": userChatContacts,
	}

	MakeApiResponseSuccess(c, data)

}

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

	pageStr := c.Query("page")
	page := GetPage(pageStr)

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
