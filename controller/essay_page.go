package controller

import (
	"github.com/gin-gonic/gin"
)

// 创建文章页面
func AddEssayPageHandler(c *gin.Context) {
	RenderIndexPage(c, "essay/add.html", nil)
}

// 文章详情页面
func EssayDetailPageHandler(c *gin.Context) {
	RenderIndexPage(c, "essay/detail.html", nil)
}

// 修改文章页面
func EditEssayPageHandler(c *gin.Context) {
	RenderIndexPage(c, "essay/edit.html", nil)
}

// 搜索文章页面
func SearchEssayPageHandler(c *gin.Context) {
	RenderIndexPage(c, "essay/search.html", nil)
}
