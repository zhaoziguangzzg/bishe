package service

import (
	"bishe/internal/app/knowledge_sharing/dao/mysql"
	"bishe/internal/app/knowledge_sharing/model"
)

// 创建菜单
func CreateMenu(newMenu *model.Menu) (err error) {
	return mysql.CreateMenu(newMenu)
}

// 根据菜单名查询菜单
func GetMenuByName(name string) (menu *model.Menu, err error) {
	return mysql.GetMenuByName(name)
}

// 更新菜单为未删除状态
func UpdateMenuNotDeletedById(id int, isDeleted int) (err error) {
	return mysql.UpdateMenuNotDeletedById(id, isDeleted)
}

// 获取全部权限菜单
func GetAllMenu(page int, pagesize int) (menus []model.Menu, err error) {
	return mysql.GetAllMenu(page, pagesize)
}

// 根据菜单ID查询菜单
func GetMenuNotDeletedById(id int) (menu *model.Menu, err error) {
	return mysql.GetMenuNotDeletedById(id)
}
