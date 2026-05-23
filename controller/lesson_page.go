package controller

import (
	"github.com/gin-gonic/gin"
)

// AddLessonPageHandler 创建课时页面
func AddLessonPageHandler(c *gin.Context) {
	RenderIndexPage(c, "lesson/add.html", nil)
}

// LessonDetailPageHandler 课时详情页面
func LessonDetailPageHandler(c *gin.Context) {
	RenderIndexPage(c, "lesson/detail.html", nil)
}

// EditLessonPageHandler 课时修改页面
func EditLessonPageHandler(c *gin.Context) {
	RenderIndexPage(c, "lesson/edit.html", nil)
}
