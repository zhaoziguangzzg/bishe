package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func StatIndexPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "stat/index.html", nil)
}
