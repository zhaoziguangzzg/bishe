package service

import (
	"bishe/internal/app/knowledge_sharing/dao/mysql"
	"bishe/internal/app/knowledge_sharing/model"
)

// create文章
func CreateEssay(newEssay *model.Essay) (err error) {
	return mysql.CreateEssay(newEssay)
}

// get用户全部文章
func GetAllEssayByUid(uid int, page int, pagesize int) (essays []model.Essay, err error) {
	return mysql.GetAllEssayByUid(uid, page, pagesize)
}

// get圈子中的文章
func GetAllEssayByCid(cid int, page int, pagesize int) (essays []model.Essay, err error) {
	return mysql.GetAllEssayByCid(cid, page, pagesize)
}

// 根据eid获取文章
func GetEssayByEid(eid int) (essay *model.Essay, err error) {
	return mysql.GetEssayByEid(eid)
}

// 根据eid更新文章信息
func UpdateEssayByEid(eid int, title string, content string) (int64, error) {
	return mysql.UpdateEssayByEid(eid, title, content)
}

// 更新IsDeleted删除essay
func UpdateEssayIsDeleted(eid int) (int64, error) {
	return mysql.UpdateEssayIsDeleted(eid)
}
