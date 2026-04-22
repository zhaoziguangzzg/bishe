package service

import (
	"bishe/internal/app/knowledge_sharing/dao/mysql"
	"bishe/internal/app/knowledge_sharing/model"
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

// 获取用户全部购买课程记录
func GetAllPurchaseByUid(uid int) (purchases []model.Purchase, err error) {
	return mysql.GetAllPurchaseByUid(uid)
}

// 获取用户购买课程记录
func GetPurchaseByUid(uid int, status int) (purchases []model.Purchase, err error) {
	return mysql.GetPurchaseByUid(uid, status)
}

// 更新用户购买记录状态
func UpdatePurchaseStatusById(id int, status int) (int64, error) {
	return mysql.UpdatePurchaseStatusById(id, status)
}

// 更新课程购买人数
func IncrCourseJoinNumByCid(cid int) (int64, error) {
	return mysql.IncrCourseJoinNumByCid(cid)
}
