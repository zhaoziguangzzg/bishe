package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CourseIndexPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "course/index.html", nil)
}

func CourseDetailPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "course/detail.html", nil)
}

func EditCoursePageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "course/edit.html", nil)
}

func AddCoursePageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "course/add.html", nil)
}
