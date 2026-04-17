package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AnnounceEditPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "announce/edit.html", nil)
}
