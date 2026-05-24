package controller

import (
	"github.com/gin-gonic/gin"
)

func AnnounceListPageHandler(c *gin.Context) {
	RenderIndexPage(c, "announce/list.html", nil)
}

func AnnounceDetailPageHandler(c *gin.Context) {
	RenderIndexPage(c, "announce/detail.html", nil)
}
