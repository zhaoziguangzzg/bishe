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

	err = DB.Model(&model.Essay{}).Where("author_id=? and is_deleted=?", uid, model.ESSAY_NOT_DELETED).
		Order("id DESC").Offset(offset).Limit(pagesize).Find(&essays).Error
	if err != nil {
		return
	}

	return
}

// get圈子中的文章
func GetAllEssayByCid(cid int, page int, pagesize int) (essays []model.Essay, err error) {
	offset := (page - 1) * pagesize

	err = DB.Model(&model.Essay{}).Where("circle_id=? and is_deleted=?", cid, model.ESSAY_NOT_DELETED).
		Order("id DESC").Offset(offset).Limit(pagesize).Find(&essays).Error
	if err != nil {
		return
	}

	return
}

// 根据eid获取文章
func GetEssayByEid(eid int) (essay *model.Essay, err error) {
	essay = new(model.Essay)
	err = DB.Model(&model.Essay{}).Where("id=? and is_deleted=?", eid, model.ESSAY_NOT_DELETED).First(&essay).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound { //没查到数据返回空
			return nil, nil
		}

		return nil, err
	}

	return essay, nil
}

// 根据eids获取文章
func GetEssayByEids(eids []int) (newEssays []model.Essay, err error) {
	//TODO in取的数据为乱序，需要排序
	var essays []model.Essay
	err = DB.Where("id IN (?)", eids).Find(&essays).Error
	if err != nil {
		return
	}

	//TODO 按eids 排序
	for {

	}

	return
}

// 根据title获取文章
func GetEssayByTitle(title string, cid int) (essay *model.Essay, err error) {
	essay = new(model.Essay)
	err = DB.Model(&model.Essay{}).Where("title=? and cid=? and is_deleted=?", title, cid, model.ESSAY_NOT_DELETED).First(&essay).Error
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

// 更新IsDeleted删除essay
func UpdateEssayIsDeleted(eid int) (int64, error) {
	result := DB.Model(&model.Essay{}).Where("id=?", eid).Update("is_deleted", model.ESSAY_IS_DELETED)
	return result.RowsAffected, result.Error
}
