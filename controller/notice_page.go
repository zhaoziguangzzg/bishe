package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// OrdersIndexPageHandler
func NoticeIndexPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "notice/index.html", nil)
}

func NoticeDetailPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "notice/detail.html", nil)
}
