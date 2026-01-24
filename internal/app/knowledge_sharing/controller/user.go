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
func AddUserHandler(c *gin.Context) { //c

	// 从表单中获取用户信息
	accountStr := c.PostForm("account")
	password := c.PostForm("password")

	// 数据验证
	if accountStr == "" {
		MakeApiResponseError(c, CODE_PARAMS_ERROR)
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

	account, err := strconv.Atoi(accountStr)
	if err != nil || account < 0 {
		service.Logger.Error("accountStrAtoi err", zap.Error(err))
		MakeApiResponseError(c, CODE_PARAMS_ERROR)
		return
	}

	// 构造用户对象
	newUser := &model.User{ //其中包含自动生成的id
		Account:  account,
		Password: password,
	}

	// 插入数据库
	err = service.CreateUser(newUser)
	if err != nil {
		service.Logger.Error("CreateUser err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	// 返回成功响应
	MakeApiResponseSuccess(c, CODE_SUCCESS)
}

// POST /api/user/login
func GetUserHandler(c *gin.Context) {
	// 从表单中获取用户信息
	accountStr := c.PostForm("account")
	password := c.PostForm("password")

	// 数据验证
	if accountStr == "" {
		MakeApiResponseError(c, CODE_PARAMS_ERROR)
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

	account, err := strconv.Atoi(accountStr)
	if err != nil || account < 0 {
		MakeApiResponseError(c, CODE_PARAMS_ERROR)
		return
	}

	passwordStr, err := service.GetPasswordFromUser(account)
	if err != nil {
		MakeApiResponseError(c, CODE_PARAMS_ERROR)
		return
	}

	if password != passwordStr {
		MakeApiResponseError(c, CODE_PARAMS_ERROR)
		return
	}

	user, err := service.GetUserByAccount(account)
	if err != nil {
		service.Logger.Error("GetUserByAccount", zap.Error(err))
		MakeApiResponseError(c, CODE_SYS_ERROR)
	}

	createTime := time.Now()

	err = service.MakeAndSendNotice(0, user.Account, "用户已登录成功", createTime)
	if err != nil {
		service.Logger.Error("MakeAndSendNotice err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	HttpCookie(c, "userAccount", strconv.Itoa(user.Account))
	HttpCookie(c, "userId", strconv.Itoa(user.Id))

	MakeApiResponseSuccess(c, CODE_SUCCESS)
}

// 更新用户信息
func UpdateUserHandler(c *gin.Context) {
	userIdStr, err := c.Cookie("userId")
	if err != nil {
		service.Logger.Error("get coolie err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	userId, _ := strconv.Atoi(userIdStr)
	if userId == 0 {
		MakeApiResponseErrorParams(c)
		return
	}
	accountStr := c.PostForm("account")
	password := c.PostForm("password")
	email := c.PostForm("email")
	ageStr := c.PostForm("age")
	phoneStr := c.PostForm("phone")

	account, err := strconv.Atoi(accountStr)
	if err != nil {
		service.Logger.Error("accountAtoi err", zap.Error(err))
		MakeApiResponseError(c, CODE_PARAMS_ERROR)
	}

	age, err := strconv.Atoi(ageStr)
	if err != nil {
		service.Logger.Error("ageAtoi err", zap.Error(err))
		MakeApiResponseError(c, CODE_PARAMS_ERROR)
		return
	}

	phone, err := strconv.Atoi(phoneStr)
	if err != nil {
		service.Logger.Error("phoneAtoi err", zap.Error(err))
		MakeApiResponseError(c, CODE_PARAMS_ERROR)
		return
	}

	user, err := service.GetUserByUserId(userId)
	if err != nil {
		service.Logger.Error("GetUserByUserId", zap.Error(err))
		MakeApiResponseError(c, CODE_SYS_ERROR)
	}

	result := service.UpdateFromUser(user, password, email, age, phone)
	if result.Error != nil {
		service.Logger.Error("UpdateFromUser err", zap.Error(err))
		MakeApiResponseError(c, CODE_SYS_ERROR)
		return
	}

	service.MakeAndSendNotice(0, account, "用户信息已更新", time.Now())

	MakeApiResponseSuccess(c, map[string]interface{}{
		"account":  account,
		"password": password,
		"email":    email,
		"age":      age,
		"phone":    phone,
	})
}
