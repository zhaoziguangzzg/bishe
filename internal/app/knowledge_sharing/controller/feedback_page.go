package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func FeedbackIndexPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "feedback/index.html", nil)
}

func FeedbackDetailPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "feedback/detail.html", nil)
}

func FeedbackEditPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "feedback/edit.html", nil)
}
