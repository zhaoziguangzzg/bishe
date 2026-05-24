package controller

import (
	"bishe/model"
	"bishe/service"
	"net"
	"net/url"
	"strconv"
	"strings"
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

	// 查询所有订单，遍历判断是否有待支付订单
	allOrders, err := service.GetUserOrdersByUidCid(uid, cid)
	if err != nil {
		service.Logger.Error("GetUserOrdersByUidCid err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	for _, order := range allOrders {
		if order.OrderStatus == model.ORDER_STATUS_UNPAID {
			MakeApiResponseError(c, CODE_HAS_UNPAY_ORDER)
			return
		}
	}

	lockKey := "user-add-order" + strconv.Itoa(uid) + "-" + strconv.Itoa(cid)
	lockValue, locked, err := service.Lock(c, lockKey, 5*time.Second)
	if err != nil {
		service.Logger.Error("Lock err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if !locked {
		MakeApiResponseError(c, CODE_LOCKED)
		return
	}

	defer service.Unlock(c, lockKey, lockValue)

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

	circle, err := service.GetCircleByCid(cid)
	if err != nil || circle == nil {
		service.Logger.Error("GetCircleByCid err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	payUrl := getPayUrl(c, orders.Id, cid, price, circle.Title)
	qrCodeUrl, err := service.QrcodeImgSave(payUrl, 200, service.FileTypeCircleImg, createTime)
	if err != nil {
		service.Logger.Error("QrcodeImgSave err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	MakeApiResponseSuccess(c, map[string]interface{}{
		"orders_id":   orders.Id,
		"pay_url":     payUrl,
		"qr_code_url": qrCodeUrl,
	})
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

// 获取已有订单的支付二维码
func GetOrdersQrcodeHandler(c *gin.Context) {
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

	if orders.OrderStatus != model.ORDER_STATUS_UNPAID {
		MakeApiResponseError(c, CODE_HAS_UNPAY_ORDER)
		return
	}

	circle, err := service.GetCircleByCid(orders.Cid)
	if err != nil || circle == nil {
		service.Logger.Error("GetCircleByCid err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	payUrl := getPayUrl(c, orders.Id, orders.Cid, orders.Price, circle.Title)
	qrCodeUrl, err := service.QrcodeImgSave(payUrl, 200, service.FileTypeCircleImg, *orders.CreateAt)
	if err != nil {
		service.Logger.Error("QrcodeImgSave err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	MakeApiResponseSuccess(c, map[string]interface{}{
		"orders_id":   orders.Id,
		"pay_url":     payUrl,
		"qr_code_url": qrCodeUrl,
	})
}

// 用户支付更新
func UpdateUserOrdersHandler(c *gin.Context) {
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

	cid := orders.Cid
	uid := orders.Uid

	// 获取圈子信息
	circle, err := service.GetCircleByCid(cid)
	if err != nil {
		service.Logger.Error("GetCircleByCid err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}
	if circle == nil {
		MakeApiResponseError(c, CODE_CIRCLE_NOT_EXIST)
		return
	}

	isFree := circle.Price == 0

	// 处理用户加入圈子
	join, err := service.GetUserCircleJoinByUidCid(uid, cid)
	if err != nil {
		service.Logger.Error("GetUserCircleJoinByUidCid", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if join != nil && join.NotJoinStatus == model.USER_CIRCLE_JOIN_STATUS_JOIN {
		// 已加入圈子，续费延长加入时间
		var startTime time.Time
		LastEndTime := *join.EndTime
		nowTime := time.Now()

		if nowTime.After(LastEndTime) {
			startTime = nowTime
		} else {
			startTime = LastEndTime
		}

		endTime := startTime.AddDate(1, 0, 0)

		affectRows, err = service.UpdateUserCircleJoinTimeByUidCid(uid, cid, startTime, endTime)
		if err != nil || affectRows == 0 {
			service.Logger.Error("UpdateUserCircleJoinTimeByUidCid err", zap.Error(err))
			MakeApiResponseErrorDefault(c)
			return
		}
	} else if join != nil && join.NotJoinStatus == model.USER_CIRCLE_JOIN_STATUS_NOT_JOIN {
		// 之前退出过，重新加入
		affectRows, err := service.UpdateUserCircleJoinStatusByJid(join.Id, model.USER_CIRCLE_JOIN_STATUS_JOIN)
		if affectRows == 0 || err != nil {
			service.Logger.Error("UpdateUserCircleJoinStatusByJid err", zap.Error(err))
			MakeApiResponseErrorDefault(c)
			return
		}

		affectRows, _, err = service.IncrUpdateCircleJoinNumByCid(c, cid, isFree)
		if affectRows == 0 || err != nil {
			service.Logger.Error("IncrUpdateCircleJoinNumByCid err", zap.Error(err))
			MakeApiResponseErrorDefault(c)
			return
		}
	} else {
		// 首次加入圈子
		nowTime := time.Now()
		endTime := nowTime.AddDate(1, 0, 0)
		joinId, err := service.CreateUserJoinCircleAndUpdateJoinNum(c, uid, cid, nowTime, endTime, isFree)
		if err != nil {
			service.Logger.Error("CreateUserJoinCircleAndUpdateJoinNum err", zap.Error(err))
			MakeApiResponseErrorDefault(c)
			return
		}

		if joinId == 0 {
			MakeApiResponseErrorDefault(c)
			return
		}

		// 发送加入通知
		noticeMsg := &model.NoticeMsg{
			Type:     model.NOTICE_TYPE_JOIN,
			Time:     nowTime.Unix(),
			JoinUid:  uid,
			CircleId: cid,
		}

		_, _, err = service.ProduceKafkaNoticeMessage(noticeMsg)
		if err != nil {
			service.Logger.Error("ProduceKafkaNoticeMessage err", zap.Error(err))
			err = nil
		}

		err = service.UserAddLevelScore(uid, cid, nowTime)
		if err != nil {
			service.Logger.Error("UserAddLevelScore err", zap.Error(err))
			MakeApiResponseErrorDefault(c)
			return
		}
	}

	MakeApiResponseSuccessDefault(c)
}

// 取消未支付订单
func CancelOrdersHandler(c *gin.Context) {
	uid := service.GetUidFromContext(c)

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

	if orders == nil {
		MakeApiResponseError(c, CODE_ORDERS_NOT_EXIST)
		return
	}

	if orders.Uid != uid {
		MakeApiResponseErrorParams(c)
		return
	}

	if orders.OrderStatus != model.ORDER_STATUS_UNPAID {
		MakeApiResponseErrorDefault(c)
		return
	}

	statusNew, err := service.MakeOrderStatus(orders.OrderStatus, model.ORDER_ACTION_CANCEL)
	if err != nil {
		service.Logger.Error("MakeOrderStatus err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	affectRows, err := service.UpdateOrderStatusById(id, orders.OrderStatus, statusNew)
	if err != nil || affectRows == 0 {
		service.Logger.Error("UpdateOrderStatusById err", zap.Error(err))
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

func getPayUrl(c *gin.Context, ordersId int, cid int, price int, title string) string {
	port := c.Request.Host
	if idx := strings.LastIndex(port, ":"); idx != -1 {
		port = port[idx+1:]
	} else {
		port = "8080"
	}

	ip := getLanIP()
	if ip == "" {
		ip = "127.0.0.1"
	}

	return "http://" + ip + ":" + port + "/page/orders/pay?orders_id=" + strconv.Itoa(ordersId) +
		"&cid=" + strconv.Itoa(cid) + "&price=" + strconv.Itoa(price) +
		"&title=" + url.QueryEscape(title)
}

func getLanIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			if ipnet.IP.IsPrivate() {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
