package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ChatIndexPageHandler 获取私信首页
func ChatIndexPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "chat/index.html", nil)
}

// ChatDetailPageHandler 获取私信详情页面
func ChatDetailPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "chat/detail.html", nil)
}
