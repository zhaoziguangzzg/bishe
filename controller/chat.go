package controller

import (
	"bishe/model"
	"bishe/service"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 发私信消息
func AddChatHandler(c *gin.Context) {
	uid := service.GetUidFromContext(c)

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

	// 返回成功响应，包含创建的私信记录
	MakeApiResponseSuccess(c, map[string]interface{}{
		"chat": chat,
	})
}

// 获取私信记录
func GetChatListHandler(c *gin.Context) {
	uid := service.GetUidFromContext(c)

	chatUidStr := c.Query("chat_uid")
	if chatUidStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	chatUid, err := strconv.Atoi(chatUidStr)
	if err != nil {
		MakeApiResponseErrorParams(c)
		return
	}

	mode := c.DefaultQuery("mode", "latest")
	limit := 10
	if l := c.Query("limit"); l != "" {
		if parsed, err := strconv.Atoi(l); err == nil && parsed > 0 && parsed <= 50 {
			limit = parsed
		}
	}

	var chats []model.Chat

	switch mode {
	case "before":
		baseIdStr := c.Query("base_id")
		if baseIdStr == "" {
			MakeApiResponseErrorParams(c)
			return
		}
		baseId, err := strconv.Atoi(baseIdStr)
		if err != nil {
			MakeApiResponseErrorParams(c)
			return
		}
		chats, err = service.GetChatListBefore(uid, chatUid, baseId, limit)
	case "after":
		baseIdStr := c.Query("base_id")
		if baseIdStr == "" {
			MakeApiResponseErrorParams(c)
			return
		}
		baseId, err := strconv.Atoi(baseIdStr)
		if err != nil {
			MakeApiResponseErrorParams(c)
			return
		}
		chats, err = service.GetChatListAfter(uid, chatUid, baseId, limit)
	default:
		chats, err = service.GetChatListLatest(uid, chatUid, limit)
	}

	if err != nil {
		service.Logger.Error("GetChatList", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if chats == nil {
		chats = make([]model.Chat, 0)
	}

	data := map[string]interface{}{
		"chats": chats,
		"uid":   uid,
	}

	MakeApiResponseSuccess(c, data)
}
