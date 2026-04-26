package service

import (
	"bishe/internal/app/knowledge_sharing/dao/mysql"
	"bishe/internal/app/knowledge_sharing/model"
	"bishe/internal/app/knowledge_sharing/utils"
	"errors"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const (
	//管理员用户密码盐
	ADMIN_USER_PASSWORD_SALT string = "kjhgfsfd"

	//管理员用户密钥
	ADMIN_USER_JWT_KEY string = "yjtfcvkyt"

	//管理员用户过期时间
	ADMIN_USER_JWT_EXPIRES_DAYS int = 7
)

// 生成新密码
func MakeAdminUserPassword(str string) (md5Str string) {
	newStr := str + ADMIN_USER_PASSWORD_SALT
	md5Str = utils.MakeMd5(newStr)
	return
}

// 管理员用户信息生成jwt结构体
type AdminUserJwtClaims struct {
	AdminUserId   int    `json:"adminUserId"`
	AdminUserName string `json:"adminUserName"`
	jwt.RegisteredClaims
}

func (adminUserJwtClaims *AdminUserJwtClaims) GetAdminUserIdNameFromJwt() (adminUserId int, adminUserName string) {
	return adminUserJwtClaims.AdminUserId, adminUserJwtClaims.AdminUserName
}

// 生成管理员用户jwt
func MakeAdminUserJwt(adminUserId int, adminUserName string, t time.Time) (string, error) {
	key := []byte(ADMIN_USER_JWT_KEY)

	claims := AdminUserJwtClaims{
		AdminUserId:   adminUserId,
		AdminUserName: adminUserName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(t),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(key)
}

// 解析管理员用户jwt
func ParseAdminUserJwt(tokenStr string) (isExpired bool, claims *AdminUserJwtClaims, err error) {
	key := []byte(ADMIN_USER_JWT_KEY)

	token, err := jwt.ParseWithClaims(tokenStr, &AdminUserJwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		// 过期专属错误
		if errors.Is(err, jwt.ErrTokenExpired) {
			return true, nil, nil
		}
		return
	}

	if claims, ok := token.Claims.(*AdminUserJwtClaims); ok && token.Valid {
		return false, claims, nil
	}
	return false, nil, err
}

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

/*
// 将管理员用户信息设置到cookie
func SetAdminUserCookie(c *gin.Context, uid int, name string) {
	c.SetCookie("adminUserId", strconv.Itoa(uid), 7*86400, "/", "", true, true)
	c.SetCookie("adminUserName", name, 7*86400, "/", "", true, true)
}
*/

// 将管理员用户信息jwt设置到cookie
func SetAdminUserJwtCookie(c *gin.Context, uid int, name string, now time.Time) (err error) {
	t := now.AddDate(0, 0, ADMIN_USER_JWT_EXPIRES_DAYS)
	adminUserJwtStr, err := MakeAdminUserJwt(uid, name, t)
	if err != nil {
		return
	}

	c.SetCookie("admin_user", adminUserJwtStr, ADMIN_USER_JWT_EXPIRES_DAYS*86400-10, "/", "", true, true)
	return
}

// 获取用户信息
func GetAdminUserJwtCookie(c *gin.Context) (uid int, name string, isExpired bool, err error) {
	adminUserJwtStr, err := c.Cookie("admin_user")
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
