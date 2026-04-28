package controller

import (
	"bishe/internal/app/knowledge_sharing/model"
	"bishe/internal/app/knowledge_sharing/service"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 创建权限角色
func AddRoleHandler(c *gin.Context) {

	name := c.PostForm("role_name")
	if name == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	// TODO传数组，自己拼mids
	mids := c.PostForm("mids")
	if mids == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	//查询name是否存在
	role, err := service.GetRoleByName(name)
	if err != nil {
		service.Logger.Error("GetRoleByName err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if role != nil {
		if role.IsDeleted == model.IS_DELETED_NO {
			MakeApiResponseError(c, CODE_ROLE_NAME_EXIST)
			return
		}

		//更新isdelete
		err = service.UpdateRoleNotDeletedById(role.Id, model.IS_DELETED_NO)
		if err != nil {
			service.Logger.Error("UpdateRoleNotDeletedById err", zap.Error(err))
			MakeApiResponseErrorDefault(c)
			return
		}
		data := map[string]interface{}{
			"role_id": role.Id,
		}

		MakeApiResponseSuccess(c, data)
		return

	}

	createTime := time.Now()

	newRole := &model.Role{ //其中包含自动生成的id
		RoleName:  name,
		Mids:      mids,
		CreateAt:  &createTime,
		UpdateAt:  &createTime,
		IsDeleted: model.IS_DELETED_NO,
	}

	err = service.CreateRole(newRole)
	if err != nil {
		service.Logger.Error("CreateRole err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}
	data := map[string]interface{}{
		"role_id": newRole.Id,
	}
	MakeApiResponseSuccess(c, data)
}

// 更新角色权限

// 获取全部权限角色
func GetAllRoleHandler(c *gin.Context) {
	pageStr := c.Query("page")
	page := GetPage(pageStr)

	pagesize := 10

	roles, err := service.GetAllRole(page, pagesize)
	if err != nil {
		service.Logger.Error("GetAllRole err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if len(roles) == 0 {
		roles = make([]model.Role, 0)
	}

	data := map[string]interface{}{
		"roles": roles,
	}

	MakeApiResponseSuccess(c, data)
}

// 根据角色ID查询权限角色信息
func GetRoleHandler(c *gin.Context) {
	idStr := c.PostForm("role_id")
	if idStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		MakeApiResponseErrorParams(c)
		return
	}

	role, err := service.GetRoleNotDeletedById(id)
	if err != nil {
		service.Logger.Error("GetRoleNotDeletedById err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if role == nil {
		MakeApiResponseError(c, CODE_ROLE_NOT_EXIST)
		return
	}

	data := map[string]interface{}{
		"role": role,
	}

	MakeApiResponseSuccess(c, data)
}

// 删除权限角色
func DeleteRoleHandler(c *gin.Context) {
	idStr := c.PostForm("role_id")
	if idStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		MakeApiResponseErrorParams(c)
		return
	}

	//查询角色是否存在
	role, err := service.GetRoleNotDeletedById(id)
	if err != nil {
		service.Logger.Error("GetRoleNotDeletedById err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if role == nil {
		MakeApiResponseError(c, CODE_ROLE_NOT_EXIST)
		return
	}

	//更新isdelete
	err = service.UpdateRoleNotDeletedById(id, model.IS_DELETED_YES)
	if err != nil {
		service.Logger.Error("UpdateRoleNotDeletedById err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	MakeApiResponseSuccessDefault(c)

}
