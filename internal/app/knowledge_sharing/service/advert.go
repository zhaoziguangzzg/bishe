package service

import (
	"bishe/internal/app/knowledge_sharing/dao/mysql"
	"bishe/internal/app/knowledge_sharing/model"
	"time"
)

// create广告
func CreateAdvert(advert *model.Advert) (err error) {
	return mysql.CreateAdvert(advert)
}

// 获取全部广告
func GetAllAdvertByTime(ctime time.Time, position int, page int, pagesize int) (adverts []model.Advert, err error) {
	return mysql.GetAllAdvertByTime(ctime, position, page, pagesize)
}

// 获取全部广告
func GetAllAdvert(page int, pagesize int) (adverts []model.Advert, err error) {
	return mysql.GetAllAdvert(page, pagesize)
}

// 根据id获取广告
func GetAdvertById(id int) (advert *model.Advert, err error) {
	return mysql.GetAdvertById(id)
}

// 根据id更新广告
func UpdateAdvertById(id int, advert map[string]interface{}) (int64, error) {
	return mysql.UpdateAdvertById(id, advert)
}

// 根据id删除广告
func UpdateAdvertIsDeleted(id int) (int64, error) {
	return mysql.UpdateAdvertIsDeleted(id)
}

// 检查position
func CheckPosition(position int) (b bool) {
	b = true
	switch position {
	case model.ADVERT_POSITION_CIRCLE_INDEX:
	case model.ADVERT_POSITION_COURSE_INDEX:
	case model.ADVERT_POSITION_USER_PROFILE:
	case model.ADVERT_POSITION_INDEX:
	default:
		return false
	}
	return
}
