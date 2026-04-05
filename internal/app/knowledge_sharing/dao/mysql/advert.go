package mysql

import (
	"bishe/internal/app/knowledge_sharing/model"
	"time"

	"gorm.io/gorm"
)

// create广告
func CreateAdvert(advert *model.Advert) (err error) {
	err = DB.Model(&model.Advert{}).Create(advert).Error
	return
}

// 获取全部广告
func GetAllAdvertByTime(ctime time.Time, position string, page int, pagesize int) (adverts []model.Advert, err error) {
	offset := (page - 1) * pagesize

	err = DB.Model(&model.Advert{}).Where("start_time<? and end_time>? and position=? and is_deleted=?", ctime, ctime, position, model.IS_DELETED_NO).
		Order("id DESC").Offset(offset).Limit(pagesize).Find(&adverts).Error
	if err != nil {
		return
	}

	return
}

// 根据id获取广告
func GetAdvertById(id int) (advert *model.Advert, err error) {
	advert = new(model.Advert)
	err = DB.Model(&model.Advert{}).Where("id=?", id).First(&advert).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound { //没查到数据返回空
			return nil, nil
		}

		return nil, err
	}

	return advert, nil
}

// 根据id更新广告
func UpdateAdvertById(id int, position string, content string, startTime time.Time, endTime time.Time) (int64, error) {
	advert := model.Advert{
		Position:  position,
		Content:   content,
		StartTime: &startTime,
		EndTime:   &endTime,
	}

	result := DB.Model(&model.Advert{}).Where("id=?", id).Updates(advert)
	return result.RowsAffected, result.Error
}

// 根据id删除广告
func UpdateAdvertIsDeleted(id int) (int64, error) {
	result := DB.Model(&model.Advert{}).Where("id=?", id).Update("is_deleted", model.IS_DELETED_YES)
	return result.RowsAffected, result.Error
}
