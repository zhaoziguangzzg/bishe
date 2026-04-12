package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 创建文章页面
func AddEssayPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "essay/add.html", nil)
}
