package controller

import (
	"github.com/gin-gonic/gin"
)

func StatIndexPageHandler(c *gin.Context) {
	RenderIndexPage(c, "stat/index.html", nil)
}
