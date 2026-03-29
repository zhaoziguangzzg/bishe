package controller

import (
	"bishe/internal/app/knowledge_sharing/model"
	"bishe/internal/app/knowledge_sharing/service"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 发私信消息
func AddChatHandler(c *gin.Context) {
	uid, _ := service.GetUserFromCookie(c)
	if uid == 0 {
		MakeApiResponseError(c, CODE_USER_NOT_LOGIN)
		return
	}

	content := c.PostForm("content")

	contentLen := len(content)
	if contentLen > model.CHAT_CONTACT_MAX_CONTENT || contentLen == 0 {
		MakeApiResponseError(c, CODE_CHAT_CONTENT_LEN_INVASLID)
		return
	}

	receiveUidStr := c.PostForm("receive_uid")
	if receiveUidStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	receiveUid, err := strconv.Atoi(receiveUidStr)
	if err != nil {
		MakeApiResponseErrorParams(c)
		return
	}

	createTime := time.Now()

	chat := &model.Chat{ //其中包含自动生成的id
		SendUid:    uid,
		ReceiveUid: receiveUid,
		Content:    content,
		CreateAt:   &createTime,
		UpdateAt:   &createTime,
		IsDeleted:  model.IS_DELETED_NO,
	}

	// 添加私信
	err = service.ChatAdd(chat)
	if err != nil {
		service.Logger.Error("ChatAdd err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	chatContact := &model.ChatContact{
		SendUid:    uid,
		ReceiveUid: receiveUid,
		Content:    content,
		CreateAt:   &createTime,
		UpdateAt:   &createTime,
		IsDeleted:  model.IS_DELETED_NO,
	}

	//添加联系人
	err = service.ChatContactInsertUpdate(chatContact)
	if err != nil {
		service.Logger.Error("ChatContactInsertUpdate err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	// 返回成功响应
	MakeApiResponseSuccessDefault(c)
}
