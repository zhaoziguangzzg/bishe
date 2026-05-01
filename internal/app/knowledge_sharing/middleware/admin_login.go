package middleware

import (
	"bishe/internal/app/knowledge_sharing/controller"
	"bishe/internal/app/knowledge_sharing/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 页面检查管理员用户登录
func PageAdminUserLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, name, isExpired, err := service.GetAdminUserJwtCookie(c)
		if err != nil {
			c.Redirect(http.StatusFound, service.GetUrlLogin())
			c.Abort()
			return
		}

		if uid == 0 || name == "" || isExpired {
			c.Redirect(http.StatusFound, service.GetUrlLogin())
			c.Abort()
			return
		}

		service.SetAdminUidToContext(c, uid)
		service.SetAdminNameToContext(c, name)

		c.Next()
	}
}

// 接口检查管理员用户登录
func ApiAdminUserLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, name, isExpired, err := service.GetAdminUserJwtCookie(c)
		if err != nil {
			controller.MakeApiResponseError(c, controller.CODE_ADMIN_USER_NOT_LOGIN)
			c.Abort()
			return
		}

		if uid == 0 || name == "" || isExpired {
			controller.MakeApiResponseError(c, controller.CODE_ADMIN_USER_NOT_LOGIN)
			c.Abort()
			return
		}

		service.SetAdminUidToContext(c, uid)
		service.SetAdminNameToContext(c, name)

		c.Next()
	}
}
