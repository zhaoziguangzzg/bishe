package controller

import (
	"github.com/gin-gonic/gin"
)

func CourseIndexPageHandler(c *gin.Context) {
	RenderIndexPage(c, "course/index.html", nil)
}

func AddCoursePageHandler(c *gin.Context) {
	RenderIndexPage(c, "course/add.html", nil)
}

func CourseDetailPageHandler(c *gin.Context) {
	RenderIndexPage(c, "course/detail.html", nil)
}

func EditCoursePageHandler(c *gin.Context) {
	RenderIndexPage(c, "course/edit.html", nil)
}
