package middleware

import (
	"bishe/internal/app/knowledge_sharing/controller"
	"bishe/internal/app/knowledge_sharing/service"

	"github.com/gin-gonic/gin"
)

func UserLoginMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, name := service.GetUserFromCookie(c)
		if uid == 0 || name == "" {
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
