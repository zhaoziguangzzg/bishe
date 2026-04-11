package controller

import (
	"github.com/gin-gonic/gin"
)

func IndexPageHandler(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func LoginPageHandler(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}

func RegisterPageHandler(c *gin.Context) {
	c.HTML(200, "register.html", nil)
}

func ProfilePageHandler(c *gin.Context) {
	c.HTML(200, "profile.html", nil)
}

func EditPageHandler(c *gin.Context) {
	c.HTML(200, "edit.html", nil)
}
