package middleware

import (
	"bishe/internal/app/knowledge_sharing/controller"
	"bishe/internal/app/knowledge_sharing/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 接口检查用户登录
func MiddlewareUserLoginApi() gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, name, isExpired, err := service.GetUserCookie(c)
		if err != nil {
			controller.MakeApiResponseErrorDefault(c)
			c.Abort()
			return
		}

		if isExpired || uid == 0 || name == "" {
			//用户未登录
			controller.MakeApiResponseError(c, controller.CODE_USER_NOT_LOGIN)
			c.Abort()
			return
		}

		c.Set("uid", uid)
		c.Set("name", name)
		// ========= 【2】调用后续的中间件/控制器 =========
		c.Next()

		// ========= 【3】请求后执行 =========
	}
}

// 在页面检查用户登录
func MiddlewareUserLoginPage() gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, name, isExpired, err := service.GetUserCookie(c)
		if err != nil {
			//错误就跳转到指定页面
			c.Redirect(http.StatusFound, service.GetUrlLogin())
			c.Abort()
			return
		}

		if isExpired || uid == 0 || name == "" {
			//用户未登录
			c.Redirect(http.StatusFound, service.GetUrlLogin())
			c.Abort()
			return
		}

		c.Set("uid", uid)
		c.Set("name", name)
		// ========= 【2】调用后续的中间件/控制器 =========
		c.Next()

		// ========= 【3】请求后执行 =========
	}
}
