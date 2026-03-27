package service

import (
	"bishe/internal/app/knowledge_sharing/dao/mysql"
	"bishe/internal/app/knowledge_sharing/model"
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

// 更新反馈状态为无问题
func UpdateFeedbackNormalById(id int) (int64, error) {
	return mysql.UpdateFeedbackNormalById(id)
}

// 更新反馈状态为有问题
func UpdateFeedbackViolateById(id int) (int64, error) {
	return mysql.UpdateFeedbackViolateById(id)
}
