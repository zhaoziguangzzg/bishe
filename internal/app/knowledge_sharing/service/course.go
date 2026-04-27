package service

import (
	"bishe/internal/app/knowledge_sharing/dao/mysql"
	"bishe/internal/app/knowledge_sharing/model"

	"github.com/gin-gonic/gin"
)

func SetCourseIdToContext(c *gin.Context, cid int) {
	c.Set("course_id", cid)
}

func GetCourseIdFromContext(c *gin.Context) (cid int) {
	cid = c.GetInt("course_id")
	return
}

func SetCourseToContext(c *gin.Context, course *model.Course) {
	c.Set("course", course)
}

// 从context获取文章
func GetCourseFromContext(c *gin.Context) (course *model.Course, ok bool) {
	courseAny, ok := c.Get("course")
	if !ok {
		return
	}

	course, ok = courseAny.(*model.Course)
	if !ok {
		return
	}
	return
}

// create课程
func CreateCourse(course *model.Course) (err error) {
	return mysql.CreateCourse(course)
}

// 获取全部课程
func GetAllCourse(page int, pagesize int) (courses []model.Course, err error) {
	return mysql.GetAllCourse(page, pagesize)
}

// get用户发布的课程列表
func GetUserAllCourseByUid(uid int, status int, page int, pagesize int) (courses []model.Course, err error) {
	return mysql.GetUserAllCourseByUid(uid, status, page, pagesize)
}

// get圈子中标题包含title的课程
func GetAllCourseByTitle(title string, page int, pagesize int) (courses []model.Course, err error) {
	return mysql.GetAllCourseByTitle(title, page, pagesize)
}

// 根据cid获取课程
func GetCourseById(cid int) (course *model.Course, err error) {
	return mysql.GetCourseById(cid)
}

// 根据courseIds获取courseMap
func GetCourseMapByCourseIds(courseIds []int) (courseMap map[int]model.Course, err error) {
	return mysql.GetCourseMapByCourseIds(courseIds)
}
