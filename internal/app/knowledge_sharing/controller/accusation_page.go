package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AccusationEditPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "accusation/edit.html", nil)
}
