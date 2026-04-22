package mysql

import (
	"bishe/internal/app/knowledge_sharing/model"

	"gorm.io/gorm"
)

// create课程
func CreateCourse(course *model.Course) (err error) {
	err = DB.Model(&model.Course{}).Create(course).Error
	return
}

// 获取全部课程
func GetAllCourse(page int, pagesize int) (courses []model.Course, err error) {
	offset := (page - 1) * pagesize
	err = DB.Model(&model.Course{}).
		Where("course_status=? and is_deleted=?", model.COURSE_STATUS_PUBLISHED, model.IS_DELETED_NO).
		Order("id DESC").Offset(offset).
		Limit(pagesize).Find(&courses).Error
	if err != nil {
		return
	}
	return
}

// get用户发布的课程列表
func GetUserAllCourseByUid(uid int, page int, pagesize int) (courses []model.Course, err error) {
	offset := (page - 1) * pagesize
	err = DB.Model(&model.Course{}).Where("uid = ?", uid).
		Order("id desc").Offset(offset).
		Limit(pagesize).Find(&courses).Error
	if err != nil {
		return
	}
	return
}

// get标题包含title的课程
func GetAllCourseByTitle(title string, page int, pagesize int) (courses []model.Course, err error) {
	offset := (page - 1) * pagesize
	err = DB.Model(&model.Course{}).
		Where("course_status=? and is_deleted=? and title like ?", model.COURSE_STATUS_PUBLISHED, model.IS_DELETED_NO, "%"+title+"%").
		Order("id DESC").Offset(offset).Limit(pagesize).Find(&courses).Error
	if err != nil {
		return
	}
	return
}

// 根据cid获取课程
func GetCourseById(cid int) (course *model.Course, err error) {
	course = new(model.Course)
	err = DB.Model(&model.Course{}).Where("id=? and is_deleted=?", cid, model.IS_DELETED_NO).First(&course).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound { //没查到数据返回空
			return nil, nil
		}

		return nil, err
	}

	return course, nil
}
