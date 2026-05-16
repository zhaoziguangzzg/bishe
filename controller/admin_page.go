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

func AdminUserDetailPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "adminuser/detail.html", nil)
}

func AdminUserRolePageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "adminuser/role.html", nil)
}

func AdminAccusationListPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "adminaccusation/list.html", nil)
}

func AdminFeedbackListPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "adminfeedback/list.html", nil)
}

func AdminAnnounceAddPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "adminannounce/add.html", nil)
}

func AdminAnnounceListPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "adminannounce/list.html", nil)
}

func AdminAdvertAddPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "adminadvert/add.html", nil)
}

func AdminAdvertListPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "adminadvert/list.html", nil)
}
