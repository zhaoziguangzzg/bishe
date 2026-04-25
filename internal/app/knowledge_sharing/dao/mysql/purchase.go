package mysql

import (
	"bishe/internal/app/knowledge_sharing/model"
	"time"

	"gorm.io/gorm"
)

// 用户购买课程
func CreatePurchase(purchase *model.Purchase) (err error) {
	err = DB.Model(&model.Purchase{}).Create(purchase).Error
	return
}

// 获取用户购买记录详情
func GetPurchaseById(id int) (purchase *model.Purchase, err error) {
	purchase = new(model.Purchase)
	err = DB.Model(&model.Purchase{}).Where("id=?", id).First(&purchase).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound { //没查到数据返回空
			return nil, nil
		}

		return nil, err
	}

	return purchase, nil
}

// 获取用户购买课程
func GetPurchaseByUidCid(uid int, cid int) (purchases []model.Purchase, err error) {
	err = DB.Model(&model.Purchase{}).
		//TODO in status
		Where("user_id=? and course_id=? and purchase_status =?", uid, cid, model.PURCHASE_STATUS_PAID).
		Or("user_id=? and course_id=? and purchase_status =?", uid, cid, model.PURCHASE_STATUS_UNPAID).
		Find(&purchases).Error

	if err != nil {
		return
	}

	return

}

// 获取购买课程
func GetUserPurchaseByUidCid(uid int, cid int) (purchase *model.Purchase, err error) {
	purchase = new(model.Purchase)
	err = DB.Model(&model.Purchase{}).
		Where("user_id=? and course_id=? and purchase_status=?", uid, cid, model.PURCHASE_STATUS_PAID).
		First(&purchase).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound { //没查到数据返回空
			return nil, nil
		}

		return nil, err
	}

	return purchase, nil

}

// 获取用户全部购买课程记录
func GetAllPurchaseByUid(uid int) (purchases []model.Purchase, err error) {
	err = DB.Model(&model.Purchase{}).Where("user_id = ?", uid).Find(&purchases).Error
	if err != nil {
		return
	}

	return
}

// 获取用户购买课程记录
func GetPurchaseByUid(uid int, status int) (purchases []model.Purchase, err error) {
	err = DB.Model(&model.Purchase{}).Where("user_id = ? and purchase_status=?", uid, status).Find(&purchases).Error
	if err != nil {
		return
	}

	return
}

// 根据状态时间获取订单
func GetPurchaseByStatusTime(status int, t time.Time, limit int) (purchases []model.Purchase, err error) {
	err = DB.Model(&model.Purchase{}).
		Where("purchase_status=? and create_at <? ", status, t).
		Order("id ASC").
		Limit(limit).
		Find(&purchases).Error
	if err != nil {
		return
	}

	return
}

// 更新用户购买记录状态
func UpdatePurchaseStatusById(id int, status int, newStatus int) (int64, error) {
	result := DB.Model(&model.Purchase{}).
		Where("id=? and purchase_status=?", id, status).
		Update("purchase_status", newStatus)
	return result.RowsAffected, result.Error
}

// 增加课程购买人数
func IncrCourseJoinNumByCid(cid int) (int64, error) {
	result := DB.Model(&model.Course{}).Where("id=?", cid).
		UpdateColumn("join_num", gorm.Expr("join_num + ?", 1))
	return result.RowsAffected, result.Error
}
