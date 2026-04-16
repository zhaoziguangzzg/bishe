package service

import (
	"bishe/internal/app/knowledge_sharing/dao/mysql"
	"bishe/internal/app/knowledge_sharing/model"
	"time"
)

// create公告
func CreateAnnounce(announce *model.Announce) (err error) {
	return mysql.CreateAnnounce(announce)
}

// 获取全部公告
func GetAllAnnounceByTime(ctime time.Time, page int, pagesize int) (announces []model.Announce, err error) {
	return mysql.GetAllAnnounceByTime(ctime, page, pagesize)
}

// 获取全部公告
func GetAllAnnounce(page int, pagesize int) (announces []model.Announce, err error) {
	return mysql.GetAllAnnounce(page, pagesize)
}

// 根据id获取公告
func GetAnnounceById(id int) (announce *model.Announce, err error) {
	return mysql.GetAnnounceById(id)
}

// 根据id更新公告
func UpdateAnnounceById(id int, announce map[string]interface{}) (int64, error) {
	return mysql.UpdateAnnounceById(id, announce)
}

// 根据id删除公告
func UpdateAnnounceIsDeleted(id int) (int64, error) {
	return mysql.UpdateAnnounceIsDeleted(id)
}
