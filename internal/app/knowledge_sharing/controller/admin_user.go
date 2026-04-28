package controller

import (
	"bishe/internal/app/knowledge_sharing/model"
	"bishe/internal/app/knowledge_sharing/service"
	"bishe/internal/app/knowledge_sharing/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 通过post查询参数添加用户的处理函数
func AddAdminUserHandler(c *gin.Context) {
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

	//查询用户是否存在
	adminUser, err := service.GetAdminUserByName(name)
	if err != nil {
		service.Logger.Error("GetAdminUserByName", zap.Error(err))
		MakeApiResponseError(c, CODE_SYS_ERROR)
		return
	}

	if adminUser != nil {
		if adminUser.IsDeleted == model.IS_DELETED_NO {
			MakeApiResponseError(c, CODE_USER_NAME_EXIST)
			return
		}

		MakeApiResponseError(c, CODE_USER_REPLACE)
		return
	}

	createTime := time.Now()

	newPassword := service.MakeAdminUserPassword(password)

	// 构造管理员用户
	newAdminUser := &model.AdminUser{ //其中包含自动生成的id
		Name:      name,
		Password:  newPassword,
		CreateAt:  &createTime,
		UpdateAt:  &createTime,
		IsDeleted: model.IS_DELETED_NO,
	}

	// 插入数据库
	err = service.CreateAdminUser(newAdminUser)
	if err != nil {
		service.Logger.Error("CreateAdminUser err", zap.Error(err))
		MakeApiResponseError(c, CODE_USER_NAME_EXIST)
		return
	}

	service.SetAdminUserJwtCookie(c, newAdminUser.Id, name, createTime)

	// 返回成功响应
	MakeApiResponseSuccessDefault(c)
}

// POST /api/user/login
func AdminUserLoginHandler(c *gin.Context) {
	nowTime := time.Now()
	// 从表单中获取用户信息
	name := c.PostForm("name")
	password := c.PostForm("password")

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

	//根据姓名获取管理员用户信息
	user, err := service.GetAdminUserByName(name)
	if err != nil {
		service.Logger.Error("GetAdminUserByName", zap.Error(err))
		MakeApiResponseError(c, CODE_SYS_ERROR)
		return
	}

	if user == nil {
		MakeApiResponseError(c, CODE_USER_NAME_NOT_EXIST)
		return
	}

	newPassword := service.MakeAdminUserPassword(password)

	//验证密码是否正确
	if newPassword != user.Password {
		MakeApiResponseError(c, CODE_PASSWORD_WRONG)
		return
	}

	service.SetAdminUserJwtCookie(c, user.Id, name, nowTime)

	MakeApiResponseSuccessDefault(c)
}

// 退出登录
func AdminUserLogoutHandler(c *gin.Context) {
	//清除cookie
	service.DeleteAdminUserCookie(c)
	MakeApiResponseSuccessDefault(c)
}

// 获取用户
func GetAdminUserHandler(c *gin.Context) {
	//从cookie获取用户信息
	uid := c.GetInt("admin_uid")

	//从数据库获取用户信息
	adminUser, err := service.GetAdminUserByUserId(uid)
	if err != nil {
		service.Logger.Error("GetAdminUserByUserId", zap.Error(err))
		MakeApiResponseError(c, CODE_SYS_ERROR)
		return
	}

	if adminUser == nil {
		MakeApiResponseError(c, CODE_USER_NOT_LOGIN)
		return
	}

	data := map[string]interface{}{
		"user": adminUser,
	}

	MakeApiResponseSuccess(c, data)

}

// 更新用户信息
func UpdateAdminUserHandler(c *gin.Context) {
	//从cookie获取用户登录信息，是验证登录
	uid := c.GetInt("admin_uid")

	userName := c.PostForm("name")
	email := c.PostForm("email")
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
	adminUser, err := service.GetAdminUserByUserId(uid)
	if err != nil {
		service.Logger.Error("GetAdminUserByUserId", zap.Error(err))
		MakeApiResponseError(c, CODE_SYS_ERROR)
		return
	}

	//用户cookie有问题，重新登录
	if adminUser == nil {
		// 清除cookie
		service.DeleteAdminUserCookie(c)
		MakeApiResponseError(c, CODE_USER_NOT_LOGIN)
		return
	}

	updateMap := map[string]interface{}{
		"name":  userName,
		"email": email,
		"phone": phone,
	}

	fileType := service.FILE_TYPE_UAER_AVATAR
	timeNow := time.Now()

	// 处理头像上传
	avatarPath := ""
	file, header, err := c.Request.FormFile("avatar")
	//判断错误不等于无文件
	if err != nil && err != http.ErrMissingFile {
		service.Logger.Error("FormFile err", zap.Error(err))
		MakeApiResponseErrorParams(c)
		return
	}

	//判断size不是空
	if err == nil && header.Size != 0 {
		avatarPath, err = service.FileSave(file, header, fileType, timeNow)
		if err != nil {
			MakeApiResponseErrorDefault(c)
			return
		}
		updateMap["avatar"] = avatarPath
	}

	//更新用户信息
	affectRows, err := service.UpdateAdminUserByUid(uid, updateMap)
	if !(affectRows > 0 && err == nil) {
		service.Logger.Error("UpdateAdminUserByUid err", zap.Error(err))
		MakeApiResponseError(c, CODE_SYS_ERROR)
		return
	}

	//修改cookie中的用户名
	service.SetAdminUserJwtCookie(c, uid, userName, timeNow)

	MakeApiResponseSuccessDefault(c)

}

// 删除管理员用户
func DeleteAdminUserHandler(c *gin.Context) {
	//更新字段
	uidStr := c.PostForm("uid")
	if uidStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	uid, err := strconv.Atoi(uidStr)
	if err != nil {
		MakeApiResponseErrorDefault(c)
		return
	}

	// 更新IsDeleted删除
	affectRows, err := service.UpdateAdminUserIsDeleted(uid)
	if err != nil || affectRows == 0 {
		service.Logger.Error("UpdateAdminUserIsDeleted err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	MakeApiResponseSuccessDefault(c)
}

// 修改管理员用户角色
func UpdateAdminUserRoleHandler(c *gin.Context) {
	uid := c.GetInt("admin_uid")

	roleIdStr := c.PostForm("role_id")
	if roleIdStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	roleId, err := strconv.Atoi(roleIdStr)
	if err != nil {
		MakeApiResponseErrorDefault(c)
		return
	}

	//查询用户是否存在
	adminUser, err := service.GetAdminUserByUserId(uid)
	if err != nil {
		service.Logger.Error("GetAdminUserByUserId", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}
	if adminUser == nil {
		MakeApiResponseError(c, CODE_USER_NOT_LOGIN)
		return
	}

	//查询角色是否存在
	role, err := service.GetRoleNotDeletedById(roleId)
	if err != nil {
		service.Logger.Error("GetRoleNotDeletedById", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if role == nil {
		MakeApiResponseError(c, CODE_ROLE_NOT_EXIST)
		return
	}

	//更新用户角色
	affectRows, err := service.UpdateAdminUserRoleId(uid, roleId)
	if err != nil || affectRows == 0 {
		service.Logger.Error("UpdateAdminUserRoleId err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	MakeApiResponseSuccessDefault(c)
}
