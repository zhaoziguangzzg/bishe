package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminIndexPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/index.html", nil)
}

func AdminEditPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/edit.html", nil)
}
