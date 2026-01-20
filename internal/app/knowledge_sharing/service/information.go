package service

import (
	"bishe/internal/app/knowledge_sharing/dao/mysql"
	"bishe/internal/app/knowledge_sharing/model"
)

func UserAddInformation(information *model.Information) (err error) {
	return mysql.UserAddInformation(information)
}
