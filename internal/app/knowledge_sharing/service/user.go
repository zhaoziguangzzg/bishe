package service

import (
	"bishe/internal/app/knowledge_sharing/dao/mysql"
	"bishe/internal/app/knowledge_sharing/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

// create用户
func CreateUser(newUser *model.User) (err error) {
	return mysql.CreateUser(newUser)
}

// 根据name获取用户
func GetUserByName(name string) (user *model.User, err error) {
	return mysql.GetUserByName(name)
}

// 根据id获取用户
func GetUserByUserId(UserId int) (user *model.User, err error) {
	return mysql.GetUserByUserId(UserId)
}

// 更新
func UpdateUserByUid(uid int, updateMap map[string]interface{}) (int64, error) {
	return mysql.UpdateUserByUid(uid, updateMap)
}

// 更新用户密码
func UpdateUserPasswordByUid(uid int, password string) (int64, error) {
	return mysql.UpdateUserPasswordByUid(uid, password)
}

// 将用户信息设置到cookie
func SetUserCookie(c *gin.Context, uid int, name string) {
	c.SetCookie("userId", strconv.Itoa(uid), 7*86400, "/", "", true, true)
	c.SetCookie("userName", name, 7*86400, "/", "", true, true)
}

// 获取cookie用户信息
func GetUserFromCookie(c *gin.Context) (uid int, name string) {
	uidStr, err := c.Cookie("userId")
	if err != nil {
		//0,""z
		return
	}

	uid, err = strconv.Atoi(uidStr)
	if err != nil {
		//0,""
		return
	}

	name, err = c.Cookie("userName")
	if err != nil {
		//uid,""
		return
	}
	//uid,name
	return
}

// 清除cookie
func DeleteUserCookie(c *gin.Context) {
	c.SetCookie("userId", "", -1, "/", "", true, true)
	c.SetCookie("userName", "", -1, "/", "", true, true)
}

// 根据uids获取userMap
func GetUserMapByUids(uids []int) (userMap map[int]model.User, err error) {
	return mysql.GetUserMapByUids(uids)
}
