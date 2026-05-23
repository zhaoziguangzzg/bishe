package service

import (
	"time"

	"github.com/gin-gonic/gin"
)

const (
	CookieUser      = "user"
	CookieAdminUser = "admin_user"
	//管理员用户过期时间
	ADMIN_USER_JWT_EXPIRES_DAYS int = 7
)

// 将用户信息jwt设置到cookie
func SetUserJwtCookie(c *gin.Context, uid int, name string, now time.Time) (err error) {
	t := now.AddDate(0, 0, USER_JWT_EXPIRE_DAYS)
	userJwtStr, err := MakeUserJwt(uid, name, t)
	if err != nil {
		return
	}

	c.SetCookie(CookieUser, userJwtStr, USER_JWT_EXPIRE_DAYS*86400-10, "/", "", true, true)
	return
}

// 获取用户信息
func GetUserCookie(c *gin.Context) (uid int, name string, isExpired bool, err error) {
	userJwtStr, err := c.Cookie(CookieUser)
	if err != nil {
		return
	}

	isExpired, claims, err := ParseUserJwt(userJwtStr)
	if isExpired {
		return
	}

	uid, name = claims.GetUserIdNameFromJwt()
	return
}

// 清除cookie
func DeleteUserCookie(c *gin.Context) {
	c.SetCookie(CookieUser, "", -1, "/", "", true, true)
}

// 将管理员用户信息jwt设置到cookie
func SetAdminUserJwtCookie(c *gin.Context, uid int, name string, now time.Time) (err error) {
	t := now.AddDate(0, 0, ADMIN_USER_JWT_EXPIRES_DAYS)
	adminUserJwtStr, err := MakeAdminUserJwt(uid, name, t)
	if err != nil {
		return
	}

	c.SetCookie(CookieAdminUser, adminUserJwtStr, ADMIN_USER_JWT_EXPIRES_DAYS*86400-10, "/", "", true, true)
	return
}

// 获取用户信息
func GetAdminUserJwtCookie(c *gin.Context) (uid int, name string, isExpired bool, err error) {
	adminUserJwtStr, err := c.Cookie(CookieAdminUser)
	if err != nil {
		return
	}

	isExpired, claims, err := ParseAdminUserJwt(adminUserJwtStr)
	if isExpired {
		return
	}

	uid, name = claims.GetAdminUserIdNameFromJwt()

	return
}

// 清除cookie
func DeleteAdminUserCookie(c *gin.Context) {
	c.SetCookie(CookieAdminUser, "", -1, "/", "", true, true)
}
