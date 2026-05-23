package service

import (
	"bishe/dao/mysql"
	"bishe/model"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	// 密码的salt
	USER_PASSWORD_SALT string = "asdgsfadf"

	//用户jwt的密钥
	USER_JWT_KEY string = "asdfardadf"

	//用户jwt的过期天数
	USER_JWT_EXPIRE_DAYS int = 7
)

// 生成新密码
func MakeUserPassword(str string) (md5Str string) {
	newStr := str + USER_PASSWORD_SALT
	md5Str = MakeMd5(newStr)
	return
}

// 用户信息生成jwt
type UserJwtClaims struct {
	UserId   int    `json:"userId"`
	UserName string `json:"userName"`
	jwt.RegisteredClaims
}

func (userJwtClaims *UserJwtClaims) GetUserIdNameFromJwt() (userId int, userName string) {
	return userJwtClaims.UserId, userJwtClaims.UserName
}

// 生成token
func MakeUserJwt(userId int, userName string, t time.Time) (string, error) {
	key := []byte(USER_JWT_KEY)

	claims := UserJwtClaims{
		UserId:   userId,
		UserName: userName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(t),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(key)
}

// 验证Token
func ParseUserJwt(tokenStr string) (isExpired bool, claims *UserJwtClaims, err error) {
	key := []byte(USER_JWT_KEY)

	token, err := jwt.ParseWithClaims(tokenStr, &UserJwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		// 过期专属错误
		if errors.Is(err, jwt.ErrTokenExpired) {
			return true, nil, nil
		}
		return
	}

	if claims, ok := token.Claims.(*UserJwtClaims); ok && token.Valid {
		return false, claims, nil
	}
	return false, nil, err
}

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

func GetUserByUidList(uidList []int) (userMap map[int]model.User, err error) {
	return mysql.GetUserMapByUids(uidList)
}

// 更新
func UpdateUserByUid(uid int, updateMap map[string]interface{}) (int64, error) {
	return mysql.UpdateUserByUid(uid, updateMap)
}

// 更新用户密码
func UpdateUserPasswordByUid(uid int, password string) (int64, error) {
	return mysql.UpdateUserPasswordByUid(uid, password)
}

// 根据uids获取userMap
func GetUserMapByUids(uids []int) (userMap map[int]model.User, err error) {
	return mysql.GetUserMapByUids(uids)
}
