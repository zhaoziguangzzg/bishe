package service

import (
	"bishe/dao/mysql"
	"bishe/model"
)

// create支付
func CreateOrders(orders *model.Orders) (err error) {
	return mysql.CreateOrders(orders)
}

// 获取用户全部支付
func GetUserAllOrdersByUid(uid int, page int, pagesize int) (orderss []model.Orders, err error) {
	return mysql.GetUserAllOrdersByUid(uid, page, pagesize)
}

// 根据id获取支付
func GetOrdersById(ordersId int) (orders *model.Orders, err error) {
	return mysql.GetOrdersById(ordersId)
}

// 根据uid，cid获取所有订单
func GetUserOrdersByUidCid(uid int, cid int) (orderss []model.Orders, err error) {
	return mysql.GetUserOrdersByUidCid(uid, cid)
}

// 根据id更新支付
func UpdateOrderStatusById(id int, status int, newStatus int) (int64, error) {
	return mysql.UpdateOrderStatusById(id, status, newStatus)
}
