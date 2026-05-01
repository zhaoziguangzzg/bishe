package controller

import (
	"bishe/internal/app/knowledge_sharing/model"
	"bishe/internal/app/knowledge_sharing/service"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 添加支付
func AddOrdersHandler(c *gin.Context) {
	uid := service.GetUidFromContext(c)

	cidStr := c.PostForm("cid")
	if cidStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	cid, err := strconv.Atoi(cidStr)
	if err != nil {
		MakeApiResponseErrorParams(c)
		return
	}

	levelStr := c.PostForm("level")
	if levelStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	level, err := strconv.Atoi(levelStr)
	if err != nil {
		MakeApiResponseErrorParams(c)
		return
	}

	priceStr := c.PostForm("price")
	if priceStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	price, err := strconv.Atoi(priceStr)
	if err != nil {
		MakeApiResponseErrorParams(c)
		return
	}

	var discount float64
	discount = 0.01 * float64(level*price)

	createTime := time.Now()

	// 构造支付
	orders := &model.Orders{ //其中包含自动生成的id
		Uid:         uid,
		Cid:         cid,
		Price:       price,
		Discount:    discount,
		CreateAt:    &createTime,
		UpdateAt:    &createTime,
		OrderStatus: model.ORDER_STATUS_UNPAID,
	}

	//添加支付
	err = service.CreateOrders(orders)
	if err != nil {
		service.Logger.Error("CreateOrders err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	MakeApiResponseSuccessDefault(c)
}

// 获取用户全部支付列表
func GetUserAllOrdersHandler(c *gin.Context) {
	uid := service.GetUidFromContext(c)

	pageStr := c.Query("page")
	page := GetPage(pageStr)

	pagesize := 10

	//获取用户全部支付
	orderss, err := service.GetUserAllOrdersByUid(uid, page, pagesize)
	if err != nil {
		service.Logger.Error("GetUserAllOrdersByUid", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if len(orderss) == 0 {
		orderss = make([]model.Orders, 0)
	}

	MakeApiResponseSuccess(c, map[string]interface{}{
		"orderss": orderss,
	})
}

// 获取某支付信息
func GetOrdersHandler(c *gin.Context) {
	//获取支付id
	ordersIdStr := c.Query("orders_id")
	if ordersIdStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	ordersId, err := strconv.Atoi(ordersIdStr)
	if err != nil {
		MakeApiResponseErrorParams(c)
		return
	}

	//根据id获取支付
	orders, err := service.GetOrdersById(ordersId)
	if err != nil {
		service.Logger.Error("GetOrdersById", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if orders == nil {
		MakeApiResponseError(c, CODE_ORDERS_NOT_EXIST)
		return
	}

	data := map[string]interface{}{
		"orders": orders,
	}

	MakeApiResponseSuccess(c, data)
}

// 用户支付更新
func UpdateUserOrdersHandler(c *gin.Context) {
	uid := service.GetUidFromContext(c)

	cidStr := c.PostForm("cid")
	if cidStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	cid, err := strconv.Atoi(cidStr)
	if err != nil {
		MakeApiResponseErrorParams(c)
		return
	}

	idStr := c.PostForm("orders_id")
	if idStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		MakeApiResponseErrorParams(c)
		return
	}

	orders, err := service.GetOrdersById(id)
	if err != nil {
		service.Logger.Error("GetOrdersById err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	status := orders.OrderStatus

	statusNew, err := service.MakeOrderStatus(status, model.ORDER_ACTION_PAY)
	if err != nil {
		service.Logger.Error("MakeOrderStatus err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	//根据id更新支付
	affectRows, err := service.UpdateOrderStatusById(id, status, statusNew)
	if err != nil || affectRows == 0 {
		service.Logger.Error("UpdateOrderById err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	//获取用户圈子
	join, err := service.GetUserCircleJoinByUidCid(uid, cid)
	if err != nil {
		service.Logger.Error("GetUserCircleJoinByUidCid", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if join == nil {
		MakeApiResponseError(c, CODE_USER_NOT_JOIN_CIRCLE)
		return
	}

	var startTime time.Time
	var endTime time.Time

	LastEndTime := *join.EndTime
	nowTime := time.Now()

	//开始时间大于当前时间
	// a after b    a再b后
	if nowTime.After(LastEndTime) {
		startTime = nowTime
	} else {
		startTime = LastEndTime
	}

	endTime = startTime.AddDate(1, 0, 0)

	//更新join starttime
	affectRows, err = service.UpdateUserCircleJoinTimeByUidCid(uid, cid, startTime, endTime)
	if err != nil || affectRows == 0 {
		service.Logger.Error("UpdateUserCircleJoinTimeByUidCid err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	MakeApiResponseSuccessDefault(c)
}

// 获取圈子待续费
func GetUserOrdersCircleHandler(c *gin.Context) {
	uid := service.GetUidFromContext(c)

	cidStr := c.Query("cid")
	if cidStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	cid, err := strconv.Atoi(cidStr)
	if err != nil {
		MakeApiResponseErrorParams(c)
		return
	}

	//付费圈子
	circle, err := service.GetCircleByCid(cid)
	if err != nil {
		service.Logger.Error("GetCircleByCid", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if circle == nil {
		MakeApiResponseError(c, CODE_CIRCLE_NOT_EXIST)
		return
	}

	if circle.Price == 0 {
		MakeApiResponseError(c, CODE_CIRCLE_FREE)
		return
	}

	userCircleJoin, err := service.GetUserJoinCircleByUidCid(uid, cid)
	if err != nil {
		service.Logger.Error("GetUserJoinCircleByUidCid", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if userCircleJoin == nil {
		MakeApiResponseError(c, CODE_USER_NOT_JOIN_CIRCLE)
		return
	}

	nowTime := time.Now()
	var endTime time.Time
	endTime = *userCircleJoin.EndTime
	eTime := endTime.AddDate(0, -1, 0)
	need := false

	if nowTime.After(eTime) {
		need = true
	}

	data := map[string]bool{
		"need": need,
	}

	MakeApiResponseSuccess(c, data)

}
