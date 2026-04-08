package mysql

import (
	"bishe/internal/app/knowledge_sharing/model"

	"gorm.io/gorm"
)

// create支付
func CreateOrders(orders *model.Orders) (err error) {
	err = DB.Model(&model.Orders{}).Create(orders).Error
	return
}

// 获取用户全部支付
func GetUserAllOrdersByUid(uid int, page int, pagesize int) (orderss []model.Orders, err error) {
	offset := (page - 1) * pagesize
	err = DB.Model(&model.Orders{}).Where("uid=? ", uid).
		Order("id DESC").Offset(offset).Limit(pagesize).Find(&orderss).Error
	if err != nil {
		return
	}

	return
}

// 根据id获取支付
func GetOrdersById(ordersId int) (orders *model.Orders, err error) {
	orders = new(model.Orders)
	err = DB.Model(&model.Orders{}).Where("id=?", ordersId).First(&orders).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound { //没查到数据返回空
			return nil, nil
		}

		return nil, err
	}

	return orders, nil
}

// 根据uid，cid获取最后支付
func GetUserOrdersByUidCid(uid int, cid int) (orders *model.Orders, err error) {
	err = DB.Model(&model.Orders{}).Where("uid=? and cid=? ", uid, cid).
		Where("order_status=? or order_status=?", model.ORDER_STATUS_PAID, model.ORDER_STATUS_END).
		Order("id DESC").Limit(1).Find(&orders).Error
	if err != nil {
		return
	}

	return
}

// 根据id更新支付
func UpdateOrderById(id int) (int64, error) {

	result := DB.Model(&model.Orders{}).Where("id=?", id).Update("order_status", model.ORDER_STATUS_PAID)
	return result.RowsAffected, result.Error
}
