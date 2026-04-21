package mysql

import "bishe/internal/app/knowledge_sharing/model"

//创建课时
func CreateLesson(lesson *model.Lesson) (err error) {
	err = DB.Model(&model.Lesson{}).Create(lesson).Error
	return
}

//获取课时详情
func GetLessonById(lessonId int) (lesson *model.Lesson, err error) {
	err = DB.Model(&model.Lesson{}).Where("id = ?", lessonId).First(&lesson).Error
	return
}

//获取课程全部课时
func GetAllLessonByCid(cid int) (lessons []model.Lesson, err error) {
	err = DB.Model(&model.Lesson{}).Where("course_id = ?", cid).Find(&lessons).Error
	return
}
