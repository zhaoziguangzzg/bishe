package controller

import (
	"github.com/gin-gonic/gin"
)

func NoticeIndexPageHandler(c *gin.Context) {
	RenderIndexPage(c, "notice/index.html", nil)
}

func NoticeDetailPageHandler(c *gin.Context) {
	RenderIndexPage(c, "notice/detail.html", nil)
}
