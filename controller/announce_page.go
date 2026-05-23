package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AnnounceListPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "announce/list.html", nil)
}

func AnnounceDetailPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "announce/detail.html", nil)
}
