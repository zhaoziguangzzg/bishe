package mysql

import (
	"bishe/internal/app/knowledge_sharing/model"

	"gorm.io/gorm"
)

// create 用户反馈
func CreateUserFeedback(newFeedback *model.Feedback) (err error) {
	err = DB.Model(&model.Feedback{}).Create(newFeedback).Error
	return
}

// 获取全部未处理反馈
func GetAllFeedback(page int, pagesize int) (feedbacks []model.Feedback, err error) {
	offset := (page - 1) * pagesize

	err = DB.Model(&model.Feedback{}).
		Where("feedback_status=? and is_deleted=?", model.FEEDBACK_STATUS_WAIT, model.FEEDBACK_NOT_DELETED).
		Order("feedback_time DESC").Offset(offset).Limit(pagesize).Find(&feedbacks).Error
	return
}

// 获取文章反馈内容
func GetFeedbackById(id int) (feedback *model.Feedback, err error) {
	feedback = new(model.Feedback)
	err = DB.Model(&model.Feedback{}).Where("id=?", id).First(&feedback).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound { //没查到数据返回空
			return nil, nil
		}

		return nil, err
	}

	return feedback, nil
}

// 更新反馈状态为无问题
func UpdateFeedbackNormalById(id int) (int64, error) {
	result := DB.Model(&model.Feedback{}).Where("id=?", id).Update("feedback_status", model.FEEDBACK_STATUS_NORMAL)
	return result.RowsAffected, result.Error
}

// 更新反馈状态为有问题
func UpdateFeedbackViolateById(id int) (int64, error) {
	result := DB.Model(&model.Feedback{}).Where("id=?", id).Update("feedback_status", model.FEEDBACK_STATUS_QUESTIONABLE)
	return result.RowsAffected, result.Error
}
