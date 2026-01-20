package mysql

import "bishe/internal/app/knowledge_sharing/model"

// create文章
func CreateEssay(newEssay *model.Essay) (err error) {
	err = DB.Model(&model.Essay{}).Create(newEssay).Error
	return
}
