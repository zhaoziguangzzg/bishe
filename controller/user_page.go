package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "user/login.html", nil)
}

func RegisterPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "user/register.html", nil)
}

func IndexPageHandler(c *gin.Context) {
	RenderIndexPage(c, "index/index.html", nil)
}

func ProfilePageHandler(c *gin.Context) {
	RenderIndexPage(c, "user/profile.html", nil)
}

func EditPageHandler(c *gin.Context) {
	RenderIndexPage(c, "user/edit.html", nil)
}

func EditPasswordPageHandler(c *gin.Context) {
	RenderIndexPage(c, "user/edit-password.html", nil)
}
