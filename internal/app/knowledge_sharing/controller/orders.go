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
	uid, _ := service.GetUserFromCookie(c)
	if uid == 0 {
		MakeApiResponseError(c, CODE_USER_NOT_LOGIN)
		return
	}

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
		StartTime:   &createTime,
		EndTime:     &createTime,
		CreateAt:    &createTime,
		UpdateAt:    &createTime,
		OrderStatus: model.ORDER_STATUS_WAIT,
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
	uid, _ := service.GetUserFromCookie(c)
	if uid == 0 {
		MakeApiResponseError(c, CODE_USER_NOT_LOGIN)
		return
	}

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
	uid, _ := service.GetUserFromCookie(c)
	if uid == 0 {
		MakeApiResponseError(c, CODE_USER_NOT_LOGIN)
		return
	}

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

	orders, err := service.GetUserOrdersByUidCid(uid, cid)
	if err != nil {
		service.Logger.Error("GetOrdersById", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if orders == nil {
		MakeApiResponseError(c, CODE_ORDERS_NOT_EXIST)
		return
	}

	ordersIdStr := c.PostForm("orders_id")
	if ordersIdStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	ordersId, err := strconv.Atoi(ordersIdStr)
	if err != nil {
		MakeApiResponseErrorParams(c)
		return
	}

	var startTime time.Time
	var endTime time.Time

	LastEndTime := *orders.EndTime
	nowTime := time.Now()

	//开始时间大于当前时间
	// a after b    a再b后
	if nowTime.After(LastEndTime) {
		startTime = nowTime
	} else {
		startTime = LastEndTime
	}

	endTime = startTime.AddDate(1, 0, 0)

	//根据id更新支付
	affectRows, err := service.UpdateOrderById(ordersId, startTime, endTime)
	if err != nil || affectRows == 0 {
		service.Logger.Error("UpdateOrderById err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	MakeApiResponseSuccessDefault(c)
}
