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
	//回复内容
	Reply string `json:"reply" gorm:"column:reply" mapstructure:"reply"`
	//回复时间
	ReplyTime *time.Time `json:"replyTime" gorm:"column:reply_time" mapstructure:"replyTime"`

	CreateAt       *time.Time `json:"createAt" gorm:"column:create_at" mapstructure:"-"`
	CreateAtStr    string     `json:"-" gorm:"-" mapstructure:"createAt"`
	UpdateAt       *time.Time `json:"updateAt" gorm:"column:update_at" mapstructure:"-"`
	UpdateAtStr    string     `json:"-" gorm:"-" mapstructure:"updateAt"`
	FeedbackStatus int        `json:"feedbackStatus" gorm:"column:feedback_status" mapstructure:"feedbackStatus"`
	IsDeleted      int        `json:"isDeleted" gorm:"column:is_deleted" mapstructture:"isDeleted"`
}

const (
	FEEDBACK_STATUS_OPEN  int = 0 //打开，未处理
	FEEDBACK_STATUS_CLOSE int = 1 //关闭，已处理

	FEEDBACK_NOT_DELETED int = 0 //未被删除
	FEEDBACK_IS_DELETED  int = 1 //被删除

	FEEDBACK_MAX_CONTENT int = 100 //反馈内容最长100字

)

// 指定Feedback对应的表名
func (Feedback) TableName() string {
	return "feedback"
}
