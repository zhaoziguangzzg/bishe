package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdvertEditPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "advert/edit.html", nil)
}
