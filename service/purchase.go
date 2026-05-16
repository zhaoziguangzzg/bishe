package service

import (
	"bishe/dao/mysql"
	"bishe/model"
	"time"
)

// 用户购买课程
func CreatePurchase(purchase *model.Purchase) (err error) {
	return mysql.CreatePurchase(purchase)
}

// 获取用户购买记录详情
func GetPurchaseById(uid int) (purchase *model.Purchase, err error) {
	return mysql.GetPurchaseById(uid)
}

// 获取用户购买课程
func GetPurchaseByUidCid(uid int, cid int) (purchases []model.Purchase, err error) {
	return mysql.GetPurchaseByUidCid(uid, cid)
}

// 获取购买课程
func GetUserPurchaseByUidCid(uid int, cid int) (purchase *model.Purchase, err error) {
	return mysql.GetUserPurchaseByUidCid(uid, cid)
}

// 获取用户全部购买课程记录
func GetAllPurchaseByUid(uid int) (purchases []model.Purchase, err error) {
	return mysql.GetAllPurchaseByUid(uid)
}

// 获取用户购买课程记录
func GetPurchaseByUid(uid int, status int) (purchases []model.Purchase, err error) {
	return mysql.GetPurchaseByUid(uid, status)
}

// 根据状态时间获取订单
func GetPurchaseByStatusTime(status int, t time.Time, limit int) (purchases []model.Purchase, err error) {
	return mysql.GetPurchaseByStatusTime(status, t, limit)
}

// 更新用户购买记录状态
func UpdatePurchaseStatusById(id int, status int, newStatus int) (int64, error) {
	return mysql.UpdatePurchaseStatusById(id, status, newStatus)
}

// 更新课程购买人数
func IncrCourseJoinNumByCid(cid int) (int64, error) {
	return mysql.IncrCourseJoinNumByCid(cid)
}
