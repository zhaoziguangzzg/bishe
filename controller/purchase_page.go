package controller

import (
	"github.com/gin-gonic/gin"
)

func PurchaseIndexPageHandler(c *gin.Context) {
	RenderIndexPage(c, "purchase/index.html", nil)
}
