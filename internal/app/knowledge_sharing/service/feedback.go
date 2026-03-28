package service

import (
	"bishe/internal/app/knowledge_sharing/dao/mysql"
	"bishe/internal/app/knowledge_sharing/model"
	"time"
)

// create 用户反馈
func CreateUserFeedback(newFeedback *model.Feedback) (err error) {
	return mysql.CreateUserFeedback(newFeedback)
}

// 获取全部未处理反馈
func GetAllFeedback(page int, pagesize int) (feedbacks []model.Feedback, err error) {
	return mysql.GetAllFeedback(page, pagesize)
}

// 获取反馈内容
func GetFeedbackById(id int) (feedback *model.Feedback, err error) {
	return mysql.GetFeedbackById(id)
}

// 更新反馈状态回复
func UpdateFeedbackStatusReplyById(id int, reply string, replyTime time.Time) (int64, error) {
	return mysql.UpdateFeedbackStatusReplyById(id, reply, replyTime)
}
