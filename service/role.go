package service

import (
	"bishe/dao/mysql"
	"bishe/model"
)

// 创建权限角色
func CreateRole(newRole *model.Role) (err error) {
	return mysql.CreateRole(newRole)
}

// 根据权限名查询权限角色
func GetRoleByName(name string) (role *model.Role, err error) {
	return mysql.GetRoleByName(name)
}

// 更新角色为未删除状态
func UpdateRoleNotDeletedById(id int, isDeleted int) (err error) {
	return mysql.UpdateRoleNotDeletedById(id, isDeleted)
}

// 更新角色权限
func UpdateRoleById(id int, roleMap map[string]interface{}) (int64, error) {
	return mysql.UpdateRoleById(id, roleMap)
}

// 获取全部权限角色
func GetAllRole(page int, pagesize int) (roles []model.Role, err error) {
	offset := (page - 1) * pagesize
	return mysql.GetAllRole(offset, pagesize)
}

// 根据角色ID查询权限角色信息
func GetRoleNotDeletedById(id int) (role *model.Role, err error) {
	return mysql.GetRoleNotDeletedById(id)
}
