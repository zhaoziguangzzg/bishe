package service

import (
	"bishe/internal/app/knowledge_sharing/dao/mysql"
	"bishe/internal/app/knowledge_sharing/model"
)

// create文章
func CreateEssay(newEssay *model.Essay) (err error) {
	return mysql.CreateEssay(newEssay)
}
