package controller

import (
	"github.com/gin-gonic/gin"
)

func AdminIndexPageHandler(c *gin.Context) {
	RenderAdminPage(c, "admin/index.html", nil)
}

func MenuAddPageHandler(c *gin.Context) {
	RenderAdminPage(c, "adminmenu/add.html", nil)
}

func MenuEditPageHandler(c *gin.Context) {
	RenderAdminPage(c, "adminmenu/edit.html", nil)
}

func MenuListPageHandler(c *gin.Context) {
	RenderAdminPage(c, "adminmenu/list.html", nil)
}

func RoleAddPageHandler(c *gin.Context) {
	RenderAdminPage(c, "adminrole/add.html", nil)
}

func RoleDetailPageHandler(c *gin.Context) {
	RenderAdminPage(c, "adminrole/detail.html", nil)
}

func AdminRoleEditPageHandler(c *gin.Context) {
	RenderAdminPage(c, "adminrole/edit.html", nil)
}

func RoleListPageHandler(c *gin.Context) {
	RenderAdminPage(c, "adminrole/list.html", nil)
}

func AdminUserListPageHandler(c *gin.Context) {
	RenderAdminPage(c, "adminuser/list.html", nil)
}

func AdminUserAddPageHandler(c *gin.Context) {
	RenderAdminPage(c, "adminuser/add.html", nil)
}

func AdminUserDetailPageHandler(c *gin.Context) {
	RenderAdminPage(c, "adminuser/detail.html", nil)
}

func AdminUserEditPageHandler(c *gin.Context) {
	RenderAdminPage(c, "adminuser/edit.html", nil)
}

func AdminUserRolePageHandler(c *gin.Context) {
	RenderAdminPage(c, "adminuser/role.html", nil)
}

func AdminAccusationListPageHandler(c *gin.Context) {
	RenderAdminPage(c, "adminaccusation/list.html", nil)
}

func AdminAccusationEditPageHandler(c *gin.Context) {
	RenderAdminPage(c, "adminaccusation/edit.html", nil)
}

func AdminFeedbackListPageHandler(c *gin.Context) {
	RenderAdminPage(c, "adminfeedback/list.html", nil)
}

func AdminAdvertListPageHandler(c *gin.Context) {
	RenderAdminPage(c, "adminadvert/list.html", nil)
}

func AdminAdvertAddPageHandler(c *gin.Context) {
	RenderAdminPage(c, "adminadvert/add.html", nil)
}

func AdvertEditPageHandler(c *gin.Context) {
	RenderAdminPage(c, "adminadvert/edit.html", nil)
}

func AdminAnnounceAddPageHandler(c *gin.Context) {
	RenderAdminPage(c, "adminannounce/add.html", nil)
}

func AdminAnnounceListPageHandler(c *gin.Context) {
	RenderAdminPage(c, "adminannounce/list.html", nil)
}

func AnnounceEditPageHandler(c *gin.Context) {
	RenderAdminPage(c, "announce/edit.html", nil)
}
