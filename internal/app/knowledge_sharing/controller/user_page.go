package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func LoginPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "user/login.html", nil)
}

func RegisterPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "user/register.html", nil)
}

func ProfilePageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "user/profile.html", nil)
}

func EditPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "user/edit.html", nil)
}

func EditPasswordPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "user/edit-password.html", nil)
}
