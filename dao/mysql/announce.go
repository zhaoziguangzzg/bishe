package mysql

import (
	"bishe/model"
	"time"

	"gorm.io/gorm"
)

// create公告
func CreateAnnounce(announce *model.Announce) (err error) {
	err = DB.Model(&model.Announce{}).Create(announce).Error
	return
}

// 获取全部公告
func GetAllAnnounceByTime(ctime time.Time, page int, pagesize int) (announces []model.Announce, err error) {
	offset := (page - 1) * pagesize

	err = DB.Model(&model.Announce{}).Where("start_time<? and end_time>? and is_deleted=?", ctime, ctime, model.IS_DELETED_NO).
		Order("id DESC").Offset(offset).Limit(pagesize).Find(&announces).Error
	if err != nil {
		return
	}

	return
}

// 获取全部公告
func GetAllAnnounce(page int, pagesize int) (announces []model.Announce, err error) {
	offset := (page - 1) * pagesize

	err = DB.Model(&model.Announce{}).
		Where("is_deleted=?", model.IS_DELETED_NO).
		Order("id DESC").
		Offset(offset).
		Limit(pagesize).
		Find(&announces).Error
	if err != nil {
		return
	}

	return
}

// 根据id获取公告
func GetAnnounceById(id int) (announce *model.Announce, err error) {
	announce = new(model.Announce)
	err = DB.Model(&model.Announce{}).Where("id=?", id).First(&announce).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound { //没查到数据返回空
			return nil, nil
		}

		return nil, err
	}

	return announce, nil
}

// 根据id更新公告
func UpdateAnnounceById(id int, announce map[string]interface{}) (int64, error) {

	result := DB.Model(&model.Announce{}).Where("id=?", id).Updates(announce)
	return result.RowsAffected, result.Error
}

// 根据id删除公告
func UpdateAnnounceIsDeleted(id int) (int64, error) {
	result := DB.Model(&model.Announce{}).Where("id=?", id).Update("is_deleted", model.IS_DELETED_YES)
	return result.RowsAffected, result.Error
}
