package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func FeedbackIndexPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "feedback/index.html", nil)
}

func FeedbackDetailPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "feedback/detail.html", nil)
}
