package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AnnounceEditPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "announce/edit.html", nil)
}

func AnnounceDetailPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "announce/detail.html", nil)
}

func AnnounceListPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "announce/list.html", nil)
}
