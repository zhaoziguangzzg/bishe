package service

import (
	"bishe/internal/app/knowledge_sharing/dao/mysql"
	"bishe/internal/app/knowledge_sharing/model"
)

// 创建课时
func CreateLesson(lesson *model.Lesson) (err error) {
	return mysql.CreateLesson(lesson)
}

// 获取课时详情
func GetLessonById(lessonId int) (lesson *model.Lesson, err error) {
	return mysql.GetLessonById(lessonId)
}

// 获取课程全部课时
func GetAllLessonByCid(cid int) (lessons []model.Lesson, err error) {
	return mysql.GetAllLessonByCid(cid)
}

// 更新课时信息
func UpdateLesson(lessonId int, updateMap map[string]interface{}) (int64, error) {
	return mysql.UpdateLesson(lessonId, updateMap)
}
