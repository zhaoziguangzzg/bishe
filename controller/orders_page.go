package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func OrdersIndexPageHandler(c *gin.Context) {
	RenderIndexPage(c, "orders/index.html", nil)
}

func OrdersPayPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "orders/pay.html", nil)
}
