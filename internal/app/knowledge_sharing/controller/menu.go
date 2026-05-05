package controller

import (
	"bishe/internal/app/knowledge_sharing/model"
	"bishe/internal/app/knowledge_sharing/service"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 创建权限菜单
func AddMenuHandler(c *gin.Context) {

	name := c.PostForm("menu_name")
	if name == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	path := c.PostForm("menu_path")
	if path == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	//查询name是否存在
	menu, err := service.GetMenuByName(name)
	if err != nil {
		service.Logger.Error("GetMenuByName err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if menu != nil {
		if menu.IsDeleted == model.IS_DELETED_NO {
			MakeApiResponseError(c, CODE_MENU_NAME_EXIST)
			return
		}

		//更新isdelete
		err = service.UpdateMenuNotDeletedById(menu.Id, model.IS_DELETED_NO)
		if err != nil {
			service.Logger.Error("UpdateMenuNotDeletedById err", zap.Error(err))
			MakeApiResponseErrorDefault(c)
			return
		}
		data := map[string]interface{}{
			"menu_id": menu.Id,
		}

		MakeApiResponseSuccess(c, data)
		return

	}

	createTime := time.Now()

	newMenu := &model.Menu{ //其中包含自动生成的id
		MenuName:  name,
		Path:      path,
		CreateAt:  &createTime,
		UpdateAt:  &createTime,
		IsDeleted: model.IS_DELETED_NO,
	}

	err = service.CreateMenu(newMenu)
	if err != nil {
		service.Logger.Error("CreateMenu err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}
	data := map[string]interface{}{
		"menu_id": newMenu.Id,
	}
	MakeApiResponseSuccess(c, data)
}

// 更新菜单
func UpdateMenuHandler(c *gin.Context) {
	idStr := c.PostForm("menu_id")
	if idStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		MakeApiResponseErrorParams(c)
		return
	}

	menuName := c.PostForm("menu_name")
	if menuName == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	path := c.PostForm("menu_path")
	if path == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	menu, err := service.GetMenuNotDeletedById(id)
	if err != nil {
		service.Logger.Error("GetMenuNotDeletedById err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if menu == nil {
		MakeApiResponseError(c, CODE_MENU_NOT_EXIST)
		return
	}

	err = service.UpdateMenuById(id, menuName, path)
	if err != nil {
		service.Logger.Error("UpdateMenuById err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	MakeApiResponseSuccessDefault(c)
}

// 获取全部权限菜单
func GetAllMenuHandler(c *gin.Context) {
	pageStr := c.Query("page")
	page := GetPage(pageStr)

	pagesize := 10

	menus, err := service.GetAllMenu(page, pagesize)
	if err != nil {
		service.Logger.Error("GetAllMenu err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if len(menus) == 0 {
		menus = make([]model.Menu, 0)
	}

	data := map[string]interface{}{
		"menus": menus,
	}

	MakeApiResponseSuccess(c, data)
}

// 删除权限菜单
func DeleteMenuHandler(c *gin.Context) {
	idStr := c.PostForm("menu_id")
	if idStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		MakeApiResponseErrorParams(c)
		return
	}

	//查询菜单是否存在
	menu, err := service.GetMenuNotDeletedById(id)
	if err != nil {
		service.Logger.Error("GetMenuNotDeletedById err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if menu == nil {
		MakeApiResponseError(c, CODE_MENU_NOT_EXIST)
		return
	}

	//更新isdelete
	err = service.UpdateMenuNotDeletedById(id, model.IS_DELETED_YES)
	if err != nil {
		service.Logger.Error("UpdateMenuNotDeletedById err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	MakeApiResponseSuccessDefault(c)

}
