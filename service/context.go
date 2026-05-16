package service

import (
	"github.com/gin-gonic/gin"
)

const (
	ContextKeyUid           = "uid"
	ContextKeyName          = "name"
	ContextKeyIsJoinCircle  = "isJoinCircle"
	ContextKeyIsCircleOwner = "isCircleOwner"
	ContextKeyAdminUid      = "admin_uid"
	ContextKeyAdminName     = "admin_name"
)

func SetUidToContext(c *gin.Context, uid int) {
	c.Set(ContextKeyUid, uid)
}

func GetUidFromContext(c *gin.Context) int {
	return c.GetInt(ContextKeyUid)
}

func SetNameToContext(c *gin.Context, name string) {
	c.Set(ContextKeyName, name)
}

func GetNameFromContext(c *gin.Context) string {
	return c.GetString(ContextKeyName)
}

func SetIsJoinCircleToContext(c *gin.Context, isJoin bool) {
	c.Set(ContextKeyIsJoinCircle, isJoin)
}

func GetIsJoinCircleFromContext(c *gin.Context) bool {
	return c.GetBool(ContextKeyIsJoinCircle)
}

func SetIsCircleOwnerToContext(c *gin.Context, isOwner bool) {
	c.Set(ContextKeyIsCircleOwner, isOwner)
}

func GetIsCircleOwnerFromContext(c *gin.Context) bool {
	return c.GetBool(ContextKeyIsCircleOwner)
}

func SetAdminUidToContext(c *gin.Context, uid int) {
	c.Set(ContextKeyAdminUid, uid)
}

func GetAdminUidFromContext(c *gin.Context) int {
	return c.GetInt(ContextKeyAdminUid)
}

func SetAdminNameToContext(c *gin.Context, name string) {
	c.Set(ContextKeyAdminName, name)
}

func GetAdminNameFromContext(c *gin.Context) string {
	return c.GetString(ContextKeyAdminName)
}
