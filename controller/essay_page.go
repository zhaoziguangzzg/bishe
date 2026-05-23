package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 创建文章页面
func AddEssayPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "essay/add.html", nil)
}

// 文章详情页面
func EssayDetailPageHandler(c *gin.Context) {
	RenderIndexPage(c, "essay/detail.html", nil)
}

// 修改文章页面
func EditEssayPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "essay/edit.html", nil)
}

// 搜索文章页面
func SearchEssayPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "essay/search.html", nil)
}
