package service

import (
	"bishe/internal/app/knowledge_sharing/dao/mysql"
	"bishe/internal/app/knowledge_sharing/model"
)

// create圈子
func CreateCircle(newCircle *model.Circle) (err error) {
	return mysql.CreateCircle(newCircle)
}

// get 付费圈子
func GetCircleAllCharge() (circles []model.Circle, err error) {
	return mysql.GetCircleAllCharge()
}

// get 免费圈子
func GetCricleAllFree() (circles []model.Circle, err error) {
	return mysql.GetCricleAllFree()
}

// get all圈子
func GetCircleAllSortByJoinNum() (circles []model.Circle, err error) {
	return mysql.GetCircleAllSortByJoinNum()
}
