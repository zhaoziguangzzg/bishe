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
func GetEssayByEids(eids []int) (essays []model.Essay, err error) {

	//in取的数据为乱序，需要排序
	err = DB.Model(&model.Essay{}).Where("id IN (?)", eids).Find(&essays).Error
	if err != nil {
		return
	}

	return
}

// 根据title获取文章
func GetEssayByTitle(title string, cid int) (essay *model.Essay, err error) {
	essay = new(model.Essay)
	err = DB.Model(&model.Essay{}).Where("title=? and circle_id=? and is_deleted=?", title, cid, model.ESSAY_NOT_DELETED).
		First(&essay).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound { //没查到数据返回空
			return nil, nil
		}

		return nil, err
	}

	return essay, nil
}

// 根据title关键词like获取文章
func GetEssayByLikeTitle(title string, cid int, page int, pagesize int) (essays []model.Essay, err error) {
	offset := (page - 1) * pagesize

	err = DB.Model(&model.Essay{}).
		Where("title like ? and is_deleted=?", "%"+title+"%", model.IS_DELETED_NO).
		Order("id DESC").Offset(offset).Limit(pagesize).Find(&essays).Error
	if err != nil {
		return
	}

	return
}

// 根据eid更新文章信息
func UpdateEssayByEid(eid int, updateMap map[string]interface{}) (int64, error) {
	result := DB.Model(&model.Essay{}).Where("id=?", eid).Updates(updateMap)
	return result.RowsAffected, result.Error
}

// 更新IsDeleted删除essay
func UpdateEssayIsDeleted(eid int) (int64, error) {
	result := DB.Model(&model.Essay{}).Where("id=?", eid).Update("is_deleted", model.ESSAY_IS_DELETED)
	return result.RowsAffected, result.Error
}

// 更新文章周刊
func UpdateEssayWeekly(eid int, isWeekly int) (int64, error) {

	result := DB.Model(&model.Essay{}).Where("id=?", eid).Update("is_weekly", isWeekly)
	return result.RowsAffected, result.Error
}

// 获取文章周刊
func GetEssayWeeklyList(cid int, page int, pagesize int) (essays []model.Essay, err error) {
	offset := (page - 1) * pagesize

	err = DB.Model(&model.Essay{}).
		Where("circle_id=? and is_deleted=? and is_weekly=?", cid, model.ESSAY_NOT_DELETED, model.ESSAY_IS_ESSENCE).
		Order("id DESC").Offset(offset).Limit(pagesize).Find(&essays).Error
	if err != nil {
		return
	}

	return
}

// update essay essence
func UpdateEssayEssence(eid int, isEssence int) (int64, error) {
	result := DB.Model(&model.Essay{}).Where("id=?", eid).Update("is_essence", isEssence)
	return result.RowsAffected, result.Error
}

// 获取文章精粹
func GetEssayEssenceList(cid int, page int, pagesize int) (essays []model.Essay, err error) {
	offset := (page - 1) * pagesize

	err = DB.Model(&model.Essay{}).
		Where("circle_id=? and is_deleted=? and is_essence=?", cid, model.ESSAY_NOT_DELETED, model.ESSAY_IS_ESSENCE).
		Order("id DESC").Offset(offset).Limit(pagesize).Find(&essays).Error
	if err != nil {
		return
	}

	return
}
