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
func GetAllAdvertByTime(ctime time.Time, position string, page int, pagesize int) (adverts []model.Advert, err error) {
	return mysql.GetAllAdvertByTime(ctime, position, page, pagesize)
}

// 根据id获取广告
func GetAdvertById(id int) (advert *model.Advert, err error) {
	return mysql.GetAdvertById(id)
}

// 根据id更新广告
func UpdateAdvertById(id int, position string, content string, startTime time.Time, endTime time.Time) (int64, error) {
	return mysql.UpdateAdvertById(id, position, content, startTime, endTime)
}

// 根据id删除广告
func UpdateAdvertIsDeleted(id int) (int64, error) {
	return mysql.UpdateAdvertIsDeleted(id)
}
