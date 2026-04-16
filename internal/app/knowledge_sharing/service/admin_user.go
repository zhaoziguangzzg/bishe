package service

import (
	"bishe/internal/app/knowledge_sharing/dao/mysql"
	"bishe/internal/app/knowledge_sharing/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 添加管理员用户
func CreateAdminUser(newAdminUser *model.AdminUser) (err error) {
	return mysql.CreateAdminUser(newAdminUser)
}

// 更新管理员用户
func UpdateAdminUserByUid(uid int, updateMap map[string]interface{}) (int64, error) {
	return mysql.UpdateAdminUserByUid(uid, updateMap)
}

// 根据name获取管理员用户
func GetAdminUserByName(name string) (adminUser *model.AdminUser, err error) {
	return mysql.GetAdminUserByName(name)
}

// 根据uid获取管理员用户
func GetAdminUserByUserId(uid int) (adminUser *model.AdminUser, err error) {
	return mysql.GetAdminUserByUserId(uid)
}

// 更新IsDeleted删除
func UpdateAdminUserIsDeleted(uid int) (int64, error) {
	return mysql.UpdateAdminUserIsDeleted(uid)
}

// 将管理员用户信息设置到cookie
func SetAdminUserCookie(c *gin.Context, uid int, name string) {
	c.SetCookie("adminUserId", strconv.Itoa(uid), 7*86400, "/", "", true, true)
	c.SetCookie("adminUserName", name, 7*86400, "/", "", true, true)
}

// 获取cookie管理员用户信息
func GetAdminUserFromCookie(c *gin.Context) (uid int, name string) {
	uidStr, err := c.Cookie("adminUserId")
	if err != nil {
		//0,""z
		return
	}

	uid, err = strconv.Atoi(uidStr)
	if err != nil {
		//0,""
		return
	}

	name, err = c.Cookie("adminUserName")
	if err != nil {
		//uid,""
		return
	}
	//uid,name
	return
}

// 清除cookie
func DeleteAdminUserCookie(c *gin.Context) {
	c.SetCookie("adminUserId", "", -1, "/", "", true, true)
	c.SetCookie("adminUserName", "", -1, "/", "", true, true)
}
