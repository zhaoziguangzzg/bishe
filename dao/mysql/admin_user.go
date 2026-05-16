package mysql

import (
	"bishe/model"

	"gorm.io/gorm"
)

// 添加管理员用户
func CreateAdminUser(newAdminUser *model.AdminUser) (err error) {
	err = DB.Model(&model.AdminUser{}).Create(newAdminUser).Error
	return
}

// 更新管理员用户
func UpdateAdminUserByUid(uid int, updateMap map[string]interface{}) (int64, error) {
	result := DB.Model(&model.AdminUser{}).Where("id=?", uid).Updates(updateMap)
	return result.RowsAffected, result.Error
}

// 根据name获取管理员用户
func GetAdminUserByName(name string) (adminUser *model.AdminUser, err error) {
	adminUser = &model.AdminUser{}
	err = DB.Model(&model.AdminUser{}).Where("name=? and is_deleted=?", name, model.IS_DELETED_NO).First(&adminUser).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound { //没查到数据返回空
			return nil, nil
		}

		return nil, err
	}

	return adminUser, nil
}

// 根据uid获取管理员用户
func GetAdminUserByUserId(uid int) (adminUser *model.AdminUser, err error) {
	adminUser = new(model.AdminUser)
	err = DB.Model(&model.AdminUser{}).
		Where("id=?", uid).First(&adminUser).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound { //没查到数据返回空
			return nil, nil
		}

		return nil, err
	}

	return adminUser, nil
}

// 获取所有管理员用户
func GetAllAdminUser(page int, pagesize int) (adminUsers []model.AdminUser, err error) {
	offset := (page - 1) * pagesize

	err = DB.Model(&model.AdminUser{}).
		Where("is_deleted=?", model.IS_DELETED_NO).
		Order("id DESC").Offset(offset).
		Limit(pagesize).Find(&adminUsers).Error
	if err != nil {
		return
	}

	return
}

// 更新管理员用户角色
func UpdateAdminUserRoleId(uid int, roleId int) (int64, error) {
	result := DB.Model(&model.AdminUser{}).Where("id=?", uid).Update("role_id", roleId)
	return result.RowsAffected, result.Error
}

// 更新IsDeleted删除
func UpdateAdminUserIsDeleted(uid int) (int64, error) {
	result := DB.Model(&model.AdminUser{}).Where("id=?", uid).Update("is_deleted", model.IS_DELETED_YES)
	return result.RowsAffected, result.Error
}
