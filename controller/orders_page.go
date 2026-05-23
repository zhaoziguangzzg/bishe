package controller

import (
	"github.com/gin-gonic/gin"
)

func OrdersIndexPageHandler(c *gin.Context) {
	RenderIndexPage(c, "orders/index.html", nil)
}
