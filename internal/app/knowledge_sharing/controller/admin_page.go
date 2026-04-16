package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminIndexPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/index.html", nil)
}

func AdminEditPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/edit.html", nil)
}

func AccusationEditPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "accusation/edit.html", nil)
}

func FeedbackEditPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "feedback/edit.html", nil)
}

func AnnounceEditPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "announce/edit.html", nil)
}

func AdvertEditPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "advert/edit.html", nil)
}