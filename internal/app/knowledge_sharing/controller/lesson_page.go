package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AddLessonPageHandler 创建课时页面
func AddLessonPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "lesson/add.html", nil)
}

// LessonDetailPageHandler 课时详情页面
func LessonDetailPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "lesson/detail.html", nil)
}

// EditLessonPageHandler 课时修改页面
func EditLessonPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "lesson/edit.html", nil)
}

//TODO 图片富文本编辑器
