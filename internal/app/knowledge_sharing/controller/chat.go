package controller

import (
	"bishe/internal/app/knowledge_sharing/model"
	"bishe/internal/app/knowledge_sharing/service"

	"github.com/gin-gonic/gin"
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

}
