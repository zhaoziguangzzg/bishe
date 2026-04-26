package middleware

import (
	"bishe/internal/app/knowledge_sharing/controller"
	"bishe/internal/app/knowledge_sharing/model"
	"bishe/internal/app/knowledge_sharing/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 检查是否加入圈子
func MiddlewareIsJoinCirclePage() gin.HandlerFunc {
	return func(c *gin.Context) {
		var cidStr string
		if c.Request.Method == http.MethodGet {
			cidStr = c.Query("circle_id")
		} else {
			cidStr = c.PostForm("circle_id")
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

		c.Set("cid", cid)

		circle, err := service.GetCircleByCid(cid)
		if err != nil {
			c.Redirect(http.StatusFound, service.GetUrlIndex())
			c.Abort()
			return
		}

		c.Set("circle", circle)

		uid := c.GetInt("uid")

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

		c.Set("isJoinCircle", true)

		// ========= 【2】调用后续的中间件/控制器 =========
		c.Next()

		// ========= 【3】请求后执行 =========
	}
}

// 检查是否加入圈子
func MiddlewareIsJoinCircleApi() gin.HandlerFunc {
	return func(c *gin.Context) {
		var cidStr string
		if c.Request.Method == http.MethodGet {
			cidStr = c.Query("circle_id")
		} else {
			cidStr = c.PostForm("circle_id")
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

		uid := c.GetInt("uid")

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

		c.Set("isJoinCircle", true)

		// ========= 【2】调用后续的中间件/控制器 =========
		c.Next()

		// ========= 【3】请求后执行 =========
	}
}

// 检查是否是圈主
func MiddlewareIsCircleOwnerApi() gin.HandlerFunc {
	return func(c *gin.Context) {
		circle, _ := service.GetCircleFromContext(c)

		uid := c.GetInt("uid")

		if uid != circle.CircleOwnerId {
			controller.MakeApiResponseError(c, controller.CODE_USER_NOT_CIRCLE_OWNER)
			c.Abort()
			return
		}

		c.Set("isCircleOwner", true)

		// ========= 【2】调用后续的中间件/控制器 =========
		c.Next()

		// ========= 【3】请求后执行 =========
	}
}

// 检查是否是圈主
func MiddlewareIsCircleOwnerPage() gin.HandlerFunc {
	return func(c *gin.Context) {
		circleAny, ok := c.Get("circle")
		if !ok {
			c.Redirect(http.StatusFound, service.GetUrlIndex())
			c.Abort()
			return
		}

		circle, ok := circleAny.(*model.Circle)
		if !ok {
			c.Redirect(http.StatusFound, service.GetUrlIndex())
			c.Abort()
			return
		}

		uid := c.GetInt("uid")

		if uid != circle.CircleOwnerId {
			c.Redirect(http.StatusFound, service.GetUrlCircleIndex(circle.Id))
			c.Abort()
			return
		}

		c.Set("isCircleOwner", true)

		// ========= 【2】调用后续的中间件/控制器 =========
		c.Next()

		// ========= 【3】请求后执行 =========
	}
}
