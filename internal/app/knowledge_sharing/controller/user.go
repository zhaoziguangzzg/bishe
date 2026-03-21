package controller

import (
	"bishe/internal/app/knowledge_sharing/model"
	"bishe/internal/app/knowledge_sharing/service"
	"bishe/internal/app/knowledge_sharing/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 通过post查询参数添加用户的处理函数
func AddUserHandler(c *gin.Context) {
	// 从表单中获取用户信息
	name := c.PostForm("name")
	password := c.PostForm("password")

	//验证 name超长
	nameLen := len(name)
	if nameLen == 0 || nameLen > 20 {
		MakeApiResponseError(c, CODE_USER_NAME_LEN_INVALID)
		return
	}

	if password == "" || len(password) < 8 {
		MakeApiResponseError(c, CODE_PARAMS_ERROR)
		return
	}

	if !utils.IsValidPassword(password) {
		MakeApiResponseError(c, CODE_PARAMS_ERROR)
		return
	}

	createTime := time.Now()

	//查询用户是否存在
	user, err := service.GetUserByName(name)
	if err != nil {
		service.Logger.Error("GetUserByName", zap.Error(err))
		MakeApiResponseError(c, CODE_SYS_ERROR)
		return
	}

	if user != nil {
		if user.IsDeleted == model.USER_NOT_DELETED {
			MakeApiResponseError(c, CODE_USER_NAME_EXIST)
			return
		}

		MakeApiResponseError(c, CODE_USER_REPLACE)
		return
	}

	// 构造用户对象
	newUser := &model.User{ //其中包含自动生成的id
		Name:       name,
		Password:   password,
		CreateAt:   &createTime,
		UpdateAt:   &createTime,
		UserStatus: model.USER_STATUS_NORMAL,
		IsDeleted:  model.USER_NOT_DELETED,
	}

	// 插入数据库
	err = service.CreateUser(newUser)
	if err != nil {
		service.Logger.Error("CreateUser err", zap.Error(err))
		MakeApiResponseError(c, CODE_USER_NAME_EXIST)
		return
	}

	service.SetUserCookie(c, newUser.Id, name)

	// 返回成功响应
	MakeApiResponseSuccessDefault(c)
}

// POST /api/user/login
func UserLoginHandler(c *gin.Context) {
	// 从表单中获取用户信息
	name := c.PostForm("name")
	password := c.PostForm("password")

	// 数据验证
	// 验证长度
	nameLen := len(name)
	if nameLen == 0 || nameLen > 20 {
		MakeApiResponseError(c, CODE_USER_NAME_LEN_INVALID)
		return
	}

	if password == "" || len(password) < 8 {
		MakeApiResponseError(c, CODE_PARAMS_ERROR)
		return
	}

	if !utils.IsValidPassword(password) {
		MakeApiResponseError(c, CODE_PARAMS_ERROR)
		return
	}

	//根据姓名获取用户信息
	user, err := service.GetUserByName(name)
	if err != nil {
		service.Logger.Error("GetUserByName", zap.Error(err))
		MakeApiResponseError(c, CODE_SYS_ERROR)
		return
	}

	if user == nil {
		MakeApiResponseError(c, CODE_USER_NAME_NOT_EXIST)
		return
	}

	//验证密码是否正确
	if password != user.Password {
		MakeApiResponseError(c, CODE_USER_PASSWORD_INVALID)
		return
	}

	service.SetUserCookie(c, user.Id, name)

	MakeApiResponseSuccessDefault(c)
}

// 退出登录
func UserLogoutHandler(c *gin.Context) {
	//清除cookie
	service.DeleteUserCookie(c)
	MakeApiResponseSuccessDefault(c)
}

// 获取用户
func GetUserHandler(c *gin.Context) {
	//从cookie获取用户信息
	uid, name := service.GetUserFromCookie(c)
	if uid == 0 || name == "" {
		//用户未登录
		MakeApiResponseError(c, CODE_USER_NOT_LOGIN)
		return
	}

	//从数据库获取用户信息
	user, err := service.GetUserByUserId(uid)
	if err != nil {
		service.Logger.Error("GetUserByUserId", zap.Error(err))
		MakeApiResponseError(c, CODE_SYS_ERROR)
		return
	}

	if user == nil {
		MakeApiResponseError(c, CODE_USER_NOT_LOGIN)
		return
	}

	data := map[string]interface{}{
		"user": user,
	}

	MakeApiResponseSuccess(c, data)

}

// 更新用户信息
func UpdateUserHandler(c *gin.Context) {
	//从cookie获取用户登录信息，是验证登录
	uid, name := service.GetUserFromCookie(c)
	if uid == 0 || name == "" {
		//用户未登录
		MakeApiResponseError(c, CODE_USER_NOT_LOGIN)
		return
	}

	userName := c.PostForm("name")
	email := c.PostForm("email")
	ageStr := c.PostForm("age")
	phoneStr := c.PostForm("phone")

	//检测name超长，name="",
	userNameLen := len(userName)
	if userNameLen == 0 || userNameLen > 20 {
		MakeApiResponseError(c, CODE_USER_NAME_LEN_INVALID)
		return
	}

	//检测email格式错误
	if !utils.IsValidEmail(email) {
		MakeApiResponseError(c, CODE_USER_EMAIL_INVALID)
		return
	}

	age, err := strconv.Atoi(ageStr)
	if err != nil {
		MakeApiResponseError(c, CODE_USER_AGE_INVALID)
		return
	}

	if age > model.USER_MAX_AGE || age == 0 {
		MakeApiResponseError(c, CODE_USER_AGE_INVALID)
		return
	}

	//检测手机号长度11位
	if len(phoneStr) != 11 {
		MakeApiResponseError(c, CODE_USER_PHONE_INVALID)
		return
	}

	phone, err := strconv.Atoi(phoneStr)
	if err != nil {
		MakeApiResponseError(c, CODE_USER_PHONE_INVALID)
		return
	}

	//根据id获取用户
	user, err := service.GetUserByUserId(uid)
	if err != nil {
		service.Logger.Error("GetUserByUserId", zap.Error(err))
		MakeApiResponseError(c, CODE_SYS_ERROR)
		return
	}

	//用户cookie有问题，重新登录
	if user == nil {
		// 清除cookie
		service.DeleteUserCookie(c)
		MakeApiResponseError(c, CODE_USER_NOT_LOGIN)
		return
	}

	//更新用户信息
	affectRows, err := service.UpdateUserByUid(uid, userName, email, age, phone)
	if !(affectRows > 0 && err == nil) {
		service.Logger.Error("UpdateUserByUid err", zap.Error(err))
		MakeApiResponseError(c, CODE_SYS_ERROR)
		return
	}

	//修改cookie中的用户名
	service.SetUserCookie(c, uid, userName)

	MakeApiResponseSuccessDefault(c)

}
