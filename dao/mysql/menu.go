package mysql

import (
	"bishe/model"

	"gorm.io/gorm"
)

// 创建权限菜单
func CreateMenu(newMenu *model.Menu) (err error) {
	err = DB.Model(&model.Menu{}).Create(newMenu).Error
	return
}

// 根据权限名查询权限菜单
func GetMenuByName(name string) (menu *model.Menu, err error) {
	menu = new(model.Menu)
	err = DB.Model(&model.Menu{}).Where("menu_name=?", name).First(&menu).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound { //没查到数据返回空
			return nil, nil
		}

		return nil, err
	}

	return menu, nil
}

// 更新菜单为未删除状态
func UpdateMenuNotDeletedById(id int, isDeleted int) (err error) {
	err = DB.Model(&model.Menu{}).Where("id=?", id).Update("is_deleted", isDeleted).Error
	return
}

// 获取全部权限菜单
func GetAllMenu() (menus []model.Menu, err error) {
	err = DB.Model(&model.Menu{}).
		Where("is_deleted=?", model.IS_DELETED_NO).
		Order("weight ASC").
		Find(&menus).Error
	return
}

// 根据菜单ID查询菜单
func GetMenuNotDeletedById(id int) (menu *model.Menu, err error) {
	menu = new(model.Menu)
	err = DB.Model(&model.Menu{}).Where("id=? and is_deleted=?", id, model.IS_DELETED_NO).First(&menu).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound { //没查到数据返回空
			return nil, nil
		}

		return nil, err
	}

	return menu, nil
}

// 更新菜单信息
func UpdateMenuById(id int, menuName string, path string) (err error) {
	err = DB.Model(&model.Menu{}).Where("id=?", id).Updates(map[string]interface{}{
		"menu_name": menuName,
		"path":      path,
	}).Error
	return
}
