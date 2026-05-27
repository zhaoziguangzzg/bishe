package middleware

import (
	"bishe/model"
	"bishe/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 检查是否购买课程
func PageIsPurchaseCourse() gin.HandlerFunc {
	return func(c *gin.Context) {
		var cidStr string
		if c.Request.Method == http.MethodGet {
			cidStr = c.Query("course_id")
			if cidStr == "" {
				cidStr = c.Query("cid")
			}
		} else {
			cidStr = c.PostForm("course_id")
			if cidStr == "" {
				cidStr = c.PostForm("cid")
			}
		}

		if cidStr == "" {
			c.Redirect(http.StatusFound, service.GetUrlCourseIndex())
			c.Abort()
			return
		}

		cid, err := strconv.Atoi(cidStr)
		if err != nil {
			c.Redirect(http.StatusFound, service.GetUrlCourseIndex())
			c.Abort()
			return
		}

		uid := service.GetUidFromContext(c)

		// 检查课程是否已购买
		purchases, err := service.GetPurchaseByUidCid(uid, cid)
		if err != nil {
			c.Redirect(http.StatusFound, service.GetUrlCourseDetail(cid))
			c.Abort()
			return
		}

		if len(purchases) == 0 {
			c.Redirect(http.StatusFound, service.GetUrlCourseDetail(cid))
			c.Abort()
			return
		}

		for _, purchase := range purchases {
			if purchase.PurchaseStatus == model.PURCHASE_ACTION_PAY {
				c.Next()
				return
			}
		}

		c.Redirect(http.StatusFound, service.GetUrlCourseDetail(cid))
		c.Abort()
	}
}
