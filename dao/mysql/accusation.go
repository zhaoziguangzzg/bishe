package mysql

import (
	"bishe/model"

	"gorm.io/gorm"
)

// create 用户举报
func CreateUserAccusation(newAccusation *model.Accusation) (err error) {
	err = DB.Model(&model.Accusation{}).Create(newAccusation).Error
	return
}

// 根据uid,eid获取举报文章
func GetUserAccusationEssay(uid int, eid int) (accusation *model.Accusation, err error) {
	accusation = new(model.Accusation)
	err = DB.Model(&model.Accusation{}).
		Where("user_id=? and essay_id=? and and is_deleted=?", uid, eid, model.ACCUSATION_NOT_DELETED).
		First(&accusation).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound { //没查到数据返回空
			return nil, nil
		}

		return nil, err
	}

	return accusation, nil
}

// 获取全部未处理举报
func GetAllAccusationEssay(page int, pagesize int) (accusations []model.Accusation, err error) {
	offset := (page - 1) * pagesize

	err = DB.Model(&model.Accusation{}).
		Where("accusation_status=? and is_deleted=?", model.ACCUSATION_STATUS_WAIT, model.ACCUSATION_NOT_DELETED).
		Order("accusation_time DESC").Offset(offset).Limit(pagesize).Find(&accusations).Error
	return
}

// 获取文章举报内容
func GetAccusationByAid(aid int) (accusation *model.Accusation, err error) {
	accusation = new(model.Accusation)
	err = DB.Model(&model.Accusation{}).Where("id=? and is_deleted=?", aid, model.ACCUSATION_NOT_DELETED).First(&accusation).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound { //没查到数据返回空
			return nil, nil
		}

		return nil, err
	}

	return accusation, nil
}

// 更新举报信息为无违规
func UpdateAccusationNormalByAid(aid int) (int64, error) {
	result := DB.Model(&model.Accusation{}).Where("id=?", aid).Update("accusation_status", model.ACCUSATION_STATUS_NORMAL)
	return result.RowsAffected, result.Error
}

// 更新举报信息为有违规
func UpdateAccusationViolateByAid(aid int) (int64, error) {
	result := DB.Model(&model.Accusation{}).Where("id=?", aid).Update("accusation_status", model.ACCUSATION_STATUS_VIOLATE)
	return result.RowsAffected, result.Error
}
