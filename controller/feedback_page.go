package controller

import (
	"github.com/gin-gonic/gin"
)

func FeedbackIndexPageHandler(c *gin.Context) {
	RenderIndexPage(c, "feedback/index.html", nil)
}

func FeedbackDetailPageHandler(c *gin.Context) {
	RenderIndexPage(c, "feedback/detail.html", nil)
}
