package model

import "time"

// Feedback 定义反馈结构体
type Feedback struct {
	Id int `json:"id" gorm:"column:id" mapstructure:"id"`
	//反馈的用户id
	UserId int `json:"userId" gorm:"column:user_id" mapstructure:"userId"`
	//反馈内容
	Content string `json:"content" gorm:"column:content" mapstructure:"content"`
	//反馈时间
	FeedbackTime *time.Time `json:"feedbackTime" gorm:"column:feedback_time" mapstructure:"feedbackTime"`

	CreateAt       *time.Time `json:"createAt" gorm:"column:create_at" mapstructure:"-"`
	CreateAtStr    string     `json:"-" gorm:"-" mapstructure:"createAt"`
	UpdateAt       *time.Time `json:"updateAt" gorm:"column:update_at" mapstructure:"-"`
	UpdateAtStr    string     `json:"-" gorm:"-" mapstructure:"updateAt"`
	FeedbackStatus int        `json:"feedbackStatus" gorm:"column:feedback_status" mapstructure:"feedbackStatus"`
	IsDeleted      int        `json:"isDeleted" gorm:"column:is_deleted" mapstructture:"isDeleted"`
}

const (
	FEEDBACK_STATUS_WAIT         int = 0 //待处理
	FEEDBACK_STATUS_NORMAL       int = 1 //无问题
	FEEDBACK_STATUS_QUESTIONABLE int = 2 //有问题

	FEEDBACK_NOT_DELETED int = 0 //未被删除
	FEEDBACK_IS_DELETED  int = 1 //被删除

	FEEDBACK_MAX_CONTENT int = 100 //反馈内容最长100字

)

// 指定Feedback对应的表名
func (Feedback) TableName() string {
	return "feedback"
}
