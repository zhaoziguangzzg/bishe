package mysql

import (
	"bishe/internal/app/knowledge_sharing/model"

	"gorm.io/gorm"
)

// 创建权限角色
func CreateRole(newRole *model.Role) (err error) {
	err = DB.Model(&model.Role{}).Create(newRole).Error
	return
}

// 根据权限名查询权限角色
func GetRoleByName(name string) (role *model.Role, err error) {
	role = new(model.Role)
	err = DB.Model(&model.Role{}).Where("role_name=?", name).First(&role).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound { //没查到数据返回空
			return nil, nil
		}

		return nil, err
	}

	return role, nil
}

// 更新角色为未删除状态
func UpdateRoleNotDeletedById(id int, isDeleted int) (err error) {
	err = DB.Model(&model.Role{}).Where("id=?", id).Update("is_deleted", isDeleted).Error
	return
}

// 更新角色权限
func UpdateRoleById(id int, roleMap map[string]interface{}) (int64, error) {
	result := DB.Model(&model.Role{}).Where("id=?", id).Updates(roleMap)
	return result.RowsAffected, result.Error
}

// 获取全部权限角色
func GetAllRole(page int, pagesize int) (roles []model.Role, err error) {
	offset := (page - 1) * pagesize

	err = DB.Model(&model.Role{}).
		Where("is_deleted=?", model.IS_DELETED_NO).
		Order("id DESC").Offset(offset).Limit(pagesize).Find(&roles).Error
	return
}

// 根据角色ID查询权限角色信息
func GetRoleNotDeletedById(id int) (role *model.Role, err error) {
	role = new(model.Role)
	err = DB.Model(&model.Role{}).Where("id=? and is_deleted=?", id, model.IS_DELETED_NO).First(&role).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound { //没查到数据返回空
			return nil, nil
		}

		return nil, err
	}

	return role, nil
}
