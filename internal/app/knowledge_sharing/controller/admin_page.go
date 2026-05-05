package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminIndexPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/index.html", nil)
}

func AdminEditPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/edit.html", nil)
}

func AdminRoleEditPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "adminrole/edit.html", nil)
}

func MenuAddPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "adminmenu/add.html", nil)
}

func MenuEditPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "adminmenu/edit.html", nil)
}

func MenuListPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "adminmenu/list.html", nil)
}

func RoleAddPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "adminrole/add.html", nil)
}

func RoleListPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "adminrole/list.html", nil)
}

func RoleDetailPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "adminrole/detail.html", nil)
}

func AdminUserAddPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "adminuser/add.html", nil)
}

func AdminUserListPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "adminuser/list.html", nil)
}
