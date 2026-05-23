package controller

import (
	"github.com/gin-gonic/gin"
)

// 创建圈子页面
func AddCirclePageHandler(c *gin.Context) {
	RenderIndexPage(c, "circle/add.html", nil)
}

// 圈子详情页面
func CircleDetailPageHandler(c *gin.Context) {
	RenderIndexPage(c, "circle/detail.html", nil)
}

// 修改圈子页面
func EditCirclePageHandler(c *gin.Context) {
	RenderIndexPage(c, "circle/edit.html", nil)
}

// 圈子内首页页面
func CircleIndexPageHandler(c *gin.Context) {
	RenderIndexPage(c, "circle/index.html", nil)
}

// 圈子排行页面
func CircleListPageHandler(c *gin.Context) {
	RenderIndexPage(c, "circle/list.html", nil)
}

// 搜索圈子页面
func CircleSearchPageHandler(c *gin.Context) {
	RenderIndexPage(c, "circle/search.html", nil)
}
