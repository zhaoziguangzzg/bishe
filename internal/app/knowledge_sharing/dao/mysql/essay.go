package mysql

import (
	"bishe/internal/app/knowledge_sharing/model"

	"gorm.io/gorm"
)

// create文章
func CreateEssay(newEssay *model.Essay) (err error) {
	err = DB.Model(&model.Essay{}).Create(newEssay).Error
	return
}

// get用户全部文章
func GetAllEssayByUid(uid int, page int, pagesize int) (essays []model.Essay, err error) {
	offset := (page - 1) * pagesize

	err = DB.Model(&model.Essay{}).Where("author_id=?", uid).Order("id ASC").
		Offset(offset).Limit(pagesize).Find(&essays).Error
	if err != nil {
		return
	}

	return
}

// 根据eid获取文章
func GetEssayByEid(eid int) (essay *model.Essay, err error) {
	essay = new(model.Essay)
	err = DB.Model(&model.Essay{}).Where("id=? and essay_status=?", eid, model.ESSAY_NOT_DELETED).First(&essay).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound { //没查到数据返回空
			return nil, nil
		}

		return nil, err
	}

	return essay, nil
}

// 根据eid更新文章信息
func UpdateEssayByEid(eid int, title string, content string) (int64, error) {
	essay := model.Essay{
		Title:   title,
		Content: content,
	}

	result := DB.Model(&model.Essay{}).Where("id=?", eid).Updates(essay)
	return result.RowsAffected, result.Error
}
