package controller

import (
	"github.com/gin-gonic/gin"
)

// ChatIndexPageHandler 获取私信首页
func ChatIndexPageHandler(c *gin.Context) {
	RenderIndexPage(c, "chat/index.html", nil)
}
