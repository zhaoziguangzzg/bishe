package controller

import (
	"github.com/gin-gonic/gin"
)

// AddLessonPageHandler 创建课时页面
func AddLessonPageHandler(c *gin.Context) {
	RenderIndexPage(c, "lesson/add.html", nil)
}

// LessonIndexPageHandler 课时列表页面
func LessonIndexPageHandler(c *gin.Context) {
	RenderIndexPage(c, "lesson/index.html", nil)
}

// EditLessonPageHandler 课时修改页面
func EditLessonPageHandler(c *gin.Context) {
	RenderIndexPage(c, "lesson/edit.html", nil)
}
