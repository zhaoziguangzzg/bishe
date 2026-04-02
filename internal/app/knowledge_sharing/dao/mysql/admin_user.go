package mysql

import (
	"bishe/internal/app/knowledge_sharing/model"

	"gorm.io/gorm"
)

// 添加管理员用户
func CreateAdminUser(newAdminUser *model.AdminUser) (err error) {
	err = DB.Model(&model.AdminUser{}).Create(newAdminUser).Error
	return
}

// 更新管理员用户
func UpdateAdminUserByUid(uid int, name string, email string, phone int) (int64, error) {
	user := model.AdminUser{
		Name:  name,
		Email: email,
		Phone: phone,
	}
	result := DB.Model(&model.AdminUser{}).Where("id=?", uid).Updates(user)
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
	err = DB.Model(&model.AdminUser{}).Where("id=? and is_deleted=?", uid, model.IS_DELETED_NO).First(&adminUser).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound { //没查到数据返回空
			return nil, nil
		}

		return nil, err
	}

	return adminUser, nil
}
