package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PurchaseIndexPageHandler(c *gin.Context) {
	RenderIndexPage(c, "purchase/index.html", nil)
}

func PurchasePayPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "purchase/pay.html", nil)
}
