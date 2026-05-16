package middleware

import (
	"bishe/controller"
	"bishe/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 接口检查用户登录
func ApiUserLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 登录、注册、退出不校验登录状态
		path := c.Request.URL.Path
		if path == "/api/user/login" || path == "/api/user/add" || path == "/api/user/logout" {
			c.Next()
			return
		}

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

		service.SetUidToContext(c, uid)
		service.SetNameToContext(c, name)
		// ========= 【2】调用后续的中间件/控制器 =========
		c.Next()

		// ========= 【3】请求后执行 =========
	}
}

// 在页面检查用户登录
func PageUserLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 登录、注册页面不校验登录状态
		path := c.Request.URL.Path
		if path == "/page/user/login" || path == "/page/user/register" {
			c.Next()
			return
		}

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

		service.SetUidToContext(c, uid)
		service.SetNameToContext(c, name)
		// ========= 【2】调用后续的中间件/控制器 =========
		c.Next()

		// ========= 【3】请求后执行 =========
	}
}
