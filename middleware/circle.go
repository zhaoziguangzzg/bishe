package middleware

import (
	"bishe/controller"
	"bishe/model"
	"bishe/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 检查是否加入圈子
func PageIsJoinCircle() gin.HandlerFunc {
	return func(c *gin.Context) {
		var cidStr string
		if c.Request.Method == http.MethodGet {
			cidStr = c.Query("circle_id")
			if cidStr == "" {
				cidStr = c.Query("cid")
			}
		} else {
			cidStr = c.PostForm("circle_id")
			if cidStr == "" {
				cidStr = c.PostForm("cid")
			}
		}

		if cidStr == "" {
			c.Redirect(http.StatusFound, service.GetUrlIndex())
			c.Abort()
			return
		}

		cid, err := strconv.Atoi(cidStr)
		if err != nil {
			c.Redirect(http.StatusFound, service.GetUrlIndex())
			c.Abort()
			return
		}

		service.SetCidToContext(c, cid)

		circle, err := service.GetCircleByCid(cid)
		if err != nil {
			c.Redirect(http.StatusFound, service.GetUrlIndex())
			c.Abort()
			return
		}

		service.SetCircleToContext(c, circle)

		uid := service.GetUidFromContext(c)

		if uid == circle.CircleOwnerId {
			service.SetIsJoinCircleToContext(c, true)
			c.Next()
			return
		}

		join, err := service.GetUserCircleJoinByUidCid(uid, cid)
		if err != nil {
			c.Redirect(http.StatusFound, service.GetUrlIndex())
			c.Abort()
			return
		}

		if join == nil {
			c.Redirect(http.StatusFound, service.GetUrlIndex())
			c.Abort()
			return
		}

		if join.NotJoinStatus != model.USER_CIRCLE_JOIN_STATUS_JOIN {
			c.Redirect(http.StatusFound, service.GetUrlIndex())
			c.Abort()
			return
		}

		service.SetIsJoinCircleToContext(c, true)

		// ========= 【2】调用后续的中间件/控制器 =========
		c.Next()

		// ========= 【3】请求后执行 =========
	}
}

// 检查是否加入圈子
func ApiIsJoinCircle() gin.HandlerFunc {
	return func(c *gin.Context) {
		var cidStr string
		if c.Request.Method == http.MethodGet {
			cidStr = c.Query("circle_id")
			if cidStr == "" {
				cidStr = c.Query("cid")
			}
		} else {
			cidStr = c.PostForm("circle_id")
			if cidStr == "" {
				cidStr = c.PostForm("cid")
			}
		}

		if cidStr == "" {
			controller.MakeApiResponseErrorParams(c)
			c.Abort()
			return
		}

		cid, err := strconv.Atoi(cidStr)
		if err != nil {
			controller.MakeApiResponseErrorParams(c)
			c.Abort()
			return
		}

		service.SetCidToContext(c, cid)

		circle, err := service.GetCircleByCid(cid)
		if err != nil {
			controller.MakeApiResponseErrorDefault(c)
			c.Abort()
			return
		}

		service.SetCircleToContext(c, circle)

		uid := service.GetUidFromContext(c)

		if uid == circle.CircleOwnerId {
			service.SetIsJoinCircleToContext(c, true)
			c.Next()
			return
		}

		join, err := service.GetUserCircleJoinByUidCid(uid, cid)
		if err != nil {
			controller.MakeApiResponseErrorDefault(c)
			c.Abort()
			return
		}

		if join == nil {
			controller.MakeApiResponseError(c, controller.CODE_USER_NOT_JOIN_CIRCLE)
			c.Abort()
			return
		}

		if join.NotJoinStatus != model.USER_CIRCLE_JOIN_STATUS_JOIN {
			controller.MakeApiResponseError(c, controller.CODE_USER_NOT_JOIN_CIRCLE)
			c.Abort()
			return
		}

		service.SetIsJoinCircleToContext(c, true)

		// ========= 【2】调用后续的中间件/控制器 =========
		c.Next()

		// ========= 【3】请求后执行 =========
	}
}

// 检查是否是圈主
func ApiIsCircleOwner() gin.HandlerFunc {
	return func(c *gin.Context) {
		circle, _ := service.GetCircleFromContext(c)

		uid := service.GetUidFromContext(c)

		if uid != circle.CircleOwnerId {
			controller.MakeApiResponseError(c, controller.CODE_USER_NOT_CIRCLE_OWNER)
			c.Abort()
			return
		}

		service.SetIsCircleOwnerToContext(c, true)

		// ========= 【2】调用后续的中间件/控制器 =========
		c.Next()

		// ========= 【3】请求后执行 =========
	}
}

// 检查是否是圈主
func PageIsCircleOwner() gin.HandlerFunc {
	return func(c *gin.Context) {
		circle, ok := service.GetCircleFromContext(c)
		if !ok {
			c.Redirect(http.StatusFound, service.GetUrlIndex())
			c.Abort()
			return
		}

		uid := service.GetUidFromContext(c)

		if uid != circle.CircleOwnerId {
			c.Redirect(http.StatusFound, service.GetUrlCircleIndex(circle.Id))
			c.Abort()
			return
		}

		service.SetIsCircleOwnerToContext(c, true)

		// ========= 【2】调用后续的中间件/控制器 =========
		c.Next()

		// ========= 【3】请求后执行 =========
	}
}
