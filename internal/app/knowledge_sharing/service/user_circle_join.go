package service

import (
	"bishe/internal/app/knowledge_sharing/dao/mysql"
	"bishe/internal/app/knowledge_sharing/model"
)

// create 用户加入圈子
func CreateUserCircleJoin(newUserCircle *model.UserCircleJoin) (err error) {
	return mysql.CreateUserCircleJoin(newUserCircle)
}
