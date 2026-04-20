package controller

import (
	"bishe/internal/app/knowledge_sharing/model"
	"bishe/internal/app/knowledge_sharing/service"
	"strconv"
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

	cidStr := c.PostForm("cid")
	if cidStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	cid, err := strconv.Atoi(cidStr)
	if err != nil {
		MakeApiResponseErrorDefault(c)
		return
	}

	uid, _ := service.GetUserFromCookie(c)
	if uid == 0 {
		MakeApiResponseError(c, CODE_USER_NOT_LOGIN)
		return
	}

	createTime := time.Now()

	// 构造课程
	course := &model.Course{ //其中包含自动生成的id
		Title:     title,
		Content:   content,
		Cid:       cid,
		Uid:       uid,
		Price:     price,
		CreateAt:  &createTime,
		UpdateAt:  &createTime,
		IsDeleted: model.IS_DELETED_NO,
	}

	err = service.CreateCourse(course)
	if err != nil {
		service.Logger.Error("CreateCourse err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	MakeApiResponseSuccess(c, map[string]interface{}{})

}

// 根据uid获取用户的课程列表
func GetUserAllCourseByUidHandler(c *gin.Context) {
	uidStr := c.Query("uid")
	if uidStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	uid, err := strconv.Atoi(uidStr)
	if err != nil {
		MakeApiResponseErrorParams(c)
		return
	}

	pageStr := c.Query("page")
	page := GetPage(pageStr)

	pagesize := 10

	//获取全部课程
	courses, err := service.GetUserAllCourseByUid(uid, page, pagesize)
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

// 获取圈子全部课程
func GetCircleAllCourseHandler(c *gin.Context) {

	cidStr := c.Query("cid")
	if cidStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	cid, err := strconv.Atoi(cidStr)
	if err != nil {
		MakeApiResponseErrorDefault(c)
		return
	}

	pageStr := c.Query("page")
	page := GetPage(pageStr)

	pagesize := 10

	//获取圈子中全部课程
	courses, err := service.GetAllCourseByCid(cid, page, pagesize)
	if err != nil {
		service.Logger.Error("GetAllCourseByCid", zap.Error(err))
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

// 获取课程
func GetCourseByTitleHandler(c *gin.Context) {

	cidStr := c.Query("cid")
	if cidStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	cid, err := strconv.Atoi(cidStr)
	if err != nil {
		MakeApiResponseErrorDefault(c)
		return
	}
	title := c.Query("title")

	if title == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	pageStr := c.Query("page")
	page := GetPage(pageStr)

	pagesize := 10

	//获取圈子中标题包含title的课程
	courses, err := service.GetAllCourseByTitle(cid, title, page, pagesize)
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
	cidStr := c.Query("cid")
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

	MakeApiResponseSuccess(c, map[string]interface{}{
		"course": course,
	})
}

func AddLessonHandler(c *gin.Context) {
	// 从表单中获取用户信息
	title := c.PostForm("title")
	content := c.PostForm("content")

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

	cidStr := c.PostForm("cid")
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
}
