package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// OrdersIndexPageHandler
func OrdersIndexPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "orders/index.html", nil)
}
