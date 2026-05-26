package controller

import (
	"bishe/model"
	"bishe/service"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 课程添加
func AddCourseHandler(c *gin.Context) {
	// 从表单中获取用户信息
	title := c.PostForm("title")
	content := c.PostForm("content")
	priceStr := c.PostForm("price")

	titleLen := len(title)
	if titleLen > model.COURSE_TITLE_MAX || titleLen == 0 {
		MakeApiResponseError(c, CODE_COURSE_TITLE_LEN_INVASLID)
		return
	}

	contentLen := len(content)
	if contentLen > model.COURSE_CONTENT_MAX || contentLen == 0 {
		MakeApiResponseError(c, CODE_COURSE_CONTENT_LEN_INVASLID)
		return
	}

	price, err := strconv.Atoi(priceStr)
	if err != nil {
		MakeApiResponseErrorParams(c)
		return
	}

	if price > model.COURSE_PRICE_MAX {
		MakeApiResponseError(c, CODE_COURSE_PRICE_MAX_INVASLID)
		return
	}

	uid := service.GetUidFromContext(c)

	courseImgPath := c.PostForm("img")
	payAccount := c.PostForm("pay_account")

	lockKey := "course-add-" + title
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

	createTime := time.Now()

	// 构造课程
	course := &model.Course{
		Title:      title,
		Content:    content,
		Img:        courseImgPath,
		PayAccount: payAccount,
		Uid:        uid,
		Price:      price,
		CreateAt:   &createTime,
		UpdateAt:   &createTime,
		IsDeleted:  model.IS_DELETED_NO,
	}

	err = service.CreateCourse(course)
	if err != nil {
		service.Logger.Error("CreateCourse err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	data := map[string]interface{}{
		"course": course,
	}

	MakeApiResponseSuccess(c, data)
}

// 获取全部课程
func GetAllCourseHandler(c *gin.Context) {
	pageStr := c.Query("page")
	page := GetPage(pageStr)

	pagesize := 10
	//获取全部课程
	courses, err := service.GetAllCourse(page, pagesize)
	if err != nil {
		service.Logger.Error("GetAllCourse err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if courses == nil {
		courses = make([]model.Course, 0)
	}

	MakeApiResponseSuccess(c, map[string]interface{}{
		"courses": courses,
	})

}

// 根据uid获取用户发布的课程列表
func GetUserAllCourseByUidHandler(c *gin.Context) {
	uid := service.GetUidFromContext(c)

	statusStr := c.Query("status")
	if statusStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	status, err := strconv.Atoi(statusStr)
	if err != nil {
		MakeApiResponseErrorParams(c)
		return
	}

	if status != model.COURSE_STATUS_PUBLISHED && status != model.COURSE_STATUS_UNPUBLISHED {
		MakeApiResponseErrorParams(c)
		return
	}

	pageStr := c.Query("page")
	page := GetPage(pageStr)

	pagesize := 10

	//获取全部课程
	courses, err := service.GetUserAllCourseByUid(uid, status, page, pagesize)
	if err != nil {
		service.Logger.Error("GetUserAllCourseByUid", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if courses == nil {
		courses = make([]model.Course, 0)
	}

	MakeApiResponseSuccess(c, map[string]interface{}{
		"courses": courses,
	})
}

// 获取课程
func GetCourseByTitleHandler(c *gin.Context) {

	title := c.Query("title")
	if title == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	pageStr := c.Query("page")
	page := GetPage(pageStr)

	pagesize := 10

	//获取标题包含title的课程
	courses, err := service.GetAllCourseByTitle(title, page, pagesize)
	if err != nil {
		service.Logger.Error("GetAllCourseByTitle", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if len(courses) == 0 {
		courses = make([]model.Course, 0)
	}

	data := map[string]interface{}{
		"courses": courses,
	}

	MakeApiResponseSuccess(c, data)
}

// 获取课程详情
func GetCourseHandler(c *gin.Context) {
	cidStr := c.Query("course_id")
	if cidStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	cid, err := strconv.Atoi(cidStr)
	if err != nil {
		MakeApiResponseErrorDefault(c)
		return
	}

	course, err := service.GetCourseById(cid)
	if err != nil {
		service.Logger.Error("GetCourseById", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if course == nil {
		course = &model.Course{}
	}

	uid := service.GetUidFromContext(c)
	isOwner := course.Uid == uid
	isFree := course.Price == 0

	data := map[string]interface{}{
		"course":   course,
		"is_owner": isOwner,
		"is_free":  isFree,
	}

	MakeApiResponseSuccess(c, data)
}

// 修改课程
func UpdateCourseHandler(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")
	priceStr := c.PostForm("price")

	titleLen := len(title)
	if titleLen > model.COURSE_TITLE_MAX || titleLen == 0 {
		MakeApiResponseError(c, CODE_COURSE_TITLE_LEN_INVASLID)
		return
	}

	contentLen := len(content)
	if contentLen > model.COURSE_CONTENT_MAX || contentLen == 0 {
		MakeApiResponseError(c, CODE_COURSE_CONTENT_LEN_INVASLID)
		return
	}

	cidStr := c.PostForm("course_id")
	if cidStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	cid, err := strconv.Atoi(cidStr)
	if err != nil {
		MakeApiResponseErrorDefault(c)
		return
	}

	price, err := strconv.Atoi(priceStr)
	if err != nil {
		MakeApiResponseErrorDefault(c)
		return
	}

	course, err := service.GetCourseById(cid)
	if err != nil {
		service.Logger.Error("GetCourseById", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if course == nil {
		service.Logger.Error("GetCourseById course == nil")
		MakeApiResponseErrorDefault(c)
		return
	}

	courseImgPath := c.PostForm("img")
	payAccount := c.PostForm("pay_account")

	courseMap := map[string]interface{}{
		"title":   title,
		"content": content,
		"price":   price,
	}
	if courseImgPath != "" {
		courseMap["img"] = courseImgPath
	}
	if payAccount != "" {
		courseMap["pay_account"] = payAccount
	}

	_, err = service.UpdateCourse(cid, courseMap)
	if err != nil {
		service.Logger.Error("UpdateCourse err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	MakeApiResponseSuccessDefault(c)
}

func PublishCourseHandler(c *gin.Context) {
	cidStr := c.PostForm("course_id")
	if cidStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	cid, err := strconv.Atoi(cidStr)
	if err != nil {
		MakeApiResponseErrorDefault(c)
		return
	}

	uid := service.GetUidFromContext(c)

	course, err := service.GetCourseById(cid)
	if err != nil {
		service.Logger.Error("GetCourseById err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if course == nil {
		MakeApiResponseError(c, CODE_COURSE_NOT_EXIST)
		return
	}

	if course.Uid != uid {
		MakeApiResponseErrorDefault(c)
		return
	}

	if course.CourseStatus == model.COURSE_STATUS_PUBLISHED {
		MakeApiResponseErrorDefault(c)
		return
	}

	updateMap := map[string]interface{}{
		"course_status": model.COURSE_STATUS_PUBLISHED,
	}

	rowsAffected, err := service.UpdateCourse(cid, updateMap)
	if err != nil || rowsAffected == 0 {
		service.Logger.Error("UpdateCourse err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	MakeApiResponseSuccessDefault(c)
}

// 添加课时
func AddLessonHandler(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")

	titleLen := len(title)
	if titleLen > model.LESSON_TITLE_MAX || titleLen == 0 {
		MakeApiResponseError(c, CODE_LESSON_TITLE_LEN_INVASLID)
		return
	}

	contentLen := len(content)
	if contentLen > model.LESSON_CONTENT_MAX || contentLen == 0 {
		MakeApiResponseError(c, CODE_LESSON_CONTENT_LEN_INVASLID)
		return
	}

	cidStr := c.PostForm("course_id")
	if cidStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	cid, err := strconv.Atoi(cidStr)
	if err != nil {
		MakeApiResponseErrorDefault(c)
		return
	}

	createTime := time.Now()

	// 构造课题
	lesson := &model.Lesson{ //其中包含自动生成的id
		Title:     title,
		Content:   content,
		CourseId:  cid,
		CreateAt:  &createTime,
		UpdateAt:  &createTime,
		IsDeleted: model.IS_DELETED_NO,
	}

	//创建课时
	err = service.CreateLesson(lesson)
	if err != nil {
		service.Logger.Error("CreateLesson err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	data := map[string]interface{}{
		"lessonId": lesson.Id,
	}

	MakeApiResponseSuccess(c, data)
}

// 获取课时详情
func GetLessonHandler(c *gin.Context) {
	lessonIdStr := c.Query("lesson_id")
	lessonId, _ := strconv.Atoi(lessonIdStr)
	if lessonId == 0 {
		lessonId = 1
	}

	lesson, err := service.GetLessonById(lessonId)
	if err != nil {
		service.Logger.Error("GetLessonById", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if lesson == nil {
		lesson = &model.Lesson{}
	}

	data := map[string]interface{}{
		"lesson": lesson,
	}

	MakeApiResponseSuccess(c, data)
}

// 获取课程全部课时
func GetCourseAllLessonHandler(c *gin.Context) {
	cidStr := c.Query("course_id")
	if cidStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	cid, err := strconv.Atoi(cidStr)
	if err != nil {
		MakeApiResponseErrorDefault(c)
		return
	}

	lessons, err := service.GetAllLessonByCid(cid)
	if err != nil {
		service.Logger.Error("GetAllLessonByCid", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if len(lessons) == 0 {
		lessons = make([]model.Lesson, 0)
	}

	data := map[string]interface{}{
		"lessons": lessons,
	}

	MakeApiResponseSuccess(c, data)
}

// 修改课时
func UpdateLessonHandler(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")

	titleLen := len(title)
	if titleLen > model.LESSON_TITLE_MAX || titleLen == 0 {
		MakeApiResponseError(c, CODE_LESSON_TITLE_LEN_INVASLID)
		return
	}

	contentLen := len(content)
	if contentLen > model.LESSON_CONTENT_MAX || contentLen == 0 {
		MakeApiResponseError(c, CODE_LESSON_CONTENT_LEN_INVASLID)
		return
	}

	lessonIdStr := c.PostForm("lesson_id")
	if lessonIdStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	lessonId, err := strconv.Atoi(lessonIdStr)
	if err != nil {
		MakeApiResponseErrorDefault(c)
		return
	}

	lesson, err := service.GetLessonById(lessonId)
	if err != nil {
		service.Logger.Error("GetLessonById err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if lesson == nil {
		service.Logger.Error("GetLessonById lesson == nil")
		MakeApiResponseErrorDefault(c)
		return
	}

	updateMap := map[string]interface{}{
		"title":   title,
		"content": content,
	}

	rowsAffected, err := service.UpdateLesson(lessonId, updateMap)
	if err != nil || rowsAffected == 0 {
		service.Logger.Error("UpdateLesson err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	MakeApiResponseSuccessDefault(c)
}

// 购买课程
func AddPurchaseHandler(c *gin.Context) {
	uid := service.GetUidFromContext(c)

	cidStr := c.PostForm("course_id")
	if cidStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	cid, err := strconv.Atoi(cidStr)
	if err != nil {
		MakeApiResponseErrorDefault(c)
		return
	}

	// 获取用户该课程的购买记录
	purchases, err := service.GetPurchaseByUidCid(uid, cid)
	if err != nil {
		service.Logger.Error("GetPurchaseByUidCid err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	// 获取课程信息
	course, err := service.GetCourseById(cid)
	if err != nil || course == nil {
		service.Logger.Error("GetCourseById err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if len(purchases) > 0 {
		for _, v := range purchases {
			if v.PurchaseStatus == model.PURCHASE_STATUS_UNPAID {
				// 已有该课程未支付订单，提示前端跳转到购买订单页
				MakeApiResponseError(c, CODE_HAS_UNPAY_ORDER)
				return
			}

			if v.PurchaseStatus == model.PURCHASE_STATUS_PAID {
				MakeApiResponseError(c, CODE_USER_PURCHASED)
				return
			}
		}
	}

	createTime := time.Now()
	purchase := &model.Purchase{
		UserId:         uid,
		CourseId:       cid,
		CreateAt:       &createTime,
		UpdateAt:       &createTime,
		PurchaseStatus: model.PURCHASE_STATUS_UNPAID,
	}

	//创建购买记录
	err = service.CreatePurchase(purchase)
	if err != nil {
		service.Logger.Error("CreatePurchase err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	payUrl := getCoursePayUrl(c, purchase.Id, cid, course.Price, course.Title)
	qrCodeUrl, err := service.QrcodeImgSave(payUrl, 200, service.FileTypeCoursePayImg, createTime)
	if err != nil {
		service.Logger.Error("QrcodeImgSave err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	MakeApiResponseSuccess(c, map[string]interface{}{
		"purchase_id": purchase.Id,
		"pay_url":     payUrl,
		"qr_code_url": qrCodeUrl,
		"price":       course.Price,
	})
}

// 学习免费课程
func AddPurchaseFreeHandler(c *gin.Context) {
	uid := service.GetUidFromContext(c)

	cidStr := c.PostForm("course_id")
	if cidStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	cid, err := strconv.Atoi(cidStr)
	if err != nil {
		MakeApiResponseErrorDefault(c)
		return
	}

	// 获取用户该课程的购买记录
	purchases, err := service.GetPurchaseByUidCid(uid, cid)
	if err != nil {
		service.Logger.Error("GetPurchaseByUidCid err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	// 获取课程信息
	course, err := service.GetCourseById(cid)
	if err != nil || course == nil {
		service.Logger.Error("GetCourseById err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if len(purchases) > 0 {
		for _, v := range purchases {
			if v.PurchaseStatus == model.PURCHASE_STATUS_PAID {
				MakeApiResponseError(c, CODE_USER_PURCHASED)
				return
			}
		}
	}

	createTime := time.Now()
	purchase := &model.Purchase{
		UserId:         uid,
		CourseId:       cid,
		CreateAt:       &createTime,
		UpdateAt:       &createTime,
		PurchaseStatus: model.PURCHASE_STATUS_PAID,
	}

	//创建购买记录
	err = service.CreatePurchase(purchase)
	if err != nil {
		service.Logger.Error("CreatePurchase err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	payUrl := getCoursePayUrl(c, purchase.Id, cid, course.Price, course.Title)
	qrCodeUrl, err := service.QrcodeImgSave(payUrl, 200, service.FileTypeCoursePayImg, createTime)
	if err != nil {
		service.Logger.Error("QrcodeImgSave err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	// 学习免费课程，更新课程加入人数
	rowsAffected, err := service.IncrCourseJoinNumByCid(cid)
	if rowsAffected == 0 || err != nil {
		service.Logger.Error("IncrCourseJoinNumByCid err", zap.Error(err), zap.Int("cid", cid))
		MakeApiResponseErrorDefault(c)
		return
	}

	MakeApiResponseSuccess(c, map[string]interface{}{
		"purchase_id": purchase.Id,
		"pay_url":     payUrl,
		"qr_code_url": qrCodeUrl,
		"price":       course.Price,
	})
}

// 获取购买记录
func GetPurchaseHandler(c *gin.Context) {
	uid := service.GetUidFromContext(c)

	cidStr := c.Query("course_id")
	if cidStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	cid, err := strconv.Atoi(cidStr)
	if err != nil {
		MakeApiResponseErrorDefault(c)
		return
	}

	// 获取用户该课程所有购买记录
	purchases, err := service.GetPurchaseByUidCid(uid, cid)
	if err != nil {
		service.Logger.Error("GetPurchaseByUidCid", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	isPurchased := false

	for _, v := range purchases {
		if v.PurchaseStatus == model.PURCHASE_STATUS_UNPAID {
			MakeApiResponseError(c, CODE_HAS_UNPAY_ORDER)
			return
		}
		if v.PurchaseStatus == model.PURCHASE_STATUS_PAID {
			isPurchased = true
		}
	}

	data := map[string]interface{}{
		"purchase":    nil,
		"isPurchased": isPurchased,
	}

	MakeApiResponseSuccess(c, data)

}

// 获取用户购买课程列表
func GetUserPurchaseListHandler(c *gin.Context) {
	uid := service.GetUidFromContext(c)

	// 获取用户购买记录
	purchases, err := service.GetAllPurchaseByUid(uid)
	if err != nil {
		service.Logger.Error("GetAllPurchaseByUid err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if len(purchases) == 0 {
		purchases = make([]model.Purchase, 0)
	}

	var courseIds []int
	for _, v := range purchases {
		courseIds = append(courseIds, v.CourseId)
	}

	if len(courseIds) == 0 {
		data := map[string]interface{}{
			"purchaseCourses": make([]model.PurchaseCourse, 0),
		}
		MakeApiResponseSuccess(c, data)
		return
	}

	//根据courseIds获取课程map
	courseMap, err := service.GetCourseMapByCourseIds(courseIds)
	if err != nil {
		service.Logger.Error("GetCourseMapByCourseIds err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	var purchaseCourses []model.PurchaseCourse

	for _, v := range purchases {
		var purchaseCourse model.PurchaseCourse
		course, _ := courseMap[v.CourseId]

		purchaseCourse.Purchase = v
		purchaseCourse.Course = course
		purchaseCourses = append(purchaseCourses, purchaseCourse)

	}

	data := map[string]interface{}{
		"purchaseCourses": purchaseCourses,
	}

	MakeApiResponseSuccess(c, data)
}

// 更新购买记录状态购买课程
func UpdatePurchaseStatusHandler(c *gin.Context) {

	purchaseIdStr := c.PostForm("purchase_id")
	if purchaseIdStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	purchaseId, err := strconv.Atoi(purchaseIdStr)
	if err != nil {
		MakeApiResponseErrorDefault(c)
		return
	}

	// 获取用户购买记录
	purchase, err := service.GetPurchaseById(purchaseId)
	if err != nil {
		service.Logger.Error("GetPurchaseById", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if purchase == nil {
		MakeApiResponseErrorDefault(c)
		return
	}

	status := purchase.PurchaseStatus

	statusNew, err := service.MakePurchaseStatus(status, model.PURCHASE_ACTION_PAY)
	if err != nil {
		service.Logger.Error("MakePurchaseStatus", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	//TODO加事务
	//roback撤销，下面两个操作数据库，要么都成功，要么都失败

	// // 更新用户购买记录状态
	// affectRows, err := service.UpdatePurchaseStatusById(purchaseId, status, statusNew)
	// if affectRows == 0 || err != nil {
	// 	service.Logger.Error("UpdatePurchaseStatusById err", zap.Error(err))
	// 	MakeApiResponseErrorDefault(c)
	// 	return
	// }

	// //更新course join_num+1
	// affectRows, err = service.IncrCourseJoinNumByCid(purchase.CourseId)
	// if affectRows == 0 || err != nil {
	// 	service.Logger.Error("IncrCourseJoinNumByCid err", zap.Error(err))
	// 	MakeApiResponseErrorDefault(c)
	// 	return
	// }

	// 更新用户购买记录状态
	err = service.UpdatePurchaseStatusAndJoinNum(purchaseId, status, statusNew, purchase.CourseId)
	if err != nil {
		service.Logger.Error("UpdatePurchaseStatusAndJoinNum err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	MakeApiResponseSuccessDefault(c)
}

// GetPurchaseByPayIdHandler 通过购买id获取购买记录（无需登录，供扫码支付轮询）
func GetPurchaseByPayIdHandler(c *gin.Context) {
	purchaseIdStr := c.Query("purchase_id")
	if purchaseIdStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	purchaseId, err := strconv.Atoi(purchaseIdStr)
	if err != nil {
		MakeApiResponseErrorParams(c)
		return
	}

	purchase, err := service.GetPurchaseById(purchaseId)
	if err != nil {
		service.Logger.Error("GetPurchaseById", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if purchase == nil {
		MakeApiResponseError(c, CODE_ORDERS_NOT_EXIST)
		return
	}

	MakeApiResponseSuccess(c, map[string]interface{}{
		"purchase": purchase,
	})
}

func getCoursePayUrl(c *gin.Context, purchaseId int, courseId int, price int, title string) string {
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

	return "http://" + ip + ":" + port + "/page/purchase/pay?purchase_id=" + strconv.Itoa(purchaseId) +
		"&course_id=" + strconv.Itoa(courseId) + "&price=" + strconv.Itoa(price) +
		"&title=" + url.QueryEscape(title)
}

// 取消未支付购买
func CancelPurchaseHandler(c *gin.Context) {
	uid := service.GetUidFromContext(c)

	idStr := c.PostForm("purchase_id")
	if idStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		MakeApiResponseErrorParams(c)
		return
	}

	purchase, err := service.GetPurchaseById(id)
	if err != nil {
		service.Logger.Error("GetPurchaseById err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if purchase == nil {
		MakeApiResponseError(c, CODE_ORDERS_NOT_EXIST)
		return
	}

	if purchase.UserId != uid {
		MakeApiResponseErrorParams(c)
		return
	}

	if purchase.PurchaseStatus != model.PURCHASE_STATUS_UNPAID {
		MakeApiResponseErrorDefault(c)
		return
	}

	statusNew, err := service.MakePurchaseStatus(purchase.PurchaseStatus, model.PURCHASE_ACTION_CANCEL)
	if err != nil {
		service.Logger.Error("MakePurchaseStatus err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	affectRows, err := service.UpdatePurchaseStatusById(id, purchase.PurchaseStatus, statusNew)
	if err != nil || affectRows == 0 {
		service.Logger.Error("UpdatePurchaseStatusById err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	MakeApiResponseSuccessDefault(c)
}

// 获取购买二维码
func GetPurchaseQrcodeHandler(c *gin.Context) {
	purchaseIdStr := c.Query("purchase_id")
	if purchaseIdStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	purchaseId, err := strconv.Atoi(purchaseIdStr)
	if err != nil {
		MakeApiResponseErrorParams(c)
		return
	}

	purchase, err := service.GetPurchaseById(purchaseId)
	if err != nil {
		service.Logger.Error("GetPurchaseById err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if purchase == nil {
		MakeApiResponseError(c, CODE_ORDERS_NOT_EXIST)
		return
	}

	if purchase.PurchaseStatus != model.PURCHASE_STATUS_UNPAID {
		MakeApiResponseError(c, CODE_HAS_UNPAY_ORDER)
		return
	}

	course, err := service.GetCourseById(purchase.CourseId)
	if err != nil || course == nil {
		service.Logger.Error("GetCourseById err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	payUrl := getCoursePayUrl(c, purchase.Id, purchase.CourseId, course.Price, course.Title)
	qrCodeUrl, err := service.QrcodeImgSave(payUrl, 200, service.FileTypeCoursePayImg, *purchase.CreateAt)
	if err != nil {
		service.Logger.Error("QrcodeImgSave err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	MakeApiResponseSuccess(c, map[string]interface{}{
		"purchase_id": purchase.Id,
		"pay_url":     payUrl,
		"qr_code_url": qrCodeUrl,
	})
}
