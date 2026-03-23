package model

import "time"

//Follow 定义用户关注结构体
type Follow struct {
	Id            int        `json:"id" gorm:"column:id" mapstructure:"id"`
	FanId         int        `json:"fanId" gorm:"column:fan_id" mapstructure:"fanId"`
	FollowerId    int        `json:"followerId" gorm:"column:follower_id" mapstructure:"followerId"`
	FollowTime    *time.Time `json:"followTime" gorm:"column:follow_time" mapstructure:"-"`
	FollowTimeStr string     `json:"-" gorm:"-" mapstructure:"followTime"`
	CreateAt      *time.Time `json:"createAt" gorm:"column:create_at" mapstructure:"-"`
	CreateAtStr   string     `json:"-" gorm:"-" mapstructure:"createAt"`
	UpdateAt      *time.Time `json:"updateAt" gorm:"column:update_at" mapstructure:"-"`
	UpdateAtStr   string     `json:"-" gorm:"-" mapstructure:"updateAt"`
	FollowStatus  int        `json:"followStatus" gorm:"column:follow_status" mapstructure:"followStatus"`
}

const (
	FOLLOW_STATUS_NORMAL int = 0 //关注
	FOLLOW_STATUS_REVIEW int = 1 //未关注
)

//指定Follow的表名
func (Follow) TableName() string {
	return "follow"
}
