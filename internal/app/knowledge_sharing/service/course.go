package service

import (
	"bishe/internal/app/knowledge_sharing/dao/mysql"
	"bishe/internal/app/knowledge_sharing/model"
)

// create课程
func CreateCourse(course *model.Course) (err error) {
	return mysql.CreateCourse(course)
}

// get用户发布的课程列表
func GetUserAllCourseByUid(uid int, page int, pagesize int) (courses []model.Course, err error) {
	return mysql.GetUserAllCourseByUid(uid, page, pagesize)
}

// get圈子中的课程
func GetAllCourseByCid(cid int, page int, pagesize int) (courses []model.Course, err error) {
	return mysql.GetAllCourseByCid(cid, page, pagesize)
}

// get圈子中标题包含title的课程
func GetAllCourseByTitle(cid int, title string, page int, pagesize int) (courses []model.Course, err error) {
	return mysql.GetAllCourseByTitle(cid, title, page, pagesize)
}

// 根据cid获取课程
func GetCourseById(cid int) (course *model.Course, err error) {
	return mysql.GetCourseById(cid)
}
