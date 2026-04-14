package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 创建圈子页面
func AddCirclePageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "circle/add.html", nil)
}

// 圈子详情页面
func CircleDetailPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "circle/detail.html", nil)
}

// 修改圈子页面
func EditCirclePageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "circle/edit.html", nil)
}

// 圈子内首页页面
func CircleIndexPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "circle/index.html", nil)
}
