package mysql

import "bishe/internal/app/knowledge_sharing/model"

// create圈子
func CreateCircle(newCircle *model.Circle) (err error) {
	err = DB.Model(&model.Circle{}).Create(newCircle).Error
	return
}

// get 付费圈子
func GetCircleAllCharge() (circles []model.Circle, err error) {
	err = DB.Model(&model.Circle{}).Where("price>?", 0).Find(&circles).Error
	if err != nil {
		return
	}

	return
}

// get 免费圈子
func GetCricleAllFree() (circles []model.Circle, err error) {
	err = DB.Model(&model.Circle{}).Where("price=?", 0).Find(&circles).Error
	if err != nil {
		return
	}

	return
}

// get all圈子
func GetCircleAllSortByJoinNum() (circles []model.Circle, err error) {
	err = DB.Model(&model.Circle{}).Order("join_num ASC").Find(&circles).Error
	if err != nil {
		return
	}

	return
}
