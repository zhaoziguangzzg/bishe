package mysql

import (
	"bishe/internal/app/knowledge_sharing/model"
	"time"

	"gorm.io/gorm"
)

// create 用户反馈
func CreateUserFeedback(newFeedback *model.Feedback) (err error) {
	err = DB.Model(&model.Feedback{}).Create(newFeedback).Error
	return
}

// 获取全部反馈
func GetAllFeedback(page int, pagesize int, status int) (feedbacks []model.Feedback, err error) {
	offset := (page - 1) * pagesize

	err = DB.Model(&model.Feedback{}).
		Where("feedback_status=? and is_deleted=?", status, model.FEEDBACK_NOT_DELETED).
		Order("feedback_time DESC").Offset(offset).Limit(pagesize).Find(&feedbacks).Error
	return
}

// 根据用户ID获取反馈列表
func GetFeedbackByUid(uid int, page int, pagesize int) (feedbacks []model.Feedback, err error) {
	offset := (page - 1) * pagesize

	err = DB.Model(&model.Feedback{}).
		Where("user_id=? and is_deleted=?", uid, model.FEEDBACK_NOT_DELETED).
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

// 更新反馈状态回复
func UpdateFeedbackStatusReplyById(id int, reply string, replyTime time.Time) (int64, error) {
	feedback := map[string]interface{}{
		"feedback_status": model.FEEDBACK_STATUS_CLOSE,
		"reply":           reply,
		"reply_time":      replyTime,
	}

	result := DB.Model(&model.Feedback{}).Where("id=?", id).Updates(feedback)
	return result.RowsAffected, result.Error
}
