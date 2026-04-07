package service

import (
	"bishe/internal/app/knowledge_sharing/dao/mysql"
	"bishe/internal/app/knowledge_sharing/model"
)

// create圈子
func CreateCircle(newCircle *model.Circle) (err error) {
	return mysql.CreateCircle(newCircle)
}

// 根据cid获取圈子
func GetCircleByCid(cid int) (circle *model.Circle, err error) {
	return mysql.GetCircleByCid(cid)
}

// 根据title获取圈子
func GetCircleByTitle(title string) (circle *model.Circle, err error) {
	return mysql.GetCircleByTitle(title)
}

// 通过like title关键词获取圈子
func GetCircleByLikeTitle(title string, page int, pagesize int) (circles []model.Circle, err error) {
	return mysql.GetCircleByLikeTitle(title, page, pagesize)
}

// get 付费圈子
func GetCircleAllChargeOrderByJoinNum(page int, pagesize int) (circles []model.Circle, err error) {
	return mysql.GetCircleAllChargeOrderByJoinNum(page, pagesize)
}

// get 免费圈子
func GetCricleAllFreeOrderByJoinNum(page int, pagesize int) (circles []model.Circle, err error) {
	return mysql.GetCricleAllFreeOrderByJoinNum(page, pagesize)
}

// get all圈子
func GetCircleAllByJoinNum(page int, pagesize int) (circles []model.Circle, err error) {
	return mysql.GetCircleAllByJoinNum(page, pagesize)
}

// get 用户创建的圈子
func GetUserCreateCircleByUid(uid int, page int, pagesize int) (circles []model.Circle, err error) {
	return mysql.GetUserCreateCircleByUid(uid, page, pagesize)
}

// get 用户加入的圈子
func GetUserJoinCircleListByUid(uid int, page int, pagesize int) (circles []model.Circle, err error) {
	return mysql.GetUserJoinCircleListByUid(uid, page, pagesize)
}

// 更新圈子join num 增加
func IncrUpdateCircleJoinNumByCid(cid int) (int64, error) {
	return mysql.IncrUpdateCircleJoinNumByCid(cid)
}

// 更新圈子join num 减少
func DecrrUpdateCircleJoinNumByCid(cid int) (int64, error) {
	return mysql.DecrrUpdateCircleJoinNumByCid(cid)
}

// 更新圈子信息
func UpdateCircleByCid(cid int, title string, price int, introduction string) (int64, error) {
	return mysql.UpdateCircleByCid(cid, title, price, introduction)
}

// 更新圈子状态
func UpdateCircleByTitle(title string) (int64, error) {
	return mysql.UpdateCircleByTitle(title)
}
